// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-timer assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ValidateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidateRequest)(nil), "model.ValidateRequest")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0x92, 0x4e,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x26, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0xd4, 0x28, 0xa9, 0x73, 0xf1, 0x87, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0x06,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7, 0xe6, 0x49,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x46, 0xde, 0x5c, 0xdc, 0x8e, 0xa5, 0x25, 0x19,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x36, 0x5c, 0x1c, 0x30, 0x7d, 0x42, 0x62, 0x7a,
	0x60, 0x8b, 0xf4, 0xd0, 0x0c, 0x92, 0x12, 0xd3, 0x83, 0xd8, 0xac, 0x07, 0xb3, 0x59, 0xcf, 0x15,
	0x64, 0xb3, 0x12, 0x83, 0x93, 0x0a, 0x97, 0x42, 0x72, 0x7e, 0xae, 0x5e, 0x72, 0x71, 0x46, 0x6a,
	0x81, 0x89, 0x5e, 0x41, 0x51, 0x6a, 0x6e, 0x66, 0x6a, 0x51, 0x41, 0x51, 0x6a, 0x4a, 0x66, 0x72,
	0x49, 0x7e, 0x91, 0x1e, 0xc8, 0x0f, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x7d, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0x51, 0xf5, 0xbf, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-timer assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.AuthService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Validate(context.Context, *ValidateRequest) (*empty.Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Validate(ctx context.Context, req *ValidateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AuthService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _AuthService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
