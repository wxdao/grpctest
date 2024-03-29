// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package greet

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

type EchoRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReply) Reset()         { *m = EchoReply{} }
func (m *EchoReply) String() string { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()    {}
func (*EchoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *EchoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReply.Unmarshal(m, b)
}
func (m *EchoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReply.Marshal(b, m, deterministic)
}
func (m *EchoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReply.Merge(m, src)
}
func (m *EchoReply) XXX_Size() int {
	return xxx_messageInfo_EchoReply.Size(m)
}
func (m *EchoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReply proto.InternalMessageInfo

func (m *EchoReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TickRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Interval             int64    `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickRequest) Reset()         { *m = TickRequest{} }
func (m *TickRequest) String() string { return proto.CompactTextString(m) }
func (*TickRequest) ProtoMessage()    {}
func (*TickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{2}
}

func (m *TickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TickRequest.Unmarshal(m, b)
}
func (m *TickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TickRequest.Marshal(b, m, deterministic)
}
func (m *TickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickRequest.Merge(m, src)
}
func (m *TickRequest) XXX_Size() int {
	return xxx_messageInfo_TickRequest.Size(m)
}
func (m *TickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TickRequest proto.InternalMessageInfo

func (m *TickRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TickRequest) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type TickReply struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickReply) Reset()         { *m = TickReply{} }
func (m *TickReply) String() string { return proto.CompactTextString(m) }
func (*TickReply) ProtoMessage()    {}
func (*TickReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{3}
}

func (m *TickReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TickReply.Unmarshal(m, b)
}
func (m *TickReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TickReply.Marshal(b, m, deterministic)
}
func (m *TickReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickReply.Merge(m, src)
}
func (m *TickReply) XXX_Size() int {
	return xxx_messageInfo_TickReply.Size(m)
}
func (m *TickReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TickReply.DiscardUnknown(m)
}

var xxx_messageInfo_TickReply proto.InternalMessageInfo

func (m *TickReply) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "greet.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "greet.EchoReply")
	proto.RegisterType((*TickRequest)(nil), "greet.TickRequest")
	proto.RegisterType((*TickReply)(nil), "greet.TickReply")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xd4, 0xb9, 0xb8, 0x5d,
	0x93, 0x33, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x55,
	0x2e, 0x4e, 0x88, 0xc2, 0x82, 0x9c, 0x4a, 0x3c, 0xca, 0xec, 0xb9, 0xb8, 0x43, 0x32, 0x93, 0xb3,
	0x61, 0xe6, 0x89, 0x70, 0xb1, 0x26, 0xe7, 0x97, 0xe6, 0x95, 0x80, 0x95, 0x31, 0x07, 0x41, 0x38,
	0x42, 0x52, 0x5c, 0x1c, 0x99, 0x79, 0x25, 0xa9, 0x45, 0x65, 0x89, 0x39, 0x12, 0x4c, 0x60, 0x09,
	0x38, 0x5f, 0x49, 0x93, 0x8b, 0x13, 0x62, 0x00, 0xc8, 0x1e, 0x19, 0x2e, 0xce, 0x92, 0xcc, 0xdc,
	0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0xa8, 0x11, 0x08, 0x01, 0xa3, 0x1c, 0x2e, 0x1e, 0x77, 0x90,
	0x27, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x74, 0xb8, 0x58, 0x40, 0x4e, 0x14, 0x12,
	0xd2, 0x83, 0x78, 0x14, 0xc9, 0x63, 0x52, 0x02, 0x28, 0x62, 0x20, 0xb3, 0xf5, 0xb8, 0x58, 0x40,
	0x16, 0xc1, 0x55, 0x23, 0x39, 0x1b, 0xae, 0x1a, 0xee, 0x12, 0x03, 0xc6, 0x24, 0x36, 0x70, 0xb8,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xd4, 0xb2, 0x21, 0x46, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (GreetService_TickClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (GreetService_TickClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/Tick", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_TickClient interface {
	Recv() (*TickReply, error)
	grpc.ClientStream
}

type greetServiceTickClient struct {
	grpc.ClientStream
}

func (x *greetServiceTickClient) Recv() (*TickReply, error) {
	m := new(TickReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
	Tick(*TickRequest, GreetService_TickServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_Tick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TickRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Tick(m, &greetServiceTickServer{stream})
}

type GreetService_TickServer interface {
	Send(*TickReply) error
	grpc.ServerStream
}

type greetServiceTickServer struct {
	grpc.ServerStream
}

func (x *greetServiceTickServer) Send(m *TickReply) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _GreetService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tick",
			Handler:       _GreetService_Tick_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greet.proto",
}
