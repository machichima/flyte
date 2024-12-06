// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/plugins/spark.proto

package plugins

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SparkApplication_Type int32

const (
	SparkApplication_PYTHON SparkApplication_Type = 0
	SparkApplication_JAVA   SparkApplication_Type = 1
	SparkApplication_SCALA  SparkApplication_Type = 2
	SparkApplication_R      SparkApplication_Type = 3
)

// Enum value maps for SparkApplication_Type.
var (
	SparkApplication_Type_name = map[int32]string{
		0: "PYTHON",
		1: "JAVA",
		2: "SCALA",
		3: "R",
	}
	SparkApplication_Type_value = map[string]int32{
		"PYTHON": 0,
		"JAVA":   1,
		"SCALA":  2,
		"R":      3,
	}
)

func (x SparkApplication_Type) Enum() *SparkApplication_Type {
	p := new(SparkApplication_Type)
	*p = x
	return p
}

func (x SparkApplication_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SparkApplication_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_plugins_spark_proto_enumTypes[0].Descriptor()
}

func (SparkApplication_Type) Type() protoreflect.EnumType {
	return &file_flyteidl_plugins_spark_proto_enumTypes[0]
}

func (x SparkApplication_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SparkApplication_Type.Descriptor instead.
func (SparkApplication_Type) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_plugins_spark_proto_rawDescGZIP(), []int{0, 0}
}

type SparkApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SparkApplication) Reset() {
	*x = SparkApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_spark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparkApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparkApplication) ProtoMessage() {}

func (x *SparkApplication) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_spark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparkApplication.ProtoReflect.Descriptor instead.
func (*SparkApplication) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_spark_proto_rawDescGZIP(), []int{0}
}

// Custom Proto for Spark Plugin.
type SparkJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationType     SparkApplication_Type `protobuf:"varint,1,opt,name=applicationType,proto3,enum=flyteidl.plugins.SparkApplication_Type" json:"applicationType,omitempty"`
	MainApplicationFile string                `protobuf:"bytes,2,opt,name=mainApplicationFile,proto3" json:"mainApplicationFile,omitempty"`
	MainClass           string                `protobuf:"bytes,3,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	SparkConf           map[string]string     `protobuf:"bytes,4,rep,name=sparkConf,proto3" json:"sparkConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HadoopConf          map[string]string     `protobuf:"bytes,5,rep,name=hadoopConf,proto3" json:"hadoopConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecutorPath        string                `protobuf:"bytes,6,opt,name=executorPath,proto3" json:"executorPath,omitempty"` // Executor path for Python jobs.
	// Databricks job configuration.
	// Config structure can be found here. https://docs.databricks.com/dev-tools/api/2.0/jobs.html#request-structure.
	DatabricksConf *structpb.Struct `protobuf:"bytes,7,opt,name=databricksConf,proto3" json:"databricksConf,omitempty"`
	// Databricks access token. https://docs.databricks.com/dev-tools/api/latest/authentication.html
	// This token can be set in either flytepropeller or flytekit.
	DatabricksToken string `protobuf:"bytes,8,opt,name=databricksToken,proto3" json:"databricksToken,omitempty"`
	// Domain name of your deployment. Use the form <account>.cloud.databricks.com.
	// This instance name can be set in either flytepropeller or flytekit.
	DatabricksInstance string `protobuf:"bytes,9,opt,name=databricksInstance,proto3" json:"databricksInstance,omitempty"`

	DriverPod *core.K8SPod `protobuf:"bytes,10,opt,name=driverPod,json=driverPod,proto3" json:"driverPod,omitempty"`
	ExecutorPod *core.K8SPod `protobuf:"bytes,11,opt,name=executorPod,json=executorPod,proto3" json:"executorPod,omitempty"`
}

func (x *SparkJob) Reset() {
	*x = SparkJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_spark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparkJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparkJob) ProtoMessage() {}

func (x *SparkJob) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_spark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparkJob.ProtoReflect.Descriptor instead.
func (*SparkJob) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_spark_proto_rawDescGZIP(), []int{1}
}

func (x *SparkJob) GetApplicationType() SparkApplication_Type {
	if x != nil {
		return x.ApplicationType
	}
	return SparkApplication_PYTHON
}

func (x *SparkJob) GetMainApplicationFile() string {
	if x != nil {
		return x.MainApplicationFile
	}
	return ""
}

func (x *SparkJob) GetMainClass() string {
	if x != nil {
		return x.MainClass
	}
	return ""
}

func (x *SparkJob) GetSparkConf() map[string]string {
	if x != nil {
		return x.SparkConf
	}
	return nil
}

func (x *SparkJob) GetHadoopConf() map[string]string {
	if x != nil {
		return x.HadoopConf
	}
	return nil
}

func (x *SparkJob) GetExecutorPath() string {
	if x != nil {
		return x.ExecutorPath
	}
	return ""
}

func (x *SparkJob) GetDatabricksConf() *structpb.Struct {
	if x != nil {
		return x.DatabricksConf
	}
	return nil
}

func (x *SparkJob) GetDatabricksToken() string {
	if x != nil {
		return x.DatabricksToken
	}
	return ""
}

func (x *SparkJob) GetDatabricksInstance() string {
	if x != nil {
		return x.DatabricksInstance
	}
	return ""
}

var File_flyteidl_plugins_spark_proto protoreflect.FileDescriptor

var file_flyteidl_plugins_spark_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42,
	0x0a, 0x10, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59,
	0x54, 0x48, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x52,
	0x10, 0x03, 0x22, 0xfe, 0x04, 0x0a, 0x08, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x12,
	0x51, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x72,
	0x6b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x4a, 0x6f,
	0x62, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x4a, 0x0a, 0x0a, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x2e, 0x48, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3f, 0x0a, 0x0e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x28, 0x0a, 0x0f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x6b, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43,
	0x6f, 0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x48, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0xc2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x42, 0x0a, 0x53, 0x70,
	0x61, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x58, 0xaa,
	0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0xca, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0xe2, 0x02, 0x1c, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a,
	0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_plugins_spark_proto_rawDescOnce sync.Once
	file_flyteidl_plugins_spark_proto_rawDescData = file_flyteidl_plugins_spark_proto_rawDesc
)

func file_flyteidl_plugins_spark_proto_rawDescGZIP() []byte {
	file_flyteidl_plugins_spark_proto_rawDescOnce.Do(func() {
		file_flyteidl_plugins_spark_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_plugins_spark_proto_rawDescData)
	})
	return file_flyteidl_plugins_spark_proto_rawDescData
}

var file_flyteidl_plugins_spark_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyteidl_plugins_spark_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flyteidl_plugins_spark_proto_goTypes = []interface{}{
	(SparkApplication_Type)(0), // 0: flyteidl.plugins.SparkApplication.Type
	(*SparkApplication)(nil),   // 1: flyteidl.plugins.SparkApplication
	(*SparkJob)(nil),           // 2: flyteidl.plugins.SparkJob
	nil,                        // 3: flyteidl.plugins.SparkJob.SparkConfEntry
	nil,                        // 4: flyteidl.plugins.SparkJob.HadoopConfEntry
	(*structpb.Struct)(nil),    // 5: google.protobuf.Struct
}
var file_flyteidl_plugins_spark_proto_depIdxs = []int32{
	0, // 0: flyteidl.plugins.SparkJob.applicationType:type_name -> flyteidl.plugins.SparkApplication.Type
	3, // 1: flyteidl.plugins.SparkJob.sparkConf:type_name -> flyteidl.plugins.SparkJob.SparkConfEntry
	4, // 2: flyteidl.plugins.SparkJob.hadoopConf:type_name -> flyteidl.plugins.SparkJob.HadoopConfEntry
	5, // 3: flyteidl.plugins.SparkJob.databricksConf:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flyteidl_plugins_spark_proto_init() }
func file_flyteidl_plugins_spark_proto_init() {
	if File_flyteidl_plugins_spark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_plugins_spark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparkApplication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_plugins_spark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparkJob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_plugins_spark_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_plugins_spark_proto_goTypes,
		DependencyIndexes: file_flyteidl_plugins_spark_proto_depIdxs,
		EnumInfos:         file_flyteidl_plugins_spark_proto_enumTypes,
		MessageInfos:      file_flyteidl_plugins_spark_proto_msgTypes,
	}.Build()
	File_flyteidl_plugins_spark_proto = out.File
	file_flyteidl_plugins_spark_proto_rawDesc = nil
	file_flyteidl_plugins_spark_proto_goTypes = nil
	file_flyteidl_plugins_spark_proto_depIdxs = nil
}
