// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/trace/v2/trace_service.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/lyft/protoc-gen-validate/validate"
import trace "istio.io/gogo-genproto/opencensus/proto/trace"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StreamTracesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTracesResponse) Reset()         { *m = StreamTracesResponse{} }
func (m *StreamTracesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTracesResponse) ProtoMessage()    {}
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_service_5901afff3db92913, []int{0}
}
func (m *StreamTracesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamTracesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesResponse.Merge(dst, src)
}
func (m *StreamTracesResponse) XXX_Size() int {
	return m.Size()
}
func (m *StreamTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesResponse proto.InternalMessageInfo

type StreamTracesMessage struct {
	// Identifier data effectively is a structured metadata.
	// As a performance optimization this will only be sent in the first message
	// on the stream.
	Identifier *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// A list of Span entries
	Spans                []*trace.Span `protobuf:"bytes,2,rep,name=spans" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StreamTracesMessage) Reset()         { *m = StreamTracesMessage{} }
func (m *StreamTracesMessage) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage) ProtoMessage()    {}
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_service_5901afff3db92913, []int{1}
}
func (m *StreamTracesMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamTracesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamTracesMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamTracesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage.Merge(dst, src)
}
func (m *StreamTracesMessage) XXX_Size() int {
	return m.Size()
}
func (m *StreamTracesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage proto.InternalMessageInfo

func (m *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTracesMessage) GetSpans() []*trace.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage_Identifier) Reset()         { *m = StreamTracesMessage_Identifier{} }
func (m *StreamTracesMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage_Identifier) ProtoMessage()    {}
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_service_5901afff3db92913, []int{1, 0}
}
func (m *StreamTracesMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamTracesMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamTracesMessage_Identifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamTracesMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage_Identifier.Merge(dst, src)
}
func (m *StreamTracesMessage_Identifier) XXX_Size() int {
	return m.Size()
}
func (m *StreamTracesMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage_Identifier proto.InternalMessageInfo

func (m *StreamTracesMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTracesResponse)(nil), "envoy.service.trace.v2.StreamTracesResponse")
	proto.RegisterType((*StreamTracesMessage)(nil), "envoy.service.trace.v2.StreamTracesMessage")
	proto.RegisterType((*StreamTracesMessage_Identifier)(nil), "envoy.service.trace.v2.StreamTracesMessage.Identifier")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TraceService service

type TraceServiceClient interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v2.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TraceService service

type TraceServiceServer interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(TraceService_StreamTracesServer) error
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v2/trace_service.proto",
}

func (m *StreamTracesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamTracesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamTracesMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamTracesMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identifier != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Identifier.Size()))
		n1, err := m.Identifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Spans) > 0 {
		for _, msg := range m.Spans {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTraceService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StreamTracesMessage_Identifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamTracesMessage_Identifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Node.Size()))
		n2, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTraceService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StreamTracesResponse) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamTracesMessage) Size() (n int) {
	var l int
	_ = l
	if m.Identifier != nil {
		l = m.Identifier.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovTraceService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StreamTracesMessage_Identifier) Size() (n int) {
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTraceService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTraceService(x uint64) (n int) {
	return sovTraceService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamTracesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamTracesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamTracesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamTracesMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamTracesMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamTracesMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Identifier == nil {
				m.Identifier = &StreamTracesMessage_Identifier{}
			}
			if err := m.Identifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, &trace.Span{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamTracesMessage_Identifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Identifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &core.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTraceService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTraceService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTraceService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTraceService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTraceService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/trace/v2/trace_service.proto", fileDescriptor_trace_service_5901afff3db92913)
}

var fileDescriptor_trace_service_5901afff3db92913 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0x4d, 0x6b, 0x1d, 0xd2, 0x0e, 0x72, 0x95, 0xb6, 0x1c, 0xa5, 0x94, 0x4e, 0x45, 0x25,
	0x07, 0x11, 0x74, 0x2f, 0x38, 0x38, 0xe8, 0x70, 0x15, 0x07, 0x17, 0x49, 0xef, 0x9e, 0x25, 0xd0,
	0xe6, 0x85, 0x4b, 0x3c, 0x70, 0x70, 0x17, 0xff, 0x24, 0x27, 0x47, 0x47, 0xff, 0x04, 0xe9, 0xa4,
	0xff, 0x85, 0xdc, 0x25, 0xad, 0x37, 0x74, 0xd0, 0x2d, 0x79, 0xef, 0xfb, 0x3e, 0x7e, 0xdf, 0xa3,
	0x87, 0xa0, 0x72, 0x7c, 0x8c, 0x0c, 0x64, 0xb9, 0x4c, 0x20, 0xb2, 0x99, 0x48, 0x20, 0xca, 0xb9,
	0x7b, 0xdc, 0xf9, 0x31, 0xd3, 0x19, 0x5a, 0x0c, 0x3a, 0xa5, 0x96, 0xad, 0x87, 0xa5, 0x84, 0xe5,
	0x3c, 0xec, 0xbb, 0x0c, 0xa1, 0x65, 0xe1, 0x4c, 0x30, 0x83, 0x68, 0x26, 0x8c, 0x77, 0x85, 0x4d,
	0xa7, 0x73, 0x9f, 0xfe, 0x1c, 0x71, 0xbe, 0x80, 0x52, 0x2b, 0x94, 0x42, 0x2b, 0xac, 0x44, 0x65,
	0xfc, 0xb6, 0x9b, 0x8b, 0x85, 0x4c, 0x85, 0x85, 0x68, 0xfd, 0x70, 0x8b, 0x51, 0x87, 0x1e, 0x4c,
	0x6d, 0x06, 0x62, 0x79, 0x5d, 0x64, 0x99, 0x18, 0x8c, 0x46, 0x65, 0x60, 0xf4, 0x45, 0x68, 0xbb,
	0xba, 0xb8, 0x04, 0x63, 0xc4, 0x1c, 0x82, 0x1b, 0x4a, 0x65, 0x0a, 0xca, 0xca, 0x7b, 0x09, 0x59,
	0x8f, 0x0c, 0xc9, 0xb8, 0xc9, 0x4f, 0xd9, 0x76, 0x7c, 0xb6, 0x25, 0x80, 0x5d, 0x6c, 0xdc, 0x71,
	0x25, 0x29, 0xe0, 0xb4, 0x61, 0xb4, 0x50, 0xa6, 0x57, 0x1b, 0xd6, 0xc7, 0x4d, 0xde, 0x67, 0xa8,
	0x41, 0x25, 0xa0, 0xcc, 0x83, 0xaf, 0xe0, 0x53, 0xa7, 0x5a, 0xa8, 0xd8, 0x49, 0xc3, 0x73, 0x4a,
	0x7f, 0xd3, 0x82, 0x33, 0xba, 0xab, 0x30, 0x05, 0xcf, 0xd4, 0xf5, 0x4c, 0x42, 0xcb, 0x82, 0xa4,
	0x38, 0x1d, 0xbb, 0xc2, 0x14, 0x26, 0xf4, 0xf5, 0xfb, 0xad, 0xde, 0x78, 0x21, 0xb5, 0x7d, 0x12,
	0x97, 0x06, 0xfe, 0x44, 0x5b, 0x25, 0xe2, 0xd4, 0xe1, 0x07, 0x4b, 0xda, 0xaa, 0x82, 0x07, 0x47,
	0xff, 0xa8, 0x17, 0x1e, 0xff, 0x45, 0xbc, 0xb9, 0xf2, 0xce, 0x98, 0x4c, 0xda, 0xef, 0xab, 0x01,
	0xf9, 0x58, 0x0d, 0xc8, 0xe7, 0x6a, 0x40, 0x6e, 0x6b, 0x39, 0x7f, 0x26, 0x64, 0xb6, 0x57, 0x76,
	0x3e, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x43, 0xf2, 0x34, 0x31, 0x45, 0x02, 0x00, 0x00,
}
