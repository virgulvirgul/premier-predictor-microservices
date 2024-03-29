// Code generated by protoc-gen-go. DO NOT EDIT.
// source: legacy.proto

package model

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

type LegacyIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LegacyIdRequest) Reset()         { *m = LegacyIdRequest{} }
func (m *LegacyIdRequest) String() string { return proto.CompactTextString(m) }
func (*LegacyIdRequest) ProtoMessage()    {}
func (*LegacyIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_legacy_0a87f9f856d8b4bc, []int{0}
}
func (m *LegacyIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LegacyIdRequest.Unmarshal(m, b)
}
func (m *LegacyIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LegacyIdRequest.Marshal(b, m, deterministic)
}
func (dst *LegacyIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyIdRequest.Merge(dst, src)
}
func (m *LegacyIdRequest) XXX_Size() int {
	return xxx_messageInfo_LegacyIdRequest.Size(m)
}
func (m *LegacyIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyIdRequest proto.InternalMessageInfo

func (m *LegacyIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_legacy_0a87f9f856d8b4bc, []int{1}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LegacyUserResponse struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Surname              string   `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Joined               string   `protobuf:"bytes,5,opt,name=joined,proto3" json:"joined,omitempty"`
	AdFree               bool     `protobuf:"varint,6,opt,name=adFree,proto3" json:"adFree,omitempty"`
	Admin                bool     `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LegacyUserResponse) Reset()         { *m = LegacyUserResponse{} }
func (m *LegacyUserResponse) String() string { return proto.CompactTextString(m) }
func (*LegacyUserResponse) ProtoMessage()    {}
func (*LegacyUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_legacy_0a87f9f856d8b4bc, []int{2}
}
func (m *LegacyUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LegacyUserResponse.Unmarshal(m, b)
}
func (m *LegacyUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LegacyUserResponse.Marshal(b, m, deterministic)
}
func (dst *LegacyUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyUserResponse.Merge(dst, src)
}
func (m *LegacyUserResponse) XXX_Size() int {
	return xxx_messageInfo_LegacyUserResponse.Size(m)
}
func (m *LegacyUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyUserResponse proto.InternalMessageInfo

func (m *LegacyUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LegacyUserResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LegacyUserResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *LegacyUserResponse) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *LegacyUserResponse) GetJoined() string {
	if m != nil {
		return m.Joined
	}
	return ""
}

func (m *LegacyUserResponse) GetAdFree() bool {
	if m != nil {
		return m.AdFree
	}
	return false
}

func (m *LegacyUserResponse) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func init() {
	proto.RegisterType((*LegacyIdRequest)(nil), "model.LegacyIdRequest")
	proto.RegisterType((*LoginRequest)(nil), "model.LoginRequest")
	proto.RegisterType((*LegacyUserResponse)(nil), "model.LegacyUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LegacyUserServiceClient is the client API for LegacyUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LegacyUserServiceClient interface {
	GetLegacyUserById(ctx context.Context, in *LegacyIdRequest, opts ...grpc.CallOption) (*LegacyUserResponse, error)
	LegacyLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LegacyUserResponse, error)
}

type legacyUserServiceClient struct {
	cc *grpc.ClientConn
}

func NewLegacyUserServiceClient(cc *grpc.ClientConn) LegacyUserServiceClient {
	return &legacyUserServiceClient{cc}
}

func (c *legacyUserServiceClient) GetLegacyUserById(ctx context.Context, in *LegacyIdRequest, opts ...grpc.CallOption) (*LegacyUserResponse, error) {
	out := new(LegacyUserResponse)
	err := c.cc.Invoke(ctx, "/model.LegacyUserService/GetLegacyUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *legacyUserServiceClient) LegacyLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LegacyUserResponse, error) {
	out := new(LegacyUserResponse)
	err := c.cc.Invoke(ctx, "/model.LegacyUserService/LegacyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LegacyUserServiceServer is the server API for LegacyUserService service.
type LegacyUserServiceServer interface {
	GetLegacyUserById(context.Context, *LegacyIdRequest) (*LegacyUserResponse, error)
	LegacyLogin(context.Context, *LoginRequest) (*LegacyUserResponse, error)
}

func RegisterLegacyUserServiceServer(s *grpc.Server, srv LegacyUserServiceServer) {
	s.RegisterService(&_LegacyUserService_serviceDesc, srv)
}

func _LegacyUserService_GetLegacyUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LegacyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegacyUserServiceServer).GetLegacyUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.LegacyUserService/GetLegacyUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegacyUserServiceServer).GetLegacyUserById(ctx, req.(*LegacyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LegacyUserService_LegacyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegacyUserServiceServer).LegacyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.LegacyUserService/LegacyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegacyUserServiceServer).LegacyLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LegacyUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.LegacyUserService",
	HandlerType: (*LegacyUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLegacyUserById",
			Handler:    _LegacyUserService_GetLegacyUserById_Handler,
		},
		{
			MethodName: "LegacyLogin",
			Handler:    _LegacyUserService_LegacyLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "legacy.proto",
}

func init() { proto.RegisterFile("legacy.proto", fileDescriptor_legacy_0a87f9f856d8b4bc) }

var fileDescriptor_legacy_0a87f9f856d8b4bc = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xc9, 0xb4, 0xfb, 0xf3, 0x3a, 0x94, 0x45, 0x19, 0x71, 0x78, 0x98, 0x3d, 0xc8, 0x4e,
	0x39, 0xa8, 0x57, 0x41, 0x76, 0x98, 0x0c, 0x86, 0x48, 0xc5, 0x0f, 0x10, 0x9b, 0xd7, 0x19, 0x69,
	0x9a, 0x98, 0x74, 0xca, 0xbe, 0x8a, 0x9f, 0xc7, 0x0f, 0x26, 0x4d, 0x57, 0xeb, 0x3c, 0x08, 0x1e,
	0x9f, 0xdf, 0xef, 0x6d, 0x93, 0x3c, 0x2f, 0xf4, 0x33, 0x5c, 0x8a, 0x74, 0xcd, 0xad, 0x33, 0x85,
	0xa1, 0x91, 0x36, 0x12, 0xb3, 0xf8, 0x14, 0x0e, 0x16, 0x01, 0xcf, 0x65, 0x82, 0xaf, 0x2b, 0xf4,
	0x05, 0xdd, 0x87, 0x96, 0x92, 0x8c, 0x8c, 0xc9, 0x24, 0x4a, 0x5a, 0x4a, 0xc6, 0xd7, 0xd0, 0x5f,
	0x98, 0xa5, 0xca, 0x6b, 0x7f, 0x04, 0x11, 0x6a, 0xa1, 0xb2, 0x30, 0xd2, 0x4b, 0xaa, 0x40, 0x47,
	0xd0, 0xb5, 0xc2, 0xfb, 0x77, 0xe3, 0x24, 0x6b, 0x05, 0xf1, 0x9d, 0xe3, 0x4f, 0x02, 0xb4, 0x3a,
	0xe5, 0xc1, 0xa3, 0x4b, 0xd0, 0x5b, 0x93, 0x7b, 0xfc, 0xff, 0x8f, 0xe8, 0x09, 0xf4, 0x9e, 0x94,
	0xf3, 0xc5, 0xad, 0xd0, 0xc8, 0x76, 0x82, 0x6c, 0x00, 0x65, 0xd0, 0xf1, 0x2b, 0x97, 0x97, 0x6e,
	0x37, 0xb8, 0x3a, 0xd2, 0x21, 0xb4, 0x5f, 0x8c, 0xca, 0x51, 0xb2, 0x28, 0x88, 0x4d, 0x2a, 0xb9,
	0x90, 0x33, 0x87, 0xc8, 0xda, 0x63, 0x32, 0xe9, 0x26, 0x9b, 0x54, 0xde, 0x4c, 0x48, 0xad, 0x72,
	0xd6, 0x09, 0xb8, 0x0a, 0xe7, 0x1f, 0x04, 0x06, 0xcd, 0x33, 0xee, 0xd1, 0xbd, 0xa9, 0x14, 0xe9,
	0x0c, 0x06, 0x37, 0x58, 0x34, 0x7c, 0xba, 0x9e, 0x4b, 0x3a, 0xe4, 0xa1, 0x5e, 0xfe, 0xab, 0xdb,
	0xd1, 0xf1, 0x16, 0xdf, 0x6a, 0xe3, 0x0a, 0xf6, 0x2a, 0x1a, 0xca, 0xa6, 0x87, 0xf5, 0xe4, 0x8f,
	0xea, 0xff, 0xf8, 0x7c, 0x7a, 0x06, 0x71, 0x6a, 0x34, 0x4f, 0xfd, 0x33, 0xda, 0x4b, 0x6e, 0x1d,
	0x6a, 0x85, 0xce, 0x3a, 0x94, 0x2a, 0x2d, 0x8c, 0xe3, 0xd5, 0xee, 0xef, 0xc8, 0x63, 0x3b, 0xac,
	0xff, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x94, 0x25, 0x9a, 0x0e, 0x02, 0x00, 0x00,
}
