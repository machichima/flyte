package admin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/flyteorg/flyte/flyteidl/clients/go/admin/cache"
	"github.com/flyteorg/flyte/flyteidl/clients/go/admin/utils"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/flyteorg/flyte/flytestdlib/logger"
)

const ProxyAuthorizationHeader = "proxy-authorization"

// MaterializeCredentials will attempt to build a TokenSource given the anonymously available information exposed by the server.
// Once established, it'll invoke PerRPCCredentialsFuture.Store() on perRPCCredentials to populate it with the appropriate values.
func MaterializeCredentials(tokenSource oauth2.TokenSource, cfg *Config, authorizationMetadataKey string,
	perRPCCredentials *PerRPCCredentialsFuture) error {
	_, err := tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to issue token. Error: %w", err)
	}

	wrappedTokenSource := NewCustomHeaderTokenSource(tokenSource, cfg.UseInsecureConnection, authorizationMetadataKey)
	perRPCCredentials.Store(wrappedTokenSource)

	return nil
}

// MaterializeInMemoryCredentials initializes the perRPCCredentials with the token source containing in memory cached token.
// This path doesn't perform the token refresh and only build the cred source with cached token.
func MaterializeInMemoryCredentials(ctx context.Context, cfg *Config, tokenCache cache.TokenCache,
	perRPCCredentials *PerRPCCredentialsFuture, authorizationMetadataKey string) error {
	tokenSource, err := NewInMemoryTokenSourceProvider(tokenCache).GetTokenSource(ctx)
	if err != nil {
		return fmt.Errorf("failed to get token source. Error: %w", err)
	}
	wrappedTokenSource := NewCustomHeaderTokenSource(tokenSource, cfg.UseInsecureConnection, authorizationMetadataKey)
	perRPCCredentials.Store(wrappedTokenSource)
	return nil
}

func GetProxyTokenSource(ctx context.Context, cfg *Config) (oauth2.TokenSource, error) {
	tokenSourceProvider, err := NewExternalTokenSourceProvider(cfg.ProxyCommand)
	if err != nil {
		return nil, fmt.Errorf("failed to initialized proxy authorization token source provider. Err: %w", err)
	}
	proxyTokenSource, err := tokenSourceProvider.GetTokenSource(ctx)
	if err != nil {
		return nil, err
	}
	return proxyTokenSource, nil
}

func MaterializeProxyAuthCredentials(ctx context.Context, cfg *Config, proxyCredentialsFuture *PerRPCCredentialsFuture) error {
	proxyTokenSource, err := GetProxyTokenSource(ctx, cfg)
	if err != nil {
		return err
	}

	wrappedTokenSource := NewCustomHeaderTokenSource(proxyTokenSource, cfg.UseInsecureConnection, ProxyAuthorizationHeader)
	proxyCredentialsFuture.Store(wrappedTokenSource)

	return nil
}

func shouldAttemptToAuthenticate(errorCode codes.Code) bool {
	return errorCode == codes.Unauthenticated
}

type proxyAuthTransport struct {
	transport              http.RoundTripper
	proxyCredentialsFuture *PerRPCCredentialsFuture
}

func (c *proxyAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// check if the proxy credentials future is initialized
	if !c.proxyCredentialsFuture.IsInitialized() {
		return nil, errors.New("proxy credentials future is not initialized")
	}

	metadata, err := c.proxyCredentialsFuture.GetRequestMetadata(context.Background(), "")
	if err != nil {
		return nil, err
	}
	token := metadata[ProxyAuthorizationHeader]
	req.Header.Add(ProxyAuthorizationHeader, token)
	return c.transport.RoundTrip(req)
}

// Set up http client used in oauth2
func setHTTPClientContext(ctx context.Context, cfg *Config, proxyCredentialsFuture *PerRPCCredentialsFuture) context.Context {
	httpClient := &http.Client{}
	transport := &http.Transport{}

	if len(cfg.HTTPProxyURL.String()) > 0 {
		// create a transport that uses the proxy
		transport.Proxy = http.ProxyURL(&cfg.HTTPProxyURL.URL)
	}

	if len(cfg.ProxyCommand) > 0 {
		httpClient.Transport = &proxyAuthTransport{
			transport:              transport,
			proxyCredentialsFuture: proxyCredentialsFuture,
		}
	} else {
		httpClient.Transport = transport
	}

	return context.WithValue(ctx, oauth2.HTTPClient, httpClient)
}

type OauthMetadataProvider struct {
	authorizationMetadataKey string
	tokenSource              oauth2.TokenSource
	once                     sync.Once
}

func (o *OauthMetadataProvider) getTokenSourceAndMetadata(cfg *Config, tokenCache cache.TokenCache, proxyCredentialsFuture *PerRPCCredentialsFuture) error {
	ctx := context.Background()

	authMetadataClient, err := InitializeAuthMetadataClient(ctx, cfg, proxyCredentialsFuture)
	if err != nil {
		return fmt.Errorf("failed to initialized Auth Metadata Client. Error: %w", err)
	}

	tokenSourceProvider, err := NewTokenSourceProvider(ctx, cfg, tokenCache, authMetadataClient)
	if err != nil {
		return fmt.Errorf("failed to initialize token source provider. Err: %w", err)
	}

	authorizationMetadataKey := cfg.AuthorizationHeader
	if len(authorizationMetadataKey) == 0 {
		clientMetadata, err := authMetadataClient.GetPublicClientConfig(ctx, &service.PublicClientAuthConfigRequest{})
		if err != nil {
			return fmt.Errorf("failed to fetch client metadata. Error: %v", err)
		}
		authorizationMetadataKey = clientMetadata.GetAuthorizationMetadataKey()
	}

	tokenSource, err := tokenSourceProvider.GetTokenSource(ctx)
	if err != nil {
		return fmt.Errorf("failed to get token source. Error: %w", err)
	}

	o.authorizationMetadataKey = authorizationMetadataKey
	o.tokenSource = tokenSource

	return nil
}

func (o *OauthMetadataProvider) GetOauthMetadata(cfg *Config, tokenCache cache.TokenCache, proxyCredentialsFuture *PerRPCCredentialsFuture) error {
	// Ensure loadTokenRelated() is only executed once
	var err error
	o.once.Do(func() {
		err = o.getTokenSourceAndMetadata(cfg, tokenCache, proxyCredentialsFuture)
		if err != nil {
			logger.Errorf(context.Background(), "Failed to load token related config. Error: %v", err)
		}
		logger.Debugf(context.Background(), "Successfully loaded token related metadata")
	})
	if err != nil {
		return err
	}
	return nil
}

// NewAuthInterceptor creates a new grpc.UnaryClientInterceptor that forwards the grpc call and inspects the error.
// It will first invoke the grpc pipeline (to proceed with the request) with no modifications. It's expected for the grpc
// pipeline to already have a grpc.WithPerRPCCredentials() DialOption. If the perRPCCredentials has already been initialized,
// it'll take care of refreshing when tokens expire... etc.
// If the first invocation succeeds (either due to grpc.PerRPCCredentials setting the right tokens or the server not
// requiring authentication), the interceptor will be no-op.
// If the first invocation fails with an auth error, this interceptor will then attempt to establish a token source once
// more. It'll fail hard if it couldn't do so (i.e. it will no longer attempt to send an unauthenticated request). Once
// a token source has been created, it'll invoke the grpc pipeline again, this time the grpc.PerRPCCredentials should
// be able to find and acquire a valid AccessToken to annotate the request with.
func NewAuthInterceptor(cfg *Config, tokenCache cache.TokenCache, credentialsFuture *PerRPCCredentialsFuture, proxyCredentialsFuture *PerRPCCredentialsFuture) grpc.UnaryClientInterceptor {

	oauthMetadataProvider := OauthMetadataProvider{
		once: sync.Once{},
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = setHTTPClientContext(ctx, cfg, proxyCredentialsFuture)
		// If there is already a token in the cache (e.g. key-ring), we should use it immediately...
		t, _ := tokenCache.GetToken()
		if t != nil {

			err := oauthMetadataProvider.GetOauthMetadata(cfg, tokenCache, proxyCredentialsFuture)
			if err != nil {
				return err
			}
			authorizationMetadataKey := oauthMetadataProvider.authorizationMetadataKey
			if isValid := utils.Valid(t); isValid {
				err := MaterializeInMemoryCredentials(ctx, cfg, tokenCache, credentialsFuture, authorizationMetadataKey)
				if err != nil {
					return fmt.Errorf("failed to materialize credentials. Error: %v", err)
				}
			}
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			logger.Debugf(ctx, "Request failed due to [%v]. If it's an unauthenticated error, we will attempt to establish an authenticated context.", err)

			if st, ok := status.FromError(err); ok {
				// If the error we receive from executing the request expects
				if shouldAttemptToAuthenticate(st.Code()) {
					err := oauthMetadataProvider.GetOauthMetadata(cfg, tokenCache, proxyCredentialsFuture)
					if err != nil {
						return err
					}
					authorizationMetadataKey := oauthMetadataProvider.authorizationMetadataKey
					tokenSource := oauthMetadataProvider.tokenSource
					err = func() error {
						if !tokenCache.TryLock() {
							tokenCache.CondWait()
							return nil
						}
						defer tokenCache.Unlock()
						_, err := tokenCache.PurgeIfEquals(t)
						if err != nil && !errors.Is(err, cache.ErrNotFound) {
							logger.Errorf(ctx, "Failed to purge cache. Error [%v]", err)
							return fmt.Errorf("failed to purge cache. Error: %w", err)
						}

						logger.Debugf(ctx, "Request failed due to [%v]. Attempting to establish an authenticated connection and trying again.", st.Code())
						newErr := MaterializeCredentials(tokenSource, cfg, authorizationMetadataKey, credentialsFuture)
						if newErr != nil {
							errString := fmt.Sprintf("authentication error! Original Error: %v, Auth Error: %v", err, newErr)
							logger.Errorf(ctx, errString)
							return fmt.Errorf(errString) //nolint
						}

						tokenCache.CondBroadcast()
						return nil
					}()

					if err != nil {
						return err
					}

					return invoker(ctx, method, req, reply, cc, opts...)
				}
			}
		}

		return err
	}
}

func NewProxyAuthInterceptor(cfg *Config, proxyCredentialsFuture *PerRPCCredentialsFuture) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			newErr := MaterializeProxyAuthCredentials(ctx, cfg, proxyCredentialsFuture)
			if newErr != nil {
				return fmt.Errorf("proxy authorization error! Original Error: %v, Proxy Auth Error: %w", err, newErr)
			}
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		return err
	}
}
