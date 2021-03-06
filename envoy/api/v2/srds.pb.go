// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/srds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SrdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrdsDummy) Reset()         { *m = SrdsDummy{} }
func (m *SrdsDummy) String() string { return proto.CompactTextString(m) }
func (*SrdsDummy) ProtoMessage()    {}
func (*SrdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_92f394721ede65e9, []int{0}
}

func (m *SrdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrdsDummy.Unmarshal(m, b)
}
func (m *SrdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrdsDummy.Marshal(b, m, deterministic)
}
func (m *SrdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrdsDummy.Merge(m, src)
}
func (m *SrdsDummy) XXX_Size() int {
	return xxx_messageInfo_SrdsDummy.Size(m)
}
func (m *SrdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SrdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SrdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SrdsDummy)(nil), "envoy.api.v2.SrdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/srds.proto", fileDescriptor_92f394721ede65e9) }

var fileDescriptor_92f394721ede65e9 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0x4e, 0xe3, 0x40,
	0x14, 0x85, 0x77, 0x12, 0x25, 0xab, 0x9d, 0xdd, 0x62, 0xe3, 0x62, 0x17, 0x99, 0xfc, 0xa0, 0x00,
	0x02, 0x21, 0x62, 0xa3, 0xa4, 0x4b, 0x19, 0x22, 0xea, 0x28, 0x2e, 0x29, 0xd0, 0x60, 0x0f, 0x66,
	0xa4, 0xd8, 0xd7, 0xcc, 0x8c, 0x2d, 0xdc, 0xa1, 0x54, 0x88, 0x86, 0x02, 0x81, 0x78, 0x00, 0x9e,
	0x82, 0x27, 0xa0, 0x45, 0xbc, 0x02, 0x4f, 0x00, 0x3d, 0x42, 0x19, 0xc7, 0x90, 0x21, 0x82, 0x8a,
	0xfa, 0x7c, 0xe7, 0xfe, 0x9d, 0x8b, 0xff, 0xd3, 0x30, 0x81, 0xd4, 0x26, 0x11, 0xb3, 0x93, 0xb6,
	0x2d, 0xb8, 0x27, 0xac, 0x88, 0x83, 0x04, 0xe3, 0x8f, 0x12, 0x2c, 0x12, 0x31, 0x2b, 0x69, 0x9b,
	0x55, 0x0d, 0xf3, 0x98, 0x70, 0x21, 0xa1, 0x3c, 0xcd, 0x58, 0xb3, 0xea, 0x03, 0xf8, 0x23, 0xaa,
	0x64, 0x12, 0x86, 0x20, 0x89, 0x64, 0x10, 0x4e, 0x2b, 0x99, 0x4b, 0x53, 0xef, 0xbb, 0x60, 0x73,
	0x2a, 0x20, 0xe6, 0x2e, 0x9d, 0x12, 0xf5, 0xd8, 0x8b, 0x88, 0x06, 0x04, 0xcc, 0xe7, 0x44, 0xe6,
	0x7a, 0x6d, 0x4e, 0x17, 0x92, 0xc8, 0x38, 0x6f, 0xd0, 0xd0, 0x77, 0x70, 0x21, 0xa2, 0xde, 0x1e,
	0x87, 0x38, 0xf7, 0x37, 0x7f, 0xe3, 0x5f, 0x0e, 0xf7, 0x44, 0x3f, 0x0e, 0x82, 0xb4, 0x7d, 0x59,
	0xc4, 0x55, 0x47, 0x31, 0xc3, 0x09, 0x22, 0xfa, 0xf9, 0x32, 0x0e, 0xe5, 0x09, 0x73, 0xa9, 0xb1,
	0x8b, 0x0d, 0x47, 0x72, 0x4a, 0x82, 0x59, 0xca, 0xa8, 0x5b, 0xb3, 0x07, 0xb1, 0xde, 0x5c, 0x43,
	0x7a, 0x14, 0x53, 0x21, 0xcd, 0xc6, 0xa7, 0xba, 0x88, 0x20, 0x14, 0xb4, 0xf9, 0x63, 0x1d, 0x6d,
	0x21, 0xc3, 0xc3, 0x95, 0x3e, 0x1d, 0x49, 0xa2, 0xd5, 0x5e, 0xfe, 0xe0, 0x9d, 0x00, 0x73, 0x0d,
	0x56, 0xbe, 0x86, 0xb4, 0x2e, 0x63, 0x84, 0x2b, 0x3b, 0x54, 0xba, 0x87, 0xdf, 0xbb, 0xc2, 0xe6,
	0xf8, 0xe1, 0xf1, 0xa2, 0x50, 0x6b, 0x2e, 0x6a, 0x4f, 0xd0, 0xcd, 0x2e, 0xde, 0x52, 0x17, 0x17,
	0x0a, 0x29, 0x76, 0xd1, 0x86, 0xd9, 0x3a, 0xbb, 0xb9, 0x7a, 0xfe, 0xb9, 0x86, 0x57, 0xb5, 0xaa,
	0x33, 0x13, 0x6d, 0x43, 0x78, 0xc0, 0xfc, 0x98, 0xab, 0x4c, 0x7b, 0xc3, 0xa7, 0xeb, 0x97, 0xf3,
	0xd2, 0x82, 0xf1, 0x2f, 0xc3, 0x45, 0x96, 0x86, 0x95, 0xe5, 0x98, 0x74, 0x6e, 0x4f, 0xee, 0xee,
	0xcb, 0x85, 0xbf, 0x08, 0x9b, 0x0c, 0xb2, 0x39, 0x23, 0x0e, 0xc7, 0xa9, 0x36, 0x72, 0x4f, 0xc5,
	0x3c, 0x98, 0x64, 0x3e, 0x40, 0xa7, 0x08, 0x0d, 0x4a, 0xfb, 0x65, 0xf5, 0x01, 0x9d, 0xd7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xad, 0x06, 0xf9, 0x5a, 0xe8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScopedRoutesDiscoveryServiceClient is the client API for ScopedRoutesDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScopedRoutesDiscoveryServiceClient interface {
	StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error)
	DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error)
	FetchScopedRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type scopedRoutesDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewScopedRoutesDiscoveryServiceClient(cc *grpc.ClientConn) ScopedRoutesDiscoveryServiceClient {
	return &scopedRoutesDiscoveryServiceClient{cc}
}

func (c *scopedRoutesDiscoveryServiceClient) StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ScopedRoutesDiscoveryService/StreamScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceStreamScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ScopedRoutesDiscoveryService/DeltaScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceDeltaScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) FetchScopedRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ScopedRoutesDiscoveryService/FetchScopedRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopedRoutesDiscoveryServiceServer is the server API for ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceServer interface {
	StreamScopedRoutes(ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error
	DeltaScopedRoutes(ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error
	FetchScopedRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedScopedRoutesDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScopedRoutesDiscoveryServiceServer struct {
}

func (*UnimplementedScopedRoutesDiscoveryServiceServer) StreamScopedRoutes(srv ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) DeltaScopedRoutes(srv ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) FetchScopedRoutes(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchScopedRoutes not implemented")
}

func RegisterScopedRoutesDiscoveryServiceServer(s *grpc.Server, srv ScopedRoutesDiscoveryServiceServer) {
	s.RegisterService(&_ScopedRoutesDiscoveryService_serviceDesc, srv)
}

func _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).StreamScopedRoutes(&scopedRoutesDiscoveryServiceStreamScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).DeltaScopedRoutes(&scopedRoutesDiscoveryServiceDeltaScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ScopedRoutesDiscoveryService/FetchScopedRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopedRoutesDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ScopedRoutesDiscoveryService",
	HandlerType: (*ScopedRoutesDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchScopedRoutes",
			Handler:    _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/srds.proto",
}
