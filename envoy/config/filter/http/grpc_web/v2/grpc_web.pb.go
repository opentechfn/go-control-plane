// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_web/v2/grpc_web.proto

package envoy_config_filter_http_grpc_web_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GrpcWeb struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcWeb) Reset()         { *m = GrpcWeb{} }
func (m *GrpcWeb) String() string { return proto.CompactTextString(m) }
func (*GrpcWeb) ProtoMessage()    {}
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fe810f98c50c033, []int{0}
}

func (m *GrpcWeb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcWeb.Unmarshal(m, b)
}
func (m *GrpcWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcWeb.Marshal(b, m, deterministic)
}
func (m *GrpcWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcWeb.Merge(m, src)
}
func (m *GrpcWeb) XXX_Size() int {
	return xxx_messageInfo_GrpcWeb.Size(m)
}
func (m *GrpcWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcWeb.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcWeb proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GrpcWeb)(nil), "envoy.config.filter.http.grpc_web.v2.GrpcWeb")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_web/v2/grpc_web.proto", fileDescriptor_8fe810f98c50c033)
}

var fileDescriptor_8fe810f98c50c033 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x4f, 0x4d, 0xd2, 0x2f, 0x33,
	0x82, 0xb3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0x9a, 0xf4, 0x20, 0x9a, 0xf4,
	0x20, 0x9a, 0xf4, 0x40, 0x9a, 0xf4, 0xe0, 0x0a, 0xcb, 0x8c, 0xa4, 0xe4, 0x4a, 0x53, 0x0a, 0x12,
	0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x73, 0x33, 0xd3,
	0x8b, 0x12, 0x4b, 0x52, 0x21, 0xa6, 0x48, 0xc9, 0x62, 0xc8, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16,
	0x43, 0xa4, 0x95, 0x38, 0xb9, 0xd8, 0xdd, 0x8b, 0x0a, 0x92, 0xc3, 0x53, 0x93, 0x9c, 0x6a, 0x3f,
	0xcd, 0xf8, 0xd7, 0xcf, 0xaa, 0x2d, 0xa4, 0x09, 0xb1, 0x37, 0xb5, 0xa2, 0x24, 0x35, 0xaf, 0x18,
	0xa4, 0x03, 0x6a, 0x77, 0x31, 0xba, 0xe5, 0xc6, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63, 0x12, 0x60,
	0xe4, 0x32, 0xca, 0xcc, 0xd7, 0x03, 0xeb, 0x2a, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xc6, 0xe1,
	0x4e, 0x3c, 0x50, 0x7b, 0x03, 0x40, 0xee, 0x08, 0x60, 0x4c, 0x62, 0x03, 0x3b, 0xc8, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0x5f, 0x26, 0x22, 0x2c, 0x01, 0x00, 0x00,
}
