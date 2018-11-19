// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/auth/auth.proto

package testingpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	FillUsername         bool     `protobuf:"varint,4,opt,name=fill_username,json=fillUsername" json:"fill_username,omitempty"`
	FillOauthScope       bool     `protobuf:"varint,5,opt,name=fill_oauth_scope,json=fillOauthScope" json:"fill_oauth_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_28e9311690aa0936, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetFillUsername() bool {
	if m != nil {
		return m.FillUsername
	}
	return false
}

func (m *Request) GetFillOauthScope() bool {
	if m != nil {
		return m.FillOauthScope
	}
	return false
}

type Response struct {
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	OauthScope           string   `protobuf:"bytes,3,opt,name=oauth_scope,json=oauthScope" json:"oauth_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_28e9311690aa0936, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Response) GetOauthScope() string {
	if m != nil {
		return m.OauthScope
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "grpc.testing.Request")
	proto.RegisterType((*Response)(nil), "grpc.testing.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	UnaryCall(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	UnaryCall(context.Context, *Request) (*Response, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/auth/auth.proto",
}

func init() { proto.RegisterFile("testdata/auth/auth.proto", fileDescriptor_auth_28e9311690aa0936) }

var fileDescriptor_auth_28e9311690aa0936 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4b, 0x87, 0x30,
	0x18, 0xc6, 0xb1, 0x9f, 0xfa, 0x6a, 0x11, 0x83, 0x42, 0xbc, 0x24, 0x76, 0xf1, 0xb4, 0xa0, 0x6e,
	0x1d, 0xeb, 0x10, 0x9d, 0x82, 0x99, 0x10, 0x5d, 0x64, 0xda, 0x9b, 0x09, 0x6b, 0x5b, 0xdb, 0x0c,
	0xfa, 0xef, 0x63, 0x4b, 0x44, 0xbe, 0x97, 0xc1, 0x3e, 0xcf, 0xc3, 0x87, 0x3d, 0x83, 0xdc, 0xa1,
	0x75, 0xef, 0xdc, 0xf1, 0x6b, 0x3e, 0xbb, 0xcf, 0x70, 0x50, 0x6d, 0x94, 0x53, 0x24, 0x1b, 0x8d,
	0x1e, 0xa8, 0x8f, 0x27, 0x39, 0x56, 0xaf, 0x70, 0xcc, 0xf0, 0x7b, 0x46, 0xeb, 0xc8, 0x15, 0x9c,
	0x7c, 0x4c, 0x42, 0x74, 0xb3, 0x45, 0x23, 0xf9, 0x17, 0xe6, 0x07, 0x65, 0x54, 0xc7, 0x2c, 0xf3,
	0xb0, 0x5d, 0x18, 0xa9, 0xe1, 0x2c, 0x94, 0x94, 0x37, 0x76, 0x76, 0x50, 0x1a, 0xf3, 0xc3, 0xd0,
	0x3b, 0xf5, 0xfc, 0xd9, 0xe3, 0xc6, 0xd3, 0xea, 0x11, 0x62, 0x86, 0x56, 0x2b, 0x69, 0x91, 0x14,
	0x10, 0xaf, 0xd6, 0xbd, 0x32, 0xaa, 0x13, 0xb6, 0xde, 0xc9, 0x25, 0xa4, 0x5b, 0xd9, 0x7e, 0x88,
	0x41, 0xad, 0xa2, 0x9b, 0x27, 0x48, 0x5f, 0xd0, 0xba, 0x06, 0xcd, 0xcf, 0x34, 0x20, 0xb9, 0x83,
	0xa4, 0x95, 0xdc, 0xfc, 0x3e, 0x70, 0x21, 0xc8, 0x39, 0xdd, 0xae, 0xa1, 0xcb, 0x94, 0xe2, 0x62,
	0x17, 0xff, 0xbf, 0xe3, 0x3e, 0x7d, 0x4b, 0x16, 0xa6, 0xfb, 0xfe, 0x28, 0xfc, 0xc7, 0xed, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xb7, 0xca, 0x13, 0x2b, 0x01, 0x00, 0x00,
}
