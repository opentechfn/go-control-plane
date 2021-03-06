// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/lightstep.proto

package envoy_config_trace_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type LightstepConfig_PropagationMode int32

const (
	LightstepConfig_ENVOY         LightstepConfig_PropagationMode = 0
	LightstepConfig_LIGHTSTEP     LightstepConfig_PropagationMode = 1
	LightstepConfig_B3            LightstepConfig_PropagationMode = 2
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

var LightstepConfig_PropagationMode_name = map[int32]string{
	0: "ENVOY",
	1: "LIGHTSTEP",
	2: "B3",
	3: "TRACE_CONTEXT",
}

var LightstepConfig_PropagationMode_value = map[string]int32{
	"ENVOY":         0,
	"LIGHTSTEP":     1,
	"B3":            2,
	"TRACE_CONTEXT": 3,
}

func (x LightstepConfig_PropagationMode) String() string {
	return proto.EnumName(LightstepConfig_PropagationMode_name, int32(x))
}

func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa3d42c7fd998072, []int{0, 0}
}

type LightstepConfig struct {
	CollectorCluster     string                            `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	AccessTokenFile      string                            `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	PropagationModes     []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.config.trace.v2.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa3d42c7fd998072, []int{0}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

func (m *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if m != nil {
		return m.PropagationModes
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v2.LightstepConfig_PropagationMode", LightstepConfig_PropagationMode_name, LightstepConfig_PropagationMode_value)
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v2/lightstep.proto", fileDescriptor_fa3d42c7fd998072)
}

var fileDescriptor_fa3d42c7fd998072 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x6a, 0xea, 0x30,
	0x18, 0x80, 0x4f, 0xea, 0x51, 0x31, 0xa0, 0xa6, 0x81, 0xc3, 0x11, 0xe1, 0x80, 0x78, 0x18, 0x78,
	0xd5, 0x82, 0x8e, 0xdd, 0x8e, 0xb5, 0xd4, 0x6d, 0xe0, 0xb4, 0xb8, 0x32, 0xb6, 0xab, 0x92, 0xb5,
	0xd1, 0x95, 0x75, 0xfd, 0x4b, 0x12, 0xcb, 0xbc, 0x1b, 0x7b, 0x84, 0x3d, 0xcd, 0xd8, 0x13, 0xec,
	0x76, 0xaf, 0xe3, 0xd5, 0xb0, 0x45, 0x87, 0xe2, 0x5d, 0xe0, 0xff, 0xbe, 0xe4, 0x4b, 0x82, 0x8f,
	0x78, 0x92, 0xc1, 0xd2, 0x0c, 0x20, 0x99, 0x45, 0x73, 0x53, 0x09, 0x16, 0x70, 0x33, 0xeb, 0x9b,
	0x71, 0x34, 0x7f, 0x50, 0x52, 0xf1, 0xd4, 0x48, 0x05, 0x28, 0xa0, 0x7f, 0x72, 0xcc, 0x28, 0x30,
	0x23, 0xc7, 0x8c, 0xac, 0xdf, 0xfe, 0xb7, 0x08, 0x53, 0x66, 0xb2, 0x24, 0x01, 0xc5, 0x54, 0x04,
	0x89, 0x34, 0xa5, 0x62, 0x6a, 0x21, 0x0b, 0xab, 0xfd, 0x37, 0x63, 0x71, 0x14, 0x32, 0xc5, 0xcd,
	0xcd, 0xa2, 0x18, 0x74, 0xdf, 0x35, 0xdc, 0x1c, 0x6d, 0x8e, 0xb0, 0xf3, 0x4d, 0xe9, 0x31, 0xd6,
	0x03, 0x88, 0x63, 0x1e, 0x28, 0x10, 0x7e, 0x10, 0x2f, 0xa4, 0xe2, 0xa2, 0x85, 0x3a, 0xa8, 0x57,
	0xb3, 0xaa, 0x2b, 0xeb, 0xb7, 0xd0, 0x3a, 0x68, 0x4a, 0xb6, 0x84, 0x5d, 0x00, 0x74, 0x80, 0x75,
	0x16, 0x04, 0x5c, 0x4a, 0x5f, 0xc1, 0x23, 0x4f, 0xfc, 0x59, 0x14, 0xf3, 0x96, 0xb6, 0x6b, 0x35,
	0x0b, 0xc2, 0x5b, 0x03, 0xc3, 0x28, 0xe6, 0x54, 0x60, 0x3d, 0x15, 0x90, 0xb2, 0x79, 0x1e, 0xed,
	0x3f, 0x41, 0xc8, 0x65, 0xab, 0xd4, 0x29, 0xf5, 0x1a, 0xfd, 0x13, 0xe3, 0xe0, 0x4d, 0x8d, 0xbd,
	0x5a, 0xc3, 0xfd, 0xf1, 0xaf, 0x20, 0xe4, 0x56, 0x7d, 0x65, 0xe1, 0x37, 0x54, 0xed, 0x96, 0x5f,
	0x91, 0x46, 0xd0, 0x94, 0xa4, 0xbb, 0x73, 0xd9, 0x1d, 0xe2, 0xe6, 0x9e, 0x43, 0x6b, 0xb8, 0xec,
	0x8c, 0x6f, 0x26, 0x77, 0xe4, 0x17, 0xad, 0xe3, 0xda, 0xe8, 0xf2, 0xfc, 0xc2, 0xbb, 0xf6, 0x1c,
	0x97, 0x20, 0x5a, 0xc1, 0x9a, 0x35, 0x20, 0x1a, 0xd5, 0x71, 0xdd, 0x9b, 0x9e, 0xd9, 0x8e, 0x6f,
	0x4f, 0xc6, 0x9e, 0x73, 0xeb, 0x91, 0x92, 0x75, 0xfa, 0xf1, 0xf2, 0xf9, 0x55, 0xd1, 0x08, 0xc2,
	0xff, 0x23, 0x28, 0x62, 0x53, 0x01, 0xcf, 0xcb, 0xc3, 0xdd, 0x56, 0x63, 0x1b, 0xee, 0xae, 0x5f,
	0xde, 0x45, 0xf7, 0x95, 0xfc, 0x0b, 0x06, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0x38, 0x87,
	0xfb, 0xfa, 0x01, 0x00, 0x00,
}
