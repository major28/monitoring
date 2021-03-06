// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitoring.proto

package monitoring

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type GetServiceInfoByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceInfoByNameRequest) Reset()         { *m = GetServiceInfoByNameRequest{} }
func (m *GetServiceInfoByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceInfoByNameRequest) ProtoMessage()    {}
func (*GetServiceInfoByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{1}
}

func (m *GetServiceInfoByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceInfoByNameRequest.Unmarshal(m, b)
}
func (m *GetServiceInfoByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceInfoByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceInfoByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceInfoByNameRequest.Merge(m, src)
}
func (m *GetServiceInfoByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceInfoByNameRequest.Size(m)
}
func (m *GetServiceInfoByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceInfoByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceInfoByNameRequest proto.InternalMessageInfo

func (m *GetServiceInfoByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetServiceInfoByNameResponse struct {
	IsAvailable          bool     `protobuf:"varint,1,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
	LastResponseTime     int64    `protobuf:"varint,2,opt,name=last_response_time,json=lastResponseTime,proto3" json:"last_response_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceInfoByNameResponse) Reset()         { *m = GetServiceInfoByNameResponse{} }
func (m *GetServiceInfoByNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceInfoByNameResponse) ProtoMessage()    {}
func (*GetServiceInfoByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{2}
}

func (m *GetServiceInfoByNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceInfoByNameResponse.Unmarshal(m, b)
}
func (m *GetServiceInfoByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceInfoByNameResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceInfoByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceInfoByNameResponse.Merge(m, src)
}
func (m *GetServiceInfoByNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceInfoByNameResponse.Size(m)
}
func (m *GetServiceInfoByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceInfoByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceInfoByNameResponse proto.InternalMessageInfo

func (m *GetServiceInfoByNameResponse) GetIsAvailable() bool {
	if m != nil {
		return m.IsAvailable
	}
	return false
}

func (m *GetServiceInfoByNameResponse) GetLastResponseTime() int64 {
	if m != nil {
		return m.LastResponseTime
	}
	return 0
}

type ServiceInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInfoResponse) Reset()         { *m = ServiceInfoResponse{} }
func (m *ServiceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceInfoResponse) ProtoMessage()    {}
func (*ServiceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{3}
}

func (m *ServiceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfoResponse.Unmarshal(m, b)
}
func (m *ServiceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfoResponse.Marshal(b, m, deterministic)
}
func (m *ServiceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfoResponse.Merge(m, src)
}
func (m *ServiceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceInfoResponse.Size(m)
}
func (m *ServiceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfoResponse proto.InternalMessageInfo

func (m *ServiceInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetHistoryRequestsResponse struct {
	GetHistoryRequests   []*GetHistoryRequest `protobuf:"bytes,1,rep,name=getHistoryRequests,proto3" json:"getHistoryRequests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetHistoryRequestsResponse) Reset()         { *m = GetHistoryRequestsResponse{} }
func (m *GetHistoryRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*GetHistoryRequestsResponse) ProtoMessage()    {}
func (*GetHistoryRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{4}
}

func (m *GetHistoryRequestsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryRequestsResponse.Unmarshal(m, b)
}
func (m *GetHistoryRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryRequestsResponse.Marshal(b, m, deterministic)
}
func (m *GetHistoryRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryRequestsResponse.Merge(m, src)
}
func (m *GetHistoryRequestsResponse) XXX_Size() int {
	return xxx_messageInfo_GetHistoryRequestsResponse.Size(m)
}
func (m *GetHistoryRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryRequestsResponse proto.InternalMessageInfo

func (m *GetHistoryRequestsResponse) GetGetHistoryRequests() []*GetHistoryRequest {
	if m != nil {
		return m.GetHistoryRequests
	}
	return nil
}

type GetHistoryRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryRequest) Reset()         { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()    {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9068e71705f3706, []int{5}
}

func (m *GetHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryRequest.Unmarshal(m, b)
}
func (m *GetHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryRequest.Merge(m, src)
}
func (m *GetHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetHistoryRequest.Size(m)
}
func (m *GetHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryRequest proto.InternalMessageInfo

func (m *GetHistoryRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *GetHistoryRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*GetServiceInfoByNameRequest)(nil), "GetServiceInfoByNameRequest")
	proto.RegisterType((*GetServiceInfoByNameResponse)(nil), "GetServiceInfoByNameResponse")
	proto.RegisterType((*ServiceInfoResponse)(nil), "serviceInfoResponse")
	proto.RegisterType((*GetHistoryRequestsResponse)(nil), "getHistoryRequestsResponse")
	proto.RegisterType((*GetHistoryRequest)(nil), "getHistoryRequest")
}

func init() { proto.RegisterFile("monitoring.proto", fileDescriptor_b9068e71705f3706) }

var fileDescriptor_b9068e71705f3706 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4f, 0x02, 0x31,
	0x10, 0x65, 0xc1, 0xa8, 0x0c, 0x68, 0x70, 0xe4, 0xb0, 0xe1, 0x23, 0xc1, 0x9e, 0x30, 0x31, 0x4d,
	0xc4, 0x5f, 0x00, 0x89, 0x5f, 0x07, 0x3c, 0x54, 0x8d, 0x47, 0x2c, 0x5a, 0xb1, 0x09, 0x6d, 0xb1,
	0xad, 0x44, 0x7e, 0x99, 0x7f, 0xcf, 0xf0, 0xb1, 0x0b, 0x71, 0x2b, 0x07, 0x6f, 0xb3, 0x7d, 0xfb,
	0xde, 0xcc, 0xbc, 0x37, 0x50, 0x51, 0x46, 0x4b, 0x6f, 0xac, 0xd4, 0x23, 0x3a, 0xb1, 0xc6, 0x1b,
	0x72, 0x08, 0xe5, 0x4b, 0x35, 0xf1, 0x33, 0x26, 0x3e, 0x3e, 0x85, 0xf3, 0xe4, 0x1c, 0xea, 0xd7,
	0xc2, 0xdf, 0x0b, 0x3b, 0x95, 0x2f, 0xe2, 0x56, 0xbf, 0x99, 0xde, 0xec, 0x8e, 0x2b, 0xb1, 0x82,
	0x11, 0x61, 0x47, 0x73, 0x25, 0xe2, 0xa8, 0x15, 0xb5, 0x8b, 0x6c, 0x51, 0x13, 0x03, 0x8d, 0x30,
	0xc5, 0x4d, 0x8c, 0x76, 0x02, 0x4f, 0xa0, 0x2c, 0xdd, 0x80, 0x4f, 0xb9, 0x1c, 0xf3, 0xe1, 0x78,
	0xc9, 0xdd, 0x67, 0x25, 0xe9, 0xba, 0xc9, 0x13, 0x9e, 0x01, 0x8e, 0xb9, 0xf3, 0x03, 0xbb, 0xe2,
	0x0c, 0xbc, 0x54, 0x22, 0xce, 0xb7, 0xa2, 0x76, 0x81, 0x55, 0xe6, 0x48, 0x22, 0xf6, 0x20, 0x95,
	0x20, 0xa7, 0x70, 0xec, 0xd6, 0xdd, 0xd2, 0x3e, 0xa1, 0xd9, 0x9e, 0xa1, 0x36, 0x12, 0xfe, 0x46,
	0x3a, 0x6f, 0x6c, 0xb2, 0xa3, 0x4b, 0x19, 0x3d, 0xc0, 0x2c, 0x1a, 0x47, 0xad, 0x42, 0xbb, 0xd4,
	0x41, 0x9a, 0x81, 0x58, 0xe0, 0x6f, 0xd2, 0x85, 0xa3, 0xcc, 0x2b, 0xc6, 0xb0, 0x67, 0x97, 0xe5,
	0x6a, 0x9a, 0xe4, 0x73, 0x3e, 0xe4, 0x2b, 0xf7, 0xcb, 0xdd, 0x8a, 0x6c, 0x51, 0x77, 0xbe, 0xf3,
	0x00, 0xfd, 0x34, 0x18, 0x7c, 0x84, 0xea, 0x28, 0xe0, 0x27, 0x36, 0xe8, 0x96, 0x64, 0x6a, 0x4d,
	0xba, 0x2d, 0x04, 0x92, 0xc3, 0x2b, 0x68, 0xae, 0x65, 0x9f, 0xa4, 0x7f, 0xef, 0x4b, 0xbd, 0x69,
	0x2b, 0x1e, 0xd0, 0xcd, 0x4b, 0xa8, 0x55, 0x69, 0xc0, 0xe4, 0xa0, 0x0e, 0xff, 0xfa, 0x8f, 0x4e,
	0xd0, 0xfc, 0xdf, 0xe4, 0x3a, 0xfd, 0x3b, 0x3e, 0x92, 0x1b, 0xee, 0x2e, 0x8e, 0xf8, 0xe2, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xba, 0x10, 0x42, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitoringClient is the client API for Monitoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitoringClient interface {
	GetServiceInfoByName(ctx context.Context, in *GetServiceInfoByNameRequest, opts ...grpc.CallOption) (*GetServiceInfoByNameResponse, error)
	GetServiceWithMinResponseTime(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error)
	GetServiceWithMaxResponseTime(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error)
	GetHistoryRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetHistoryRequestsResponse, error)
}

type monitoringClient struct {
	cc *grpc.ClientConn
}

func NewMonitoringClient(cc *grpc.ClientConn) MonitoringClient {
	return &monitoringClient{cc}
}

func (c *monitoringClient) GetServiceInfoByName(ctx context.Context, in *GetServiceInfoByNameRequest, opts ...grpc.CallOption) (*GetServiceInfoByNameResponse, error) {
	out := new(GetServiceInfoByNameResponse)
	err := c.cc.Invoke(ctx, "/Monitoring/getServiceInfoByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetServiceWithMinResponseTime(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error) {
	out := new(ServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/Monitoring/getServiceWithMinResponseTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetServiceWithMaxResponseTime(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error) {
	out := new(ServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/Monitoring/getServiceWithMaxResponseTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetHistoryRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetHistoryRequestsResponse, error) {
	out := new(GetHistoryRequestsResponse)
	err := c.cc.Invoke(ctx, "/Monitoring/getHistoryRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServer is the server API for Monitoring service.
type MonitoringServer interface {
	GetServiceInfoByName(context.Context, *GetServiceInfoByNameRequest) (*GetServiceInfoByNameResponse, error)
	GetServiceWithMinResponseTime(context.Context, *EmptyRequest) (*ServiceInfoResponse, error)
	GetServiceWithMaxResponseTime(context.Context, *EmptyRequest) (*ServiceInfoResponse, error)
	GetHistoryRequests(context.Context, *EmptyRequest) (*GetHistoryRequestsResponse, error)
}

func RegisterMonitoringServer(s *grpc.Server, srv MonitoringServer) {
	s.RegisterService(&_Monitoring_serviceDesc, srv)
}

func _Monitoring_GetServiceInfoByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceInfoByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetServiceInfoByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Monitoring/GetServiceInfoByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetServiceInfoByName(ctx, req.(*GetServiceInfoByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetServiceWithMinResponseTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetServiceWithMinResponseTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Monitoring/GetServiceWithMinResponseTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetServiceWithMinResponseTime(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetServiceWithMaxResponseTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetServiceWithMaxResponseTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Monitoring/GetServiceWithMaxResponseTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetServiceWithMaxResponseTime(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetHistoryRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetHistoryRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Monitoring/GetHistoryRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetHistoryRequests(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Monitoring",
	HandlerType: (*MonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getServiceInfoByName",
			Handler:    _Monitoring_GetServiceInfoByName_Handler,
		},
		{
			MethodName: "getServiceWithMinResponseTime",
			Handler:    _Monitoring_GetServiceWithMinResponseTime_Handler,
		},
		{
			MethodName: "getServiceWithMaxResponseTime",
			Handler:    _Monitoring_GetServiceWithMaxResponseTime_Handler,
		},
		{
			MethodName: "getHistoryRequests",
			Handler:    _Monitoring_GetHistoryRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring.proto",
}
