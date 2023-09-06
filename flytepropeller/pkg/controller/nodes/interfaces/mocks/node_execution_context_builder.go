// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	executors "github.com/flyteorg/flytepropeller/pkg/controller/executors"
	interfaces "github.com/flyteorg/flytepropeller/pkg/controller/nodes/interfaces"

	mock "github.com/stretchr/testify/mock"
)

// NodeExecutionContextBuilder is an autogenerated mock type for the NodeExecutionContextBuilder type
type NodeExecutionContextBuilder struct {
	mock.Mock
}

type NodeExecutionContextBuilder_BuildNodeExecutionContext struct {
	*mock.Call
}

func (_m NodeExecutionContextBuilder_BuildNodeExecutionContext) Return(_a0 interfaces.NodeExecutionContext, _a1 error) *NodeExecutionContextBuilder_BuildNodeExecutionContext {
	return &NodeExecutionContextBuilder_BuildNodeExecutionContext{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *NodeExecutionContextBuilder) OnBuildNodeExecutionContext(ctx context.Context, executionContext executors.ExecutionContext, nl executors.NodeLookup, currentNodeID string) *NodeExecutionContextBuilder_BuildNodeExecutionContext {
	c_call := _m.On("BuildNodeExecutionContext", ctx, executionContext, nl, currentNodeID)
	return &NodeExecutionContextBuilder_BuildNodeExecutionContext{Call: c_call}
}

func (_m *NodeExecutionContextBuilder) OnBuildNodeExecutionContextMatch(matchers ...interface{}) *NodeExecutionContextBuilder_BuildNodeExecutionContext {
	c_call := _m.On("BuildNodeExecutionContext", matchers...)
	return &NodeExecutionContextBuilder_BuildNodeExecutionContext{Call: c_call}
}

// BuildNodeExecutionContext provides a mock function with given fields: ctx, executionContext, nl, currentNodeID
func (_m *NodeExecutionContextBuilder) BuildNodeExecutionContext(ctx context.Context, executionContext executors.ExecutionContext, nl executors.NodeLookup, currentNodeID string) (interfaces.NodeExecutionContext, error) {
	ret := _m.Called(ctx, executionContext, nl, currentNodeID)

	var r0 interfaces.NodeExecutionContext
	if rf, ok := ret.Get(0).(func(context.Context, executors.ExecutionContext, executors.NodeLookup, string) interfaces.NodeExecutionContext); ok {
		r0 = rf(ctx, executionContext, nl, currentNodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.NodeExecutionContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, executors.ExecutionContext, executors.NodeLookup, string) error); ok {
		r1 = rf(ctx, executionContext, nl, currentNodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
