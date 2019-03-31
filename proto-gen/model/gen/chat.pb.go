// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Message_MsgType int32

const (
	Message_READ    Message_MsgType = 0
	Message_MESSAGE Message_MsgType = 1
)

var Message_MsgType_name = map[int32]string{
	0: "READ",
	1: "MESSAGE",
}

var Message_MsgType_value = map[string]int32{
	"READ":    0,
	"MESSAGE": 1,
}

func (x Message_MsgType) String() string {
	return proto.EnumName(Message_MsgType_name, int32(x))
}

func (Message_MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{6, 0}
}

type AddRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChatId               string   `protobuf:"bytes,2,opt,name=chatId,proto3" json:"chatId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

type LatestMessagesRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatestMessagesRequest) Reset()         { *m = LatestMessagesRequest{} }
func (m *LatestMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*LatestMessagesRequest) ProtoMessage()    {}
func (*LatestMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *LatestMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestMessagesRequest.Unmarshal(m, b)
}
func (m *LatestMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestMessagesRequest.Marshal(b, m, deterministic)
}
func (m *LatestMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestMessagesRequest.Merge(m, src)
}
func (m *LatestMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_LatestMessagesRequest.Size(m)
}
func (m *LatestMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LatestMessagesRequest proto.InternalMessageInfo

func (m *LatestMessagesRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

type PreviousMessagesRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	CurrentMessageId     int64    `protobuf:"varint,2,opt,name=currentMessageId,proto3" json:"currentMessageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousMessagesRequest) Reset()         { *m = PreviousMessagesRequest{} }
func (m *PreviousMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*PreviousMessagesRequest) ProtoMessage()    {}
func (*PreviousMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *PreviousMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreviousMessagesRequest.Unmarshal(m, b)
}
func (m *PreviousMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreviousMessagesRequest.Marshal(b, m, deterministic)
}
func (m *PreviousMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousMessagesRequest.Merge(m, src)
}
func (m *PreviousMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_PreviousMessagesRequest.Size(m)
}
func (m *PreviousMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousMessagesRequest proto.InternalMessageInfo

func (m *PreviousMessagesRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *PreviousMessagesRequest) GetCurrentMessageId() int64 {
	if m != nil {
		return m.CurrentMessageId
	}
	return 0
}

type MessageList struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (m *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(m, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type SendRequest struct {
	Message              string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId               string               `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ChatId               string               `protobuf:"bytes,3,opt,name=chatId,proto3" json:"chatId,omitempty"`
	DateTime             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *SendRequest) GetDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DateTime
	}
	return nil
}

type ReadReceipt struct {
	UserId               string               `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChatId               string               `protobuf:"bytes,2,opt,name=chatId,proto3" json:"chatId,omitempty"`
	MessageId            int64                `protobuf:"varint,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	DateTime             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReadReceipt) Reset()         { *m = ReadReceipt{} }
func (m *ReadReceipt) String() string { return proto.CompactTextString(m) }
func (*ReadReceipt) ProtoMessage()    {}
func (*ReadReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{5}
}

func (m *ReadReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadReceipt.Unmarshal(m, b)
}
func (m *ReadReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadReceipt.Marshal(b, m, deterministic)
}
func (m *ReadReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadReceipt.Merge(m, src)
}
func (m *ReadReceipt) XXX_Size() int {
	return xxx_messageInfo_ReadReceipt.Size(m)
}
func (m *ReadReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ReadReceipt proto.InternalMessageInfo

func (m *ReadReceipt) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReadReceipt) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ReadReceipt) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *ReadReceipt) GetDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DateTime
	}
	return nil
}

type Message struct {
	SenderId             string               `protobuf:"bytes,1,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Type                 Message_MsgType      `protobuf:"varint,2,opt,name=type,proto3,enum=chat.Message_MsgType" json:"type,omitempty"`
	Text                 string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	DateTime             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{6}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *Message) GetType() Message_MsgType {
	if m != nil {
		return m.Type
	}
	return Message_READ
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DateTime
	}
	return nil
}

type SubscribeRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChatId               string   `protobuf:"bytes,2,opt,name=chatId,proto3" json:"chatId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{7}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SubscribeRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func init() {
	proto.RegisterEnum("chat.Message_MsgType", Message_MsgType_name, Message_MsgType_value)
	proto.RegisterType((*AddRequest)(nil), "chat.AddRequest")
	proto.RegisterType((*LatestMessagesRequest)(nil), "chat.LatestMessagesRequest")
	proto.RegisterType((*PreviousMessagesRequest)(nil), "chat.PreviousMessagesRequest")
	proto.RegisterType((*MessageList)(nil), "chat.MessageList")
	proto.RegisterType((*SendRequest)(nil), "chat.SendRequest")
	proto.RegisterType((*ReadReceipt)(nil), "chat.ReadReceipt")
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterType((*SubscribeRequest)(nil), "chat.SubscribeRequest")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8f, 0xd2, 0x40,
	0x18, 0x66, 0x96, 0x66, 0x81, 0xb7, 0xd1, 0xb0, 0x63, 0x76, 0x25, 0x5d, 0x8d, 0xa4, 0xf1, 0xc0,
	0x7a, 0x28, 0x06, 0x8c, 0xee, 0xc1, 0x0b, 0xbb, 0x12, 0x82, 0x81, 0x64, 0x53, 0xb8, 0x7a, 0x28,
	0xed, 0x2b, 0x4c, 0x62, 0x3f, 0x9c, 0x99, 0x12, 0xf9, 0x15, 0x1e, 0xfc, 0x2b, 0x9e, 0xfd, 0x6d,
	0xa6, 0xed, 0xb4, 0xdb, 0xee, 0x87, 0x11, 0x6e, 0x7d, 0xbf, 0x9f, 0x79, 0x9e, 0x3e, 0x00, 0xee,
	0xc6, 0x91, 0x56, 0xc4, 0x43, 0x19, 0x52, 0x2d, 0xf9, 0x36, 0xce, 0xd7, 0x61, 0xb8, 0xfe, 0x86,
	0xfd, 0x34, 0xb7, 0x8a, 0xbf, 0xf6, 0xd1, 0x8f, 0xe4, 0x2e, 0x6b, 0x31, 0x5e, 0xdd, 0x2d, 0x4a,
	0xe6, 0xa3, 0x90, 0x8e, 0x1f, 0x65, 0x0d, 0xe6, 0x47, 0x80, 0x91, 0xe7, 0xd9, 0xf8, 0x3d, 0x46,
	0x21, 0xe9, 0x19, 0x1c, 0xc7, 0x02, 0xf9, 0xd4, 0xeb, 0x90, 0x2e, 0xe9, 0xb5, 0x6c, 0x15, 0x25,
	0xf9, 0xe4, 0xd6, 0xd4, 0xeb, 0x1c, 0x65, 0xf9, 0x2c, 0x32, 0xfb, 0x70, 0x3a, 0x73, 0x24, 0x0a,
	0x39, 0x47, 0x21, 0x9c, 0x35, 0x8a, 0xd2, 0x22, 0x35, 0x40, 0x2a, 0x03, 0x5f, 0xe0, 0xf9, 0x0d,
	0xc7, 0x2d, 0x0b, 0x63, 0xf1, 0x9f, 0x23, 0xf4, 0x0d, 0xb4, 0xdd, 0x98, 0x73, 0x0c, 0xf2, 0x23,
	0x0a, 0x45, 0xdd, 0xbe, 0x97, 0x37, 0x2f, 0x41, 0x57, 0xc1, 0x8c, 0x09, 0x49, 0x2f, 0xa0, 0xe9,
	0xab, 0x2b, 0x1d, 0xd2, 0xad, 0xf7, 0xf4, 0xc1, 0x13, 0x2b, 0xe5, 0x4f, 0x35, 0xd9, 0x45, 0xd9,
	0xfc, 0x49, 0x40, 0x5f, 0x60, 0x50, 0x30, 0xd1, 0x81, 0x86, 0xaa, 0x29, 0x38, 0x79, 0x58, 0xe2,
	0xe8, 0xe8, 0x11, 0x8e, 0xea, 0x15, 0xfc, 0xef, 0xa1, 0xe9, 0x39, 0x12, 0x97, 0xcc, 0xc7, 0x8e,
	0xd6, 0x25, 0x3d, 0x7d, 0x60, 0x58, 0x99, 0x2a, 0x56, 0xae, 0x8a, 0xb5, 0xcc, 0x55, 0xb1, 0x8b,
	0x5e, 0xf3, 0x17, 0x01, 0xdd, 0x46, 0xc7, 0xb3, 0xd1, 0x45, 0x16, 0xed, 0xad, 0x0d, 0x7d, 0x01,
	0x2d, 0xbf, 0x20, 0xac, 0x9e, 0x12, 0x76, 0x9b, 0x38, 0x18, 0xd5, 0x1f, 0x02, 0x0d, 0xc5, 0x1e,
	0x35, 0xa0, 0x29, 0x30, 0xf0, 0x4a, 0x98, 0x8a, 0x98, 0x5e, 0x80, 0x26, 0x77, 0x11, 0xa6, 0x98,
	0x9e, 0x0e, 0x4e, 0x2b, 0xb4, 0x5b, 0x73, 0xb1, 0x5e, 0xee, 0x22, 0xb4, 0xd3, 0x16, 0x4a, 0x41,
	0x93, 0xf8, 0x43, 0x2a, 0xda, 0xd2, 0xef, 0x83, 0xe1, 0x75, 0xa1, 0xa1, 0x96, 0xd3, 0x26, 0x68,
	0xf6, 0x78, 0xf4, 0xa9, 0x5d, 0xa3, 0x3a, 0x34, 0xe6, 0xe3, 0xc5, 0x62, 0x34, 0x19, 0xb7, 0x89,
	0x79, 0x05, 0xed, 0x45, 0xbc, 0x12, 0x2e, 0x67, 0x2b, 0x3c, 0xf0, 0xb7, 0x1f, 0xfc, 0xae, 0x83,
	0x7e, 0xbd, 0x71, 0xe4, 0x02, 0xf9, 0x96, 0xb9, 0x48, 0x2f, 0x01, 0xae, 0x39, 0x3a, 0x12, 0x93,
	0x24, 0x6d, 0x67, 0x8f, 0xbd, 0xb5, 0x95, 0x71, 0x76, 0x0f, 0xfb, 0x38, 0xf1, 0xa8, 0x59, 0x4b,
	0xde, 0xf9, 0x39, 0x64, 0xc1, 0xde, 0x73, 0x1f, 0xa0, 0x35, 0x43, 0x67, 0xbb, 0xff, 0xc1, 0x31,
	0x9c, 0x4c, 0x50, 0x56, 0x4d, 0x4b, 0xcf, 0xb3, 0x05, 0x0f, 0x5a, 0xd9, 0x38, 0xa9, 0x68, 0x97,
	0xf8, 0xca, 0xac, 0xd1, 0x29, 0x3c, 0x9b, 0xa0, 0xbc, 0x6b, 0x65, 0xfa, 0x32, 0xeb, 0x7d, 0xc4,
	0xe2, 0x0f, 0xaf, 0x1a, 0x82, 0x96, 0x18, 0x8f, 0xaa, 0x62, 0xc9, 0x84, 0xff, 0x78, 0xc6, 0x10,
	0x5a, 0x85, 0x8a, 0xf9, 0x64, 0xc9, 0x2c, 0x46, 0xd5, 0xe7, 0x66, 0xad, 0x47, 0xde, 0x92, 0xab,
	0x1e, 0xbc, 0x76, 0x43, 0xdf, 0x72, 0xc5, 0x06, 0xa3, 0x77, 0x56, 0xc4, 0xd1, 0x67, 0xc8, 0x23,
	0x8e, 0x1e, 0x73, 0x65, 0xc8, 0xb3, 0x89, 0x35, 0x06, 0x37, 0x64, 0x75, 0x9c, 0x1e, 0x1c, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x68, 0x10, 0xfb, 0x2f, 0x6e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	JoinChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	LeaveChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetLatestMessages(ctx context.Context, in *LatestMessagesRequest, opts ...grpc.CallOption) (*MessageList, error)
	GetPreviousMessages(ctx context.Context, in *PreviousMessagesRequest, opts ...grpc.CallOption) (*MessageList, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (ChatService_SubscribeClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/JoinChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveChat(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/LeaveChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetLatestMessages(ctx context.Context, in *LatestMessagesRequest, opts ...grpc.CallOption) (*MessageList, error) {
	out := new(MessageList)
	err := c.cc.Invoke(ctx, "/chat.ChatService/GetLatestMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetPreviousMessages(ctx context.Context, in *PreviousMessagesRequest, opts ...grpc.CallOption) (*MessageList, error) {
	out := new(MessageList)
	err := c.cc.Invoke(ctx, "/chat.ChatService/GetPreviousMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (ChatService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeClient{stream}
	return x, nil
}

type ChatService_SubscribeClient interface {
	Send(*ReadReceipt) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeClient) Send(m *ReadReceipt) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	CreateChat(context.Context, *AddRequest) (*empty.Empty, error)
	JoinChat(context.Context, *AddRequest) (*empty.Empty, error)
	LeaveChat(context.Context, *AddRequest) (*empty.Empty, error)
	GetLatestMessages(context.Context, *LatestMessagesRequest) (*MessageList, error)
	GetPreviousMessages(context.Context, *PreviousMessagesRequest) (*MessageList, error)
	Send(context.Context, *SendRequest) (*empty.Empty, error)
	Subscribe(ChatService_SubscribeServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/JoinChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinChat(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeaveChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeaveChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/LeaveChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeaveChat(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetLatestMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetLatestMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/GetLatestMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetLatestMessages(ctx, req.(*LatestMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetPreviousMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviousMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetPreviousMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/GetPreviousMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetPreviousMessages(ctx, req.(*PreviousMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Subscribe(&chatServiceSubscribeServer{stream})
}

type ChatService_SubscribeServer interface {
	Send(*Message) error
	Recv() (*ReadReceipt, error)
	grpc.ServerStream
}

type chatServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSubscribeServer) Recv() (*ReadReceipt, error) {
	m := new(ReadReceipt)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "JoinChat",
			Handler:    _ChatService_JoinChat_Handler,
		},
		{
			MethodName: "LeaveChat",
			Handler:    _ChatService_LeaveChat_Handler,
		},
		{
			MethodName: "GetLatestMessages",
			Handler:    _ChatService_GetLatestMessages_Handler,
		},
		{
			MethodName: "GetPreviousMessages",
			Handler:    _ChatService_GetPreviousMessages_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _ChatService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ChatService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
