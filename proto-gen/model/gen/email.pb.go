// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type SendEmailRequest struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	SenderEmail          string   `protobuf:"bytes,3,opt,name=senderEmail,proto3" json:"senderEmail,omitempty"`
	RecipientEmail       string   `protobuf:"bytes,4,opt,name=recipientEmail,proto3" json:"recipientEmail,omitempty"`
	Subject              string   `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailRequest) Reset()         { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()    {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_email_e26ad237bf868118, []int{0}
}
func (m *SendEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailRequest.Unmarshal(m, b)
}
func (m *SendEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailRequest.Marshal(b, m, deterministic)
}
func (dst *SendEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailRequest.Merge(dst, src)
}
func (m *SendEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendEmailRequest.Size(m)
}
func (m *SendEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailRequest proto.InternalMessageInfo

func (m *SendEmailRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SendEmailRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *SendEmailRequest) GetSenderEmail() string {
	if m != nil {
		return m.SenderEmail
	}
	return ""
}

func (m *SendEmailRequest) GetRecipientEmail() string {
	if m != nil {
		return m.RecipientEmail
	}
	return ""
}

func (m *SendEmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendEmailRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*SendEmailRequest)(nil), "model.SendEmailRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Send(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Send(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.EmailService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Send(context.Context, *SendEmailRequest) (*empty.Empty, error)
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}

func init() { proto.RegisterFile("email.proto", fileDescriptor_email_e26ad237bf868118) }

var fileDescriptor_email_e26ad237bf868118 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0x3b, 0x31,
	0x18, 0xc4, 0xff, 0xfb, 0xb7, 0x5d, 0xe9, 0x57, 0x11, 0xc9, 0xa1, 0x86, 0xea, 0xa1, 0x16, 0x14,
	0x4f, 0x29, 0xa8, 0x17, 0xaf, 0x42, 0x0f, 0xde, 0xa4, 0x7d, 0x02, 0x37, 0x3b, 0xd6, 0xc8, 0x66,
	0x13, 0xb3, 0x59, 0xc1, 0x17, 0xf4, 0xb9, 0x24, 0x5f, 0x5a, 0x95, 0x1e, 0x67, 0xe6, 0x37, 0x09,
	0xdf, 0xd0, 0x18, 0xf6, 0xd9, 0x34, 0xca, 0x07, 0x17, 0x9d, 0x18, 0x5a, 0x57, 0xa3, 0x99, 0x9e,
	0x6d, 0x9c, 0xdb, 0x34, 0x58, 0xb0, 0x59, 0xf5, 0x2f, 0x0b, 0x58, 0x1f, 0x3f, 0x33, 0x33, 0xff,
	0x2a, 0xe8, 0x64, 0x8d, 0xb6, 0x5e, 0xa6, 0xde, 0x0a, 0xef, 0x3d, 0xba, 0x28, 0x26, 0x54, 0x76,
	0x68, 0x6b, 0x04, 0x59, 0xcc, 0x8a, 0xeb, 0xd1, 0x6a, 0xab, 0xc4, 0x39, 0x8d, 0x02, 0xb4, 0xf1,
	0x06, 0x6d, 0x94, 0xff, 0x39, 0xfa, 0x35, 0xc4, 0x8c, 0xc6, 0x99, 0xe3, 0xb7, 0xe4, 0x01, 0xe7,
	0x7f, 0x2d, 0x71, 0x45, 0xc7, 0x3f, 0x78, 0x86, 0x06, 0x0c, 0xed, 0xb9, 0x42, 0xd2, 0x61, 0xd7,
	0x57, 0x6f, 0xd0, 0x51, 0x0e, 0x19, 0xd8, 0xc9, 0x94, 0x68, 0xd7, 0xc6, 0xf4, 0x7f, 0x99, 0x93,
	0xad, 0xbc, 0x79, 0xa4, 0x23, 0x2e, 0xaf, 0x11, 0x3e, 0x8c, 0x86, 0xb8, 0xa7, 0x41, 0xba, 0x4b,
	0x9c, 0x2a, 0x5e, 0x41, 0xed, 0x1f, 0x39, 0x9d, 0xa8, 0xbc, 0x8b, 0xda, 0xed, 0xa2, 0x96, 0x69,
	0x97, 0xf9, 0xbf, 0x87, 0x4b, 0xba, 0xd0, 0xce, 0x2a, 0xdd, 0xbd, 0xc2, 0xdf, 0x29, 0x1f, 0x60,
	0x0d, 0x82, 0x0f, 0xa8, 0x8d, 0x8e, 0x2e, 0x28, 0x9e, 0xf8, 0xa9, 0xa8, 0x4a, 0x2e, 0xde, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x46, 0xeb, 0x8d, 0xb3, 0x74, 0x01, 0x00, 0x00,
}
