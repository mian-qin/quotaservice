// Code generated by protoc-gen-go.
// source: protos/config/configs.proto
// DO NOT EDIT!

/*
Package quotaservice_configs is a generated protocol buffer package.

It is generated from these files:
	protos/config/configs.proto

It has these top-level messages:
	ServiceConfig
	NamespaceConfig
	BucketConfig
*/
package quotaservice_configs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Representations of configuration elements, for persisting and sharing across nodes.
type ServiceConfig struct {
	GlobalDefaultBucket *BucketConfig               `protobuf:"bytes,1,opt,name=global_default_bucket,json=globalDefaultBucket" json:"global_default_bucket,omitempty" yaml:"global_default_bucket"`
	Namespaces          map[string]*NamespaceConfig `protobuf:"bytes,2,rep,name=namespaces" json:"namespaces,omitempty" yaml:"namespaces" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Version             int32                       `protobuf:"varint,3,opt,name=version" json:"version,omitempty" yaml:"version"`
}

func (m *ServiceConfig) Reset()                    { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()               {}
func (*ServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceConfig) GetGlobalDefaultBucket() *BucketConfig {
	if m != nil {
		return m.GlobalDefaultBucket
	}
	return nil
}

func (m *ServiceConfig) GetNamespaces() map[string]*NamespaceConfig {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type NamespaceConfig struct {
	Name                  string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty" yaml:"name"`
	DefaultBucket         *BucketConfig            `protobuf:"bytes,2,opt,name=default_bucket,json=defaultBucket" json:"default_bucket,omitempty" yaml:"default_bucket"`
	DynamicBucketTemplate *BucketConfig            `protobuf:"bytes,3,opt,name=dynamic_bucket_template,json=dynamicBucketTemplate" json:"dynamic_bucket_template,omitempty" yaml:"dynamic_bucket_template"`
	MaxDynamicBuckets     int32                    `protobuf:"varint,4,opt,name=max_dynamic_buckets,json=maxDynamicBuckets" json:"max_dynamic_buckets,omitempty" yaml:"max_dynamic_buckets"`
	Buckets               map[string]*BucketConfig `protobuf:"bytes,5,rep,name=buckets" json:"buckets,omitempty" yaml:"buckets" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NamespaceConfig) Reset()                    { *m = NamespaceConfig{} }
func (m *NamespaceConfig) String() string            { return proto.CompactTextString(m) }
func (*NamespaceConfig) ProtoMessage()               {}
func (*NamespaceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NamespaceConfig) GetDefaultBucket() *BucketConfig {
	if m != nil {
		return m.DefaultBucket
	}
	return nil
}

func (m *NamespaceConfig) GetDynamicBucketTemplate() *BucketConfig {
	if m != nil {
		return m.DynamicBucketTemplate
	}
	return nil
}

func (m *NamespaceConfig) GetBuckets() map[string]*BucketConfig {
	if m != nil {
		return m.Buckets
	}
	return nil
}

type BucketConfig struct {
	Name                string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty" yaml:"name"`
	Namespace           string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty" yaml:"namespace"`
	Size                int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty" yaml:"size"`
	FillRate            int64  `protobuf:"varint,4,opt,name=fill_rate,json=fillRate" json:"fill_rate,omitempty" yaml:"fill_rate"`
	WaitTimeoutMillis   int64  `protobuf:"varint,5,opt,name=wait_timeout_millis,json=waitTimeoutMillis" json:"wait_timeout_millis,omitempty" yaml:"wait_timeout_millis"`
	MaxIdleMillis       int64  `protobuf:"varint,6,opt,name=max_idle_millis,json=maxIdleMillis" json:"max_idle_millis,omitempty" yaml:"max_idle_millis"`
	MaxDebtMillis       int64  `protobuf:"varint,7,opt,name=max_debt_millis,json=maxDebtMillis" json:"max_debt_millis,omitempty" yaml:"max_debt_millis"`
	MaxTokensPerRequest int64  `protobuf:"varint,8,opt,name=max_tokens_per_request,json=maxTokensPerRequest" json:"max_tokens_per_request,omitempty" yaml:"max_tokens_per_request"`
}

func (m *BucketConfig) Reset()                    { *m = BucketConfig{} }
func (m *BucketConfig) String() string            { return proto.CompactTextString(m) }
func (*BucketConfig) ProtoMessage()               {}
func (*BucketConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "quotaservice.configs.ServiceConfig")
	proto.RegisterType((*NamespaceConfig)(nil), "quotaservice.configs.NamespaceConfig")
	proto.RegisterType((*BucketConfig)(nil), "quotaservice.configs.BucketConfig")
}

func init() { proto.RegisterFile("protos/config/configs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xec, 0xa4, 0x69, 0xa6, 0x84, 0xd2, 0x2d, 0x05, 0xab, 0xe5, 0x10, 0x45, 0x02, 0xe5,
	0x64, 0xa4, 0xe4, 0x52, 0xc1, 0x0d, 0xc2, 0xa1, 0x12, 0x20, 0xe4, 0x46, 0x1c, 0x38, 0x60, 0xad,
	0xed, 0x49, 0xb5, 0xca, 0xda, 0x4e, 0xbd, 0xeb, 0x90, 0xf0, 0x83, 0x38, 0xf0, 0x1b, 0xf8, 0x71,
	0x68, 0x3f, 0x6c, 0xec, 0x2a, 0x42, 0x39, 0x65, 0x33, 0xef, 0xcd, 0x9b, 0x99, 0xf7, 0xa2, 0xc0,
	0xd5, 0xba, 0xc8, 0x65, 0x2e, 0x5e, 0xc7, 0x79, 0xb6, 0x64, 0x77, 0xf6, 0x43, 0xf8, 0xba, 0x4a,
	0x9e, 0xde, 0x97, 0xb9, 0xa4, 0x02, 0x8b, 0x0d, 0x8b, 0xd1, 0xb7, 0xd8, 0xf8, 0x8f, 0x03, 0xc3,
	0x5b, 0x53, 0x7b, 0xaf, 0x4b, 0xe4, 0x2b, 0x5c, 0xdc, 0xf1, 0x3c, 0xa2, 0x3c, 0x4c, 0x70, 0x49,
	0x4b, 0x2e, 0xc3, 0xa8, 0x8c, 0x57, 0x28, 0xbd, 0xce, 0xa8, 0x33, 0x39, 0x99, 0x8e, 0xfd, 0x7d,
	0x3a, 0xfe, 0x3b, 0xcd, 0x31, 0x12, 0xc1, 0xb9, 0x11, 0x98, 0x9b, 0x7e, 0x03, 0x91, 0x5b, 0x80,
	0x8c, 0xa6, 0x28, 0xd6, 0x34, 0x46, 0xe1, 0x39, 0x23, 0x77, 0x72, 0x32, 0x9d, 0xed, 0x17, 0x6b,
	0x2d, 0xe4, 0x7f, 0xae, 0xbb, 0x3e, 0x64, 0xb2, 0xd8, 0x05, 0x0d, 0x19, 0xe2, 0x41, 0x7f, 0x83,
	0x85, 0x60, 0x79, 0xe6, 0xb9, 0xa3, 0xce, 0xa4, 0x17, 0x54, 0x5f, 0x2f, 0x13, 0x38, 0x7d, 0xd0,
	0x48, 0x9e, 0x80, 0xbb, 0xc2, 0x9d, 0xbe, 0x63, 0x10, 0xa8, 0x27, 0x79, 0x0b, 0xbd, 0x0d, 0xe5,
	0x25, 0x7a, 0x8e, 0xbe, 0xed, 0xe5, 0xfe, 0x75, 0x6a, 0x1d, 0x7b, 0x9e, 0xe9, 0x79, 0xe3, 0x5c,
	0x77, 0xc6, 0xbf, 0xdd, 0xc6, 0x18, 0x6b, 0x20, 0x81, 0xae, 0xda, 0xd0, 0xce, 0xd1, 0x6f, 0x72,
	0x03, 0x8f, 0x1f, 0xb8, 0xe9, 0x1c, 0xec, 0xe6, 0x30, 0x69, 0xf9, 0xf8, 0x0d, 0x9e, 0x27, 0xbb,
	0x8c, 0xa6, 0x2c, 0xb6, 0x52, 0xa1, 0xc4, 0x74, 0xcd, 0xa9, 0x44, 0x6d, 0xc1, 0x61, 0x9a, 0x17,
	0x56, 0xc2, 0x14, 0x17, 0x56, 0x80, 0xf8, 0x70, 0x9e, 0xd2, 0x6d, 0xd8, 0xd6, 0x17, 0x5e, 0x57,
	0x5b, 0x7b, 0x96, 0xd2, 0xed, 0xbc, 0xd9, 0x26, 0xc8, 0x47, 0xe8, 0x57, 0x9c, 0x9e, 0x0e, 0x74,
	0x7a, 0x90, 0x83, 0x76, 0x17, 0x9b, 0x67, 0x25, 0x71, 0xf9, 0x1d, 0x1e, 0x35, 0x81, 0x3d, 0x79,
	0x5d, 0xb7, 0xf3, 0x3a, 0xe4, 0xd2, 0x46, 0x58, 0xbf, 0x9c, 0x6a, 0xc0, 0x7f, 0x92, 0x7a, 0x01,
	0x83, 0xfa, 0xf7, 0xa5, 0xc7, 0x0c, 0x82, 0x7f, 0x05, 0xd5, 0x21, 0xd8, 0x4f, 0xe3, 0xb4, 0x1b,
	0xe8, 0x37, 0xb9, 0x82, 0xc1, 0x92, 0x71, 0x1e, 0x16, 0x2a, 0x82, 0xae, 0x06, 0x8e, 0x55, 0x21,
	0xb0, 0x8e, 0xfe, 0xa0, 0x4c, 0x86, 0x92, 0xa5, 0x98, 0x97, 0x32, 0x4c, 0x19, 0xe7, 0x4c, 0xb9,
	0xa5, 0x68, 0x67, 0x0a, 0x5a, 0x18, 0xe4, 0x93, 0x06, 0xc8, 0x2b, 0x38, 0x55, 0x09, 0xb0, 0x84,
	0x63, 0xc5, 0x3d, 0xd2, 0xdc, 0x61, 0x4a, 0xb7, 0x37, 0x09, 0xc7, 0x36, 0x2f, 0xc1, 0xa8, 0xd6,
	0xec, 0xd7, 0xbc, 0x39, 0x46, 0x95, 0xde, 0x0c, 0x9e, 0x29, 0x9e, 0xcc, 0x57, 0x98, 0x89, 0x70,
	0x8d, 0x45, 0x58, 0xe0, 0x7d, 0x89, 0x42, 0x7a, 0xc7, 0x9a, 0xae, 0xf2, 0x5e, 0x68, 0xf0, 0x0b,
	0x16, 0x81, 0x81, 0xa2, 0x23, 0xfd, 0x8f, 0x31, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x81,
	0x15, 0xa2, 0x50, 0x04, 0x00, 0x00,
}
