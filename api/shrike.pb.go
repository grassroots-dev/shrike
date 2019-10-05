// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shrike.proto

package shrike

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

// Simple request message for stubbed service calls..
type StubRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StubRequest) Reset()         { *m = StubRequest{} }
func (m *StubRequest) String() string { return proto.CompactTextString(m) }
func (*StubRequest) ProtoMessage()    {}
func (*StubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2311c9d8f30f83c3, []int{0}
}

func (m *StubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StubRequest.Unmarshal(m, b)
}
func (m *StubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StubRequest.Marshal(b, m, deterministic)
}
func (m *StubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StubRequest.Merge(m, src)
}
func (m *StubRequest) XXX_Size() int {
	return xxx_messageInfo_StubRequest.Size(m)
}
func (m *StubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StubRequest proto.InternalMessageInfo

func (m *StubRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Simple response message for stubbed service calls..
type StubResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StubResponse) Reset()         { *m = StubResponse{} }
func (m *StubResponse) String() string { return proto.CompactTextString(m) }
func (*StubResponse) ProtoMessage()    {}
func (*StubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2311c9d8f30f83c3, []int{1}
}

func (m *StubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StubResponse.Unmarshal(m, b)
}
func (m *StubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StubResponse.Marshal(b, m, deterministic)
}
func (m *StubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StubResponse.Merge(m, src)
}
func (m *StubResponse) XXX_Size() int {
	return xxx_messageInfo_StubResponse.Size(m)
}
func (m *StubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StubResponse proto.InternalMessageInfo

func (m *StubResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*StubRequest)(nil), "shrike.StubRequest")
	proto.RegisterType((*StubResponse)(nil), "shrike.StubResponse")
}

func init() { proto.RegisterFile("shrike.proto", fileDescriptor_2311c9d8f30f83c3) }

var fileDescriptor_2311c9d8f30f83c3 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x89, 0x54, 0x05, 0xf5, 0xa8, 0x40, 0x1c, 0x0c, 0x55, 0x27, 0xc8, 0xd4, 0xa9, 0x42,
	0xb0, 0xc0, 0x00, 0x22, 0xca, 0xc0, 0x40, 0x91, 0x50, 0x23, 0xd8, 0xdd, 0xf0, 0x93, 0x44, 0x25,
	0x76, 0xb1, 0x2f, 0x7d, 0x49, 0x5e, 0x0a, 0x11, 0x37, 0x12, 0x1d, 0xdd, 0xcd, 0x67, 0xdd, 0xf7,
	0x7f, 0x27, 0xfb, 0x68, 0xe4, 0x2a, 0x5b, 0xaf, 0x30, 0x5b, 0x5b, 0x23, 0x86, 0x63, 0x5f, 0x25,
	0x97, 0x74, 0x94, 0x4b, 0xbb, 0x5c, 0xe0, 0xbb, 0x85, 0x13, 0x66, 0x1a, 0x68, 0xd5, 0x60, 0x1c,
	0x5d, 0x44, 0xd3, 0xe1, 0xa2, 0x3b, 0x27, 0x53, 0x1a, 0xf9, 0x16, 0xb7, 0x36, 0xda, 0x81, 0xc7,
	0x74, 0xd8, 0xc0, 0x39, 0x55, 0xf6, 0x6d, 0x7d, 0x79, 0xfd, 0x33, 0xa0, 0x38, 0xef, 0x72, 0x39,
	0x25, 0xce, 0x2a, 0x14, 0xab, 0xcc, 0xe8, 0xcf, 0xba, 0x6c, 0xad, 0x92, 0xda, 0x68, 0x3e, 0x9b,
	0x6d, 0x87, 0xf8, 0xe7, 0x9c, 0x9c, 0xef, 0x5e, 0x7a, 0x4b, 0x72, 0xc0, 0xb7, 0x34, 0xec, 0x69,
	0x84, 0x93, 0x7f, 0xf2, 0x37, 0x07, 0x1b, 0x46, 0xde, 0x11, 0x65, 0x16, 0x4a, 0x10, 0x8e, 0xde,
	0xd3, 0xb1, 0x47, 0x5f, 0xcc, 0x06, 0x0d, 0xb4, 0x84, 0xe1, 0x0f, 0x74, 0x92, 0xda, 0xa2, 0xaa,
	0x37, 0x7b, 0xf2, 0x8f, 0x74, 0xea, 0xf5, 0x73, 0xa5, 0x3f, 0x6a, 0x5d, 0xbe, 0xaa, 0x32, 0xf0,
	0xd5, 0x52, 0xe2, 0xed, 0x04, 0x7b, 0x47, 0x3c, 0xd3, 0xe4, 0x09, 0xf2, 0x6e, 0xbe, 0x5a, 0x2d,
	0x80, 0x9d, 0x9b, 0xa2, 0xfb, 0xf7, 0x5c, 0x2c, 0x54, 0x13, 0x14, 0x75, 0x15, 0x2d, 0xe3, 0x6e,
	0x53, 0x6f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x45, 0xd1, 0xb5, 0xc9, 0xb9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShrikeClient is the client API for Shrike service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShrikeClient interface {
	// Checks for the current configuration state of service.
	CheckConfiguration(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	// Configures a fresh instance of the shrike service when booting without a configuration.
	Configure(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	CheckUser(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	CreateUser(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	CreateMovement(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	ArchiveMovement(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	CreateLandingPage(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	ArchiveLandingPage(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error)
	GetVolunteerLocationStream(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (Shrike_GetVolunteerLocationStreamClient, error)
}

type shrikeClient struct {
	cc *grpc.ClientConn
}

func NewShrikeClient(cc *grpc.ClientConn) ShrikeClient {
	return &shrikeClient{cc}
}

func (c *shrikeClient) CheckConfiguration(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/CheckConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) Configure(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) CheckUser(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) CreateUser(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) CreateMovement(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/CreateMovement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) ArchiveMovement(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/ArchiveMovement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) CreateLandingPage(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/CreateLandingPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) ArchiveLandingPage(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (*StubResponse, error) {
	out := new(StubResponse)
	err := c.cc.Invoke(ctx, "/shrike.Shrike/ArchiveLandingPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrikeClient) GetVolunteerLocationStream(ctx context.Context, in *StubRequest, opts ...grpc.CallOption) (Shrike_GetVolunteerLocationStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Shrike_serviceDesc.Streams[0], "/shrike.Shrike/GetVolunteerLocationStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &shrikeGetVolunteerLocationStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Shrike_GetVolunteerLocationStreamClient interface {
	Recv() (*StubResponse, error)
	grpc.ClientStream
}

type shrikeGetVolunteerLocationStreamClient struct {
	grpc.ClientStream
}

func (x *shrikeGetVolunteerLocationStreamClient) Recv() (*StubResponse, error) {
	m := new(StubResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShrikeServer is the server API for Shrike service.
type ShrikeServer interface {
	// Checks for the current configuration state of service.
	CheckConfiguration(context.Context, *StubRequest) (*StubResponse, error)
	// Configures a fresh instance of the shrike service when booting without a configuration.
	Configure(context.Context, *StubRequest) (*StubResponse, error)
	CheckUser(context.Context, *StubRequest) (*StubResponse, error)
	CreateUser(context.Context, *StubRequest) (*StubResponse, error)
	CreateMovement(context.Context, *StubRequest) (*StubResponse, error)
	ArchiveMovement(context.Context, *StubRequest) (*StubResponse, error)
	CreateLandingPage(context.Context, *StubRequest) (*StubResponse, error)
	ArchiveLandingPage(context.Context, *StubRequest) (*StubResponse, error)
	GetVolunteerLocationStream(*StubRequest, Shrike_GetVolunteerLocationStreamServer) error
}

// UnimplementedShrikeServer can be embedded to have forward compatible implementations.
type UnimplementedShrikeServer struct {
}

func (*UnimplementedShrikeServer) CheckConfiguration(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckConfiguration not implemented")
}
func (*UnimplementedShrikeServer) Configure(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedShrikeServer) CheckUser(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (*UnimplementedShrikeServer) CreateUser(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedShrikeServer) CreateMovement(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovement not implemented")
}
func (*UnimplementedShrikeServer) ArchiveMovement(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveMovement not implemented")
}
func (*UnimplementedShrikeServer) CreateLandingPage(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLandingPage not implemented")
}
func (*UnimplementedShrikeServer) ArchiveLandingPage(ctx context.Context, req *StubRequest) (*StubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveLandingPage not implemented")
}
func (*UnimplementedShrikeServer) GetVolunteerLocationStream(req *StubRequest, srv Shrike_GetVolunteerLocationStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVolunteerLocationStream not implemented")
}

func RegisterShrikeServer(s *grpc.Server, srv ShrikeServer) {
	s.RegisterService(&_Shrike_serviceDesc, srv)
}

func _Shrike_CheckConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).CheckConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/CheckConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).CheckConfiguration(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).Configure(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).CheckUser(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).CreateUser(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_CreateMovement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).CreateMovement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/CreateMovement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).CreateMovement(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_ArchiveMovement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).ArchiveMovement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/ArchiveMovement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).ArchiveMovement(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_CreateLandingPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).CreateLandingPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/CreateLandingPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).CreateLandingPage(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_ArchiveLandingPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrikeServer).ArchiveLandingPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shrike.Shrike/ArchiveLandingPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrikeServer).ArchiveLandingPage(ctx, req.(*StubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrike_GetVolunteerLocationStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StubRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShrikeServer).GetVolunteerLocationStream(m, &shrikeGetVolunteerLocationStreamServer{stream})
}

type Shrike_GetVolunteerLocationStreamServer interface {
	Send(*StubResponse) error
	grpc.ServerStream
}

type shrikeGetVolunteerLocationStreamServer struct {
	grpc.ServerStream
}

func (x *shrikeGetVolunteerLocationStreamServer) Send(m *StubResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Shrike_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shrike.Shrike",
	HandlerType: (*ShrikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckConfiguration",
			Handler:    _Shrike_CheckConfiguration_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Shrike_Configure_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _Shrike_CheckUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Shrike_CreateUser_Handler,
		},
		{
			MethodName: "CreateMovement",
			Handler:    _Shrike_CreateMovement_Handler,
		},
		{
			MethodName: "ArchiveMovement",
			Handler:    _Shrike_ArchiveMovement_Handler,
		},
		{
			MethodName: "CreateLandingPage",
			Handler:    _Shrike_CreateLandingPage_Handler,
		},
		{
			MethodName: "ArchiveLandingPage",
			Handler:    _Shrike_ArchiveLandingPage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetVolunteerLocationStream",
			Handler:       _Shrike_GetVolunteerLocationStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shrike.proto",
}
