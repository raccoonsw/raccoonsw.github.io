// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_email_server/email.proto

package grpc_email_server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f5fdf5681c6555, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Request) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Request) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Response struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f5fdf5681c6555, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "grpc_email_server.Request")
	proto.RegisterType((*Response)(nil), "grpc_email_server.Response")
}

func init() {
	proto.RegisterFile("grpc_email_server/email.proto", fileDescriptor_c8f5fdf5681c6555)
}

var fileDescriptor_c8f5fdf5681c6555 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2f, 0x2a, 0x48,
	0x8e, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x07, 0x73,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x31, 0xa4, 0x95, 0x02, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0xf3, 0x8b, 0x52, 0x52, 0x8b, 0x3c, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x60, 0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xcc, 0x92, 0xd4,
	0x5c, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0x04, 0x94, 0x27, 0x24, 0xc2, 0xc5, 0x0a, 0x36, 0x4c, 0x82,
	0x59, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x51, 0x52, 0xe3, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe2, 0xe2, 0x28, 0x82, 0xb2, 0xc1, 0x86, 0x72, 0x06, 0xc1, 0xf9,
	0x46, 0x81, 0x5c, 0x3c, 0xae, 0x20, 0x0d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x8e,
	0x5c, 0x2c, 0xc1, 0xa9, 0x79, 0x29, 0x42, 0x52, 0x7a, 0x18, 0xce, 0xd4, 0x83, 0xba, 0x51, 0x4a,
	0x1a, 0xab, 0x1c, 0xc4, 0x40, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x3f, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xee, 0x33, 0x62, 0x04, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc_email_server.EmailService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Send(context.Context, *Request) (*Response, error)
}

// UnimplementedEmailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (*UnimplementedEmailServiceServer) Send(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_email_server.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_email_server.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_email_server/email.proto",
}