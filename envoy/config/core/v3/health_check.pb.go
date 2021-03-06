// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/health_check.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0}
}

type HealthCheck struct {
	Timeout                      *duration.Duration      `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval                     *duration.Duration      `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter                *duration.Duration      `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter               *duration.Duration      `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent        uint32                  `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold           *wrappers.UInt32Value   `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold             *wrappers.UInt32Value   `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort                      *wrappers.UInt32Value   `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection              *wrappers.BoolValue     `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	NoTrafficInterval            *duration.Duration      `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration      `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration      `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration      `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                  `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	EventService                 *EventServiceConfig     `protobuf:"bytes,22,opt,name=event_service,json=eventService,proto3" json:"event_service,omitempty"`
	AlwaysLogHealthCheckFailures bool                    `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	TlsOptions                   *HealthCheck_TlsOptions `protobuf:"bytes,21,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty"`
	TransportSocketMatchCriteria *_struct.Struct         `protobuf:"bytes,23,opt,name=transport_socket_match_criteria,json=transportSocketMatchCriteria,proto3" json:"transport_socket_match_criteria,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker        isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetEventService() *EventServiceConfig {
	if m != nil {
		return m.EventService
	}
	return nil
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

func (m *HealthCheck) GetTlsOptions() *HealthCheck_TlsOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

func (m *HealthCheck) GetTransportSocketMatchCriteria() *_struct.Struct {
	if m != nil {
		return m.TransportSocketMatchCriteria
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                             string               `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                             string               `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                             *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                          *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	RequestHeadersToAdd              []*HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove           []string             `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	ExpectedStatuses                 []*v3.Int64Range     `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType                  v3.CodecClientType   `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.v3.CodecClientType" json:"codec_client_type,omitempty"`
	ServiceNameMatcher               *v31.StringMatcher   `protobuf:"bytes,11,opt,name=service_name_matcher,json=serviceNameMatcher,proto3" json:"service_name_matcher,omitempty"`
	HiddenEnvoyDeprecatedServiceName string               `protobuf:"bytes,5,opt,name=hidden_envoy_deprecated_service_name,json=hiddenEnvoyDeprecatedServiceName,proto3" json:"hidden_envoy_deprecated_service_name,omitempty"` // Deprecated: Do not use.
	HiddenEnvoyDeprecatedUseHttp2    bool                 `protobuf:"varint,7,opt,name=hidden_envoy_deprecated_use_http2,json=hiddenEnvoyDeprecatedUseHttp2,proto3" json:"hidden_envoy_deprecated_use_http2,omitempty"`         // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*v3.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() v3.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return v3.CodecClientType_HTTP1
}

func (m *HealthCheck_HttpHealthCheck) GetServiceNameMatcher() *v31.StringMatcher {
	if m != nil {
		return m.ServiceNameMatcher
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedServiceName() string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedServiceName
	}
	return ""
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedUseHttp2() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseHttp2
	}
	return false
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	//	*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {
}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_CustomHealthCheck) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
		(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig)(nil),
	}
}

type HealthCheck_TlsOptions struct {
	AlpnProtocols        []string `protobuf:"bytes,1,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_TlsOptions) Reset()         { *m = HealthCheck_TlsOptions{} }
func (m *HealthCheck_TlsOptions) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TlsOptions) ProtoMessage()    {}
func (*HealthCheck_TlsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 6}
}

func (m *HealthCheck_TlsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TlsOptions.Unmarshal(m, b)
}
func (m *HealthCheck_TlsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TlsOptions.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TlsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TlsOptions.Merge(m, src)
}
func (m *HealthCheck_TlsOptions) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TlsOptions.Size(m)
}
func (m *HealthCheck_TlsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TlsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TlsOptions proto.InternalMessageInfo

func (m *HealthCheck_TlsOptions) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.core.v3.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.config.core.v3.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.CustomHealthCheck")
	proto.RegisterType((*HealthCheck_TlsOptions)(nil), "envoy.config.core.v3.HealthCheck.TlsOptions")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/health_check.proto", fileDescriptor_2ba40dc0efbf9537)
}

var fileDescriptor_2ba40dc0efbf9537 = []byte{
	// 1655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x97, 0xcf, 0x6f, 0x1b, 0xc7,
	0x15, 0xc7, 0xb9, 0x14, 0x2d, 0x92, 0x8f, 0x22, 0xb5, 0x1a, 0x59, 0xd6, 0x8a, 0x91, 0x6d, 0xda,
	0x75, 0x60, 0xc6, 0x70, 0xc9, 0x44, 0x4a, 0x5d, 0x54, 0x40, 0xd1, 0x68, 0x25, 0xd9, 0x54, 0x62,
	0xc9, 0xc2, 0x88, 0x8a, 0x51, 0xa4, 0xe8, 0x76, 0xb4, 0x3b, 0x22, 0xb7, 0x5e, 0xed, 0x6c, 0x67,
	0x87, 0x8c, 0x79, 0x0b, 0x7a, 0x2a, 0x7a, 0x0c, 0xd0, 0x1e, 0xfa, 0x27, 0xe4, 0x2f, 0x28, 0xda,
	0x4b, 0x2f, 0x05, 0xd2, 0x63, 0xfb, 0x2f, 0xf4, 0xd2, 0x53, 0x0f, 0x3d, 0x15, 0x3a, 0x15, 0x33,
	0xbb, 0x4b, 0x72, 0x49, 0x5a, 0x3f, 0x9a, 0x9b, 0x38, 0xef, 0x7d, 0x3f, 0xef, 0xcd, 0xec, 0xbc,
	0xf7, 0x46, 0xf0, 0x98, 0xfa, 0x7d, 0x36, 0x68, 0xda, 0xcc, 0x3f, 0x73, 0x3b, 0x4d, 0x9b, 0x71,
	0xda, 0xec, 0x6f, 0x36, 0xbb, 0x94, 0x78, 0xa2, 0x6b, 0xd9, 0x5d, 0x6a, 0xbf, 0x69, 0x04, 0x9c,
	0x09, 0x86, 0x6e, 0x2b, 0xc7, 0x46, 0xe4, 0xd8, 0x90, 0x8e, 0x8d, 0xfe, 0x66, 0xf5, 0xfe, 0x4c,
	0xf9, 0x29, 0x09, 0x69, 0x24, 0xab, 0x36, 0x67, 0x3a, 0xd0, 0x3e, 0xf5, 0x85, 0x15, 0x52, 0xde,
	0x77, 0x6d, 0x6a, 0xc5, 0xcc, 0x48, 0xf0, 0x30, 0x12, 0x88, 0x41, 0x40, 0x9b, 0xe7, 0x44, 0xd8,
	0x5d, 0xca, 0xa5, 0x22, 0x14, 0xdc, 0xf5, 0x13, 0x1f, 0x63, 0xcc, 0x47, 0x66, 0x2b, 0x44, 0x10,
	0x5b, 0xd6, 0xd2, 0x16, 0x4e, 0xfc, 0x4e, 0x92, 0xc9, 0x5a, 0x87, 0xb1, 0x8e, 0x47, 0x9b, 0xea,
	0xd7, 0x69, 0xef, 0xac, 0x49, 0xfc, 0x41, 0x6c, 0xba, 0x37, 0x69, 0x72, 0x7a, 0x9c, 0x08, 0x97,
	0xf9, 0xb1, 0x7d, 0x7d, 0xd2, 0x1e, 0x0a, 0xde, 0xb3, 0xc5, 0xbb, 0xd4, 0x5f, 0x72, 0x12, 0x04,
	0x94, 0x87, 0xb1, 0xfd, 0x7b, 0x51, 0x4e, 0xc4, 0xf7, 0x99, 0x50, 0xd4, 0xb0, 0xe9, 0xd0, 0x80,
	0x53, 0x7b, 0x3c, 0xc4, 0xdd, 0x9e, 0x13, 0x90, 0x94, 0x4f, 0x28, 0x88, 0xe8, 0x25, 0x8c, 0x07,
	0x53, 0xe6, 0x3e, 0xe5, 0xa1, 0xcb, 0xfc, 0xd1, 0xa1, 0xac, 0xf6, 0x89, 0xe7, 0x3a, 0x44, 0xd0,
	0x66, 0xf2, 0x47, 0x64, 0x78, 0xf8, 0xe7, 0x75, 0x28, 0xb5, 0xd4, 0x07, 0xdd, 0x91, 0xdf, 0x13,
	0xfd, 0x04, 0xf2, 0xc2, 0x3d, 0xa7, 0xac, 0x27, 0x0c, 0xad, 0xa6, 0xd5, 0x4b, 0x1b, 0x6b, 0x8d,
	0x68, 0x07, 0x8d, 0x64, 0x07, 0x8d, 0xdd, 0x78, 0xff, 0x26, 0x5c, 0x98, 0xf9, 0x6f, 0xb4, 0x5c,
	0x41, 0x7b, 0x92, 0xc1, 0x89, 0x0a, 0x6d, 0x43, 0xc1, 0xf5, 0x05, 0xe5, 0x7d, 0xe2, 0x19, 0xd9,
	0x9b, 0x10, 0x86, 0x32, 0xf4, 0x09, 0x54, 0x5c, 0xdf, 0x15, 0x2e, 0xf1, 0xac, 0x5f, 0xba, 0x42,
	0x50, 0x6e, 0xdc, 0xbe, 0x02, 0x84, 0xcb, 0xb1, 0xe0, 0x53, 0xe5, 0x8f, 0x4c, 0x58, 0x4c, 0x68,
	0x09, 0x62, 0xee, 0x2a, 0x44, 0x25, 0x51, 0xc4, 0x8c, 0x67, 0xb0, 0x3a, 0xc1, 0xb0, 0x02, 0xca,
	0x6d, 0xea, 0x0b, 0x03, 0xd5, 0xb4, 0x7a, 0x19, 0xaf, 0xa4, 0x05, 0x47, 0x91, 0x11, 0xbd, 0x86,
	0xe5, 0x9e, 0x1f, 0xd5, 0xc8, 0xc0, 0x12, 0x5d, 0x4e, 0xc3, 0x2e, 0xf3, 0x1c, 0x23, 0xa7, 0xe2,
	0xaf, 0x4f, 0xc5, 0x3f, 0xd9, 0xf7, 0xc5, 0xe6, 0xc6, 0xe7, 0xc4, 0xeb, 0x51, 0xb3, 0x70, 0x61,
	0xde, 0xfa, 0xad, 0x96, 0xd5, 0x35, 0x8c, 0x86, 0x88, 0x76, 0x42, 0x40, 0xc7, 0xb0, 0x34, 0x8d,
	0xbd, 0x75, 0x23, 0xac, 0x3e, 0x05, 0xfd, 0x21, 0x14, 0x88, 0x27, 0xac, 0x80, 0x71, 0x61, 0xcc,
	0x5f, 0xcd, 0xc2, 0x79, 0xe2, 0x89, 0x23, 0xc6, 0x05, 0xda, 0x03, 0x9d, 0xd3, 0x5e, 0xa8, 0x0a,
	0xd4, 0xa7, 0xb6, 0x3c, 0x42, 0x23, 0xaf, 0x00, 0xd5, 0x29, 0x80, 0xc9, 0x98, 0x17, 0xc9, 0x17,
	0x95, 0x66, 0x67, 0x28, 0x41, 0xc7, 0xb0, 0xec, 0x33, 0x4b, 0x70, 0x72, 0x76, 0xe6, 0xda, 0xd6,
	0xf0, 0xe6, 0x2c, 0x5c, 0x75, 0x73, 0xe4, 0x9e, 0xbe, 0xd1, 0xb2, 0x4f, 0x32, 0x78, 0xc9, 0x67,
	0xed, 0x48, 0xbe, 0x9f, 0x5c, 0x20, 0x0c, 0xa3, 0xf3, 0x1b, 0x31, 0x2b, 0x37, 0x60, 0x0e, 0xe5,
	0x43, 0xe6, 0x17, 0xb0, 0x3a, 0x62, 0x52, 0xa7, 0x43, 0x47, 0xe0, 0xc5, 0xeb, 0x83, 0x57, 0x86,
	0x8c, 0x3d, 0xa7, 0x43, 0x87, 0xf0, 0xd7, 0xb0, 0x32, 0x1b, 0xad, 0x5f, 0x1f, 0xbd, 0x3c, 0x0b,
	0xfc, 0x08, 0x2a, 0x51, 0x3b, 0xf5, 0x58, 0xc7, 0x0a, 0x88, 0xe8, 0x1a, 0x4b, 0x35, 0xad, 0x5e,
	0xc4, 0x0b, 0x6a, 0xf5, 0x25, 0xeb, 0x1c, 0x11, 0xd1, 0x45, 0x07, 0x50, 0x4e, 0x35, 0x5d, 0xe3,
	0x8e, 0x0a, 0x5b, 0x6f, 0xcc, 0x6a, 0xeb, 0x8d, 0x3d, 0xe9, 0x7a, 0x1c, 0x79, 0xee, 0x28, 0x53,
	0x8c, 0x8b, 0xd7, 0xd0, 0x73, 0xa8, 0x11, 0xef, 0x4b, 0x32, 0x08, 0x55, 0xd4, 0xf1, 0x71, 0x61,
	0x9d, 0x11, 0xd7, 0xeb, 0x71, 0x1a, 0x1a, 0xcb, 0x35, 0xad, 0x5e, 0xc0, 0xeb, 0x91, 0xdf, 0x4b,
	0xd6, 0x19, 0xeb, 0x41, 0xcf, 0x63, 0x1f, 0x74, 0x00, 0x25, 0xe1, 0x85, 0x16, 0x0b, 0x54, 0x53,
	0x33, 0x56, 0x54, 0x52, 0x4f, 0x67, 0x27, 0x35, 0xa6, 0x6f, 0xb4, 0xbd, 0xf0, 0x55, 0xa4, 0xc1,
	0x20, 0x86, 0x7f, 0xa3, 0x9f, 0xc3, 0x7d, 0xc1, 0x89, 0x1f, 0xca, 0xbb, 0x6e, 0x85, 0xcc, 0x7e,
	0x43, 0x85, 0xa5, 0x86, 0x88, 0x65, 0x73, 0x57, 0x50, 0xee, 0x12, 0x63, 0x55, 0x85, 0x58, 0x9d,
	0x3a, 0xee, 0x63, 0xd5, 0xd2, 0xf1, 0xfa, 0x50, 0x7f, 0xac, 0xe4, 0x07, 0x52, 0xbd, 0x13, 0x8b,
	0x91, 0x05, 0x4b, 0x72, 0xd8, 0xa4, 0x36, 0x6c, 0x14, 0x14, 0xf1, 0xa3, 0xab, 0x93, 0x6e, 0x09,
	0x11, 0x8c, 0xfd, 0x6e, 0x65, 0xf0, 0x62, 0x37, 0xbd, 0x84, 0x7e, 0x06, 0xba, 0xb0, 0x27, 0xf8,
	0x45, 0xc5, 0xff, 0xf0, 0x1a, 0x87, 0x62, 0x4f, 0xe0, 0x2b, 0x22, 0xb5, 0x22, 0xd3, 0xef, 0xf0,
	0xc0, 0x4e, 0xe3, 0x4b, 0xd7, 0x4d, 0xff, 0x05, 0x0f, 0xec, 0x89, 0xf4, 0x3b, 0xe9, 0x25, 0x44,
	0x61, 0xd9, 0xee, 0x85, 0x82, 0x9d, 0xa7, 0x43, 0x94, 0x55, 0x88, 0xcd, 0xab, 0x43, 0xec, 0x28,
	0x71, 0x3a, 0xc8, 0x92, 0x3d, 0xb9, 0x58, 0xfd, 0x4a, 0x83, 0xfc, 0x11, 0x19, 0x78, 0x8c, 0x38,
	0xe8, 0x2e, 0xe4, 0x04, 0x7d, 0x1b, 0x8d, 0xb2, 0xa2, 0x99, 0xbf, 0x30, 0x73, 0x3c, 0x5b, 0xd3,
	0x5a, 0x19, 0xac, 0x96, 0x91, 0x01, 0xf3, 0xa7, 0xae, 0x4f, 0xf8, 0x40, 0x4d, 0xaa, 0x85, 0x56,
	0x06, 0xc7, 0xbf, 0xb7, 0x9e, 0xfe, 0xe1, 0xaf, 0xbf, 0xb9, 0xf7, 0x18, 0xde, 0x8f, 0x92, 0x22,
	0x81, 0xdb, 0xe8, 0x6f, 0x44, 0x49, 0x8d, 0x67, 0x14, 0x87, 0x31, 0x2b, 0x90, 0x0f, 0xe2, 0x88,
	0x73, 0xff, 0x35, 0xb5, 0xea, 0x3f, 0xe7, 0x61, 0x71, 0xe2, 0x7b, 0xa2, 0xfb, 0x90, 0xeb, 0xb2,
	0x30, 0x49, 0xa5, 0x74, 0x61, 0x16, 0xf8, 0xfc, 0x5f, 0xb4, 0xec, 0xb7, 0x5a, 0x06, 0x2b, 0x03,
	0x7a, 0x00, 0x39, 0x55, 0xa0, 0x59, 0xe5, 0x50, 0xbe, 0x30, 0x81, 0x17, 0x6a, 0x5a, 0xe2, 0x22,
	0x4d, 0xe8, 0xc7, 0x90, 0x0b, 0xa9, 0xef, 0xc4, 0xb3, 0xec, 0x83, 0xab, 0x8f, 0x2c, 0x4e, 0x10,
	0x2b, 0x19, 0xda, 0x81, 0x3c, 0xa7, 0x36, 0x75, 0xfb, 0x34, 0x9e, 0x46, 0x37, 0x20, 0x24, 0x4a,
	0xd4, 0x81, 0x3b, 0x9c, 0xfe, 0xaa, 0x47, 0x43, 0x21, 0x3f, 0xa3, 0x43, 0x79, 0x68, 0x09, 0x66,
	0x11, 0xc7, 0x31, 0xe6, 0x6b, 0x73, 0xf5, 0xd2, 0xc6, 0xe3, 0x77, 0x32, 0x1d, 0xca, 0xd5, 0x10,
	0x88, 0xea, 0xd1, 0x2c, 0x5e, 0x98, 0xf3, 0x5f, 0x6b, 0x73, 0xfa, 0xbf, 0xf2, 0x78, 0x39, 0x26,
	0x46, 0x4e, 0x61, 0x9b, 0x6d, 0x3b, 0x0e, 0xfa, 0x0c, 0xd6, 0x66, 0x04, 0xe2, 0xf4, 0x9c, 0xf5,
	0xa9, 0x51, 0xa8, 0xcd, 0xd5, 0x8b, 0xa6, 0x7e, 0x61, 0x96, 0xbf, 0xd6, 0xe0, 0xa1, 0x3a, 0x4a,
	0x4d, 0x9e, 0xd3, 0x9d, 0x49, 0x12, 0x56, 0xfe, 0xe8, 0x39, 0x2c, 0xd1, 0xb7, 0x01, 0xb5, 0x05,
	0x75, 0xac, 0xe8, 0xed, 0x44, 0x43, 0xa3, 0xa8, 0x12, 0x5e, 0x8b, 0x13, 0x96, 0xcf, 0x42, 0x99,
	0xe9, 0xbe, 0x2f, 0x9e, 0x7d, 0x8c, 0xe5, 0xdb, 0x10, 0xeb, 0x89, 0xe6, 0x38, 0x96, 0xa0, 0x36,
	0x2c, 0xd9, 0xcc, 0xa1, 0xb6, 0x65, 0x7b, 0xae, 0x6c, 0x98, 0x52, 0x64, 0x40, 0x4d, 0xab, 0x57,
	0x36, 0xee, 0x4d, 0x70, 0x76, 0xa4, 0xdf, 0x8e, 0x72, 0x6b, 0x0f, 0x82, 0x68, 0x0a, 0xff, 0x5a,
	0x4d, 0xe1, 0x45, 0x3b, 0x6d, 0x42, 0x9f, 0xc3, 0xed, 0xe4, 0xb9, 0xeb, 0x93, 0x73, 0x6a, 0xc5,
	0x4f, 0xdb, 0xb8, 0xfa, 0x1e, 0x8d, 0x83, 0x63, 0x93, 0x0c, 0x70, 0xac, 0x5e, 0xbd, 0x07, 0xd1,
	0x02, 0x46, 0x31, 0xe1, 0x90, 0x9c, 0xd3, 0x78, 0x0d, 0x61, 0x78, 0xd4, 0x75, 0x1d, 0x87, 0xfa,
	0x96, 0x22, 0x58, 0xc9, 0xcb, 0x52, 0x1e, 0xc2, 0x58, 0x3c, 0xf5, 0x88, 0x28, 0x9a, 0x59, 0x43,
	0xc3, 0xb5, 0xc8, 0x7f, 0x4f, 0xba, 0xef, 0x0e, 0xbd, 0x8f, 0x47, 0x68, 0x84, 0xe1, 0xc1, 0xbb,
	0x98, 0xf2, 0x35, 0x20, 0x7b, 0xd6, 0x86, 0x7a, 0x08, 0x14, 0xcc, 0xc2, 0x1f, 0xff, 0xfd, 0x9f,
	0x7f, 0xdc, 0xd2, 0x0c, 0x0d, 0xdf, 0x9d, 0x89, 0x3d, 0x09, 0xa9, 0xac, 0x92, 0x8d, 0xad, 0x8f,
	0x65, 0xb5, 0x35, 0xe1, 0xfb, 0x97, 0x57, 0xdb, 0x44, 0x45, 0x55, 0xff, 0xa6, 0x41, 0x25, 0xdd,
	0xd5, 0x86, 0x05, 0xa2, 0x7d, 0xe7, 0x02, 0xc9, 0xaa, 0xbb, 0xf1, 0x7f, 0x14, 0xc8, 0xd6, 0xa6,
	0xdc, 0x4c, 0x03, 0x9e, 0x5e, 0xbe, 0x99, 0x74, 0xe2, 0xd5, 0x2f, 0x40, 0xc7, 0xd4, 0x71, 0xc3,
	0xf1, 0xcd, 0xe8, 0x30, 0xf7, 0x86, 0x0e, 0xa2, 0x86, 0x81, 0xe5, 0x9f, 0x5b, 0x3f, 0x90, 0xe8,
	0x0f, 0xa1, 0x71, 0x39, 0x7a, 0x12, 0x54, 0xfd, 0x9d, 0x06, 0x8b, 0x13, 0xfd, 0x19, 0x3d, 0x80,
	0x85, 0xd4, 0x15, 0x88, 0xa2, 0x94, 0xc6, 0x2e, 0x11, 0xfa, 0x00, 0x8a, 0xa4, 0x27, 0xba, 0x8c,
	0xbb, 0x62, 0x10, 0x77, 0xa5, 0x54, 0xdb, 0x1a, 0x59, 0xaf, 0xf9, 0x01, 0x27, 0x72, 0xa8, 0xfe,
	0x3e, 0x0b, 0x4b, 0x53, 0x4d, 0x1d, 0xbd, 0x07, 0xb9, 0x51, 0x46, 0xc3, 0x9e, 0x8d, 0xd5, 0x22,
	0xfa, 0x11, 0x2c, 0xc8, 0x32, 0x70, 0xe2, 0x7f, 0x0b, 0xe3, 0x4e, 0x78, 0x7b, 0x6a, 0x60, 0x6f,
	0xfb, 0x83, 0x56, 0x06, 0x97, 0x94, 0x6f, 0xf4, 0x46, 0x41, 0xa7, 0x70, 0xef, 0x5d, 0x17, 0x37,
	0x86, 0x65, 0x2f, 0x9d, 0xfe, 0xb2, 0x3e, 0x5a, 0x19, 0xfc, 0xde, 0xcc, 0xab, 0x1c, 0xc5, 0xd8,
	0x7a, 0x26, 0xcf, 0xe1, 0x23, 0x68, 0x5e, 0x7e, 0x0e, 0x53, 0x7b, 0x36, 0xcb, 0x50, 0x8a, 0x72,
	0x50, 0x0d, 0xa5, 0xea, 0x00, 0x8c, 0xde, 0x30, 0xe8, 0x7d, 0xa8, 0x10, 0x2f, 0xf0, 0x2d, 0x95,
	0x8f, 0xcd, 0xbc, 0xd0, 0xd0, 0x64, 0xf7, 0xc3, 0x65, 0xb9, 0x7a, 0x94, 0x2c, 0x6e, 0x35, 0x65,
	0xec, 0x27, 0x50, 0xbf, 0xe2, 0xde, 0x0d, 0xb9, 0x5b, 0x8f, 0xa4, 0xe0, 0x3e, 0xdc, 0xbd, 0x54,
	0x60, 0xae, 0x40, 0x65, 0x7c, 0x5c, 0x53, 0xae, 0x46, 0xdc, 0xa7, 0xb9, 0x02, 0xe8, 0xa5, 0x27,
	0xbf, 0x80, 0x85, 0xc8, 0x37, 0x6a, 0x90, 0xa8, 0x04, 0xf9, 0x93, 0xc3, 0xcf, 0x0e, 0x5f, 0xbd,
	0x3e, 0xd4, 0x33, 0xf2, 0x47, 0x6b, 0x6f, 0xfb, 0x65, 0xbb, 0xf5, 0x53, 0x5d, 0x43, 0x65, 0x28,
	0x9e, 0x1c, 0x26, 0x3f, 0xb3, 0x68, 0x01, 0x0a, 0xbb, 0x78, 0x7b, 0xff, 0x70, 0xff, 0xf0, 0x85,
	0x3e, 0x27, 0x3d, 0xdb, 0xfb, 0x07, 0x7b, 0xaf, 0x4e, 0xda, 0x7a, 0x4e, 0x99, 0xf6, 0x5e, 0xe0,
	0xed, 0xdd, 0xbd, 0x5d, 0xfd, 0x96, 0xf9, 0xc9, 0x9f, 0xbe, 0xfa, 0xf6, 0xef, 0xf3, 0x59, 0x3d,
	0x0b, 0x0f, 0x5d, 0x16, 0x5d, 0xfc, 0x80, 0xb3, 0xb7, 0x83, 0x99, 0x85, 0x69, 0xea, 0x63, 0x99,
	0xab, 0x93, 0x39, 0xd2, 0x4e, 0xe7, 0xd5, 0xb9, 0x6d, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0xb2, 0x56, 0xf7, 0xcd, 0x10, 0x00, 0x00,
}
