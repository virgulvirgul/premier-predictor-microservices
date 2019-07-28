// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prediction.proto

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

type Prediction struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MatchId              string   `protobuf:"bytes,2,opt,name=matchId,proto3" json:"matchId,omitempty"`
	HGoals               int32    `protobuf:"varint,3,opt,name=hGoals,proto3" json:"hGoals,omitempty"`
	AGoals               int32    `protobuf:"varint,4,opt,name=aGoals,proto3" json:"aGoals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prediction) Reset()         { *m = Prediction{} }
func (m *Prediction) String() string { return proto.CompactTextString(m) }
func (*Prediction) ProtoMessage()    {}
func (*Prediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_58557f179e517d63, []int{0}
}
func (m *Prediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prediction.Unmarshal(m, b)
}
func (m *Prediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prediction.Marshal(b, m, deterministic)
}
func (dst *Prediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prediction.Merge(dst, src)
}
func (m *Prediction) XXX_Size() int {
	return xxx_messageInfo_Prediction.Size(m)
}
func (m *Prediction) XXX_DiscardUnknown() {
	xxx_messageInfo_Prediction.DiscardUnknown(m)
}

var xxx_messageInfo_Prediction proto.InternalMessageInfo

func (m *Prediction) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Prediction) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func (m *Prediction) GetHGoals() int32 {
	if m != nil {
		return m.HGoals
	}
	return 0
}

func (m *Prediction) GetAGoals() int32 {
	if m != nil {
		return m.AGoals
	}
	return 0
}

type MatchPredictionSummary struct {
	HomeWin              int32    `protobuf:"varint,1,opt,name=homeWin,proto3" json:"homeWin,omitempty"`
	Draw                 int32    `protobuf:"varint,2,opt,name=draw,proto3" json:"draw,omitempty"`
	AwayWin              int32    `protobuf:"varint,3,opt,name=awayWin,proto3" json:"awayWin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchPredictionSummary) Reset()         { *m = MatchPredictionSummary{} }
func (m *MatchPredictionSummary) String() string { return proto.CompactTextString(m) }
func (*MatchPredictionSummary) ProtoMessage()    {}
func (*MatchPredictionSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_58557f179e517d63, []int{1}
}
func (m *MatchPredictionSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredictionSummary.Unmarshal(m, b)
}
func (m *MatchPredictionSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredictionSummary.Marshal(b, m, deterministic)
}
func (dst *MatchPredictionSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredictionSummary.Merge(dst, src)
}
func (m *MatchPredictionSummary) XXX_Size() int {
	return xxx_messageInfo_MatchPredictionSummary.Size(m)
}
func (m *MatchPredictionSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredictionSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredictionSummary proto.InternalMessageInfo

func (m *MatchPredictionSummary) GetHomeWin() int32 {
	if m != nil {
		return m.HomeWin
	}
	return 0
}

func (m *MatchPredictionSummary) GetDraw() int32 {
	if m != nil {
		return m.Draw
	}
	return 0
}

func (m *MatchPredictionSummary) GetAwayWin() int32 {
	if m != nil {
		return m.AwayWin
	}
	return 0
}

func init() {
	proto.RegisterType((*Prediction)(nil), "model.Prediction")
	proto.RegisterType((*MatchPredictionSummary)(nil), "model.MatchPredictionSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*Prediction, error)
	GetPredictionSummary(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MatchPredictionSummary, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*Prediction, error) {
	out := new(Prediction)
	err := c.cc.Invoke(ctx, "/model.PredictionService/GetPrediction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetPredictionSummary(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MatchPredictionSummary, error) {
	out := new(MatchPredictionSummary)
	err := c.cc.Invoke(ctx, "/model.PredictionService/GetPredictionSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	GetPrediction(context.Context, *PredictionRequest) (*Prediction, error)
	GetPredictionSummary(context.Context, *IdRequest) (*MatchPredictionSummary, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_GetPrediction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetPrediction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.PredictionService/GetPrediction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetPrediction(ctx, req.(*PredictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetPredictionSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetPredictionSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.PredictionService/GetPredictionSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetPredictionSummary(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrediction",
			Handler:    _PredictionService_GetPrediction_Handler,
		},
		{
			MethodName: "GetPredictionSummary",
			Handler:    _PredictionService_GetPredictionSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prediction.proto",
}

func init() { proto.RegisterFile("prediction.proto", fileDescriptor_prediction_58557f179e517d63) }

var fileDescriptor_prediction_58557f179e517d63 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xb4, 0xa9, 0x38, 0x50, 0x68, 0x17, 0x29, 0xa1, 0x20, 0x94, 0x1c, 0xa4, 0x78, 0xd8,
	0x83, 0x7a, 0xf6, 0xe0, 0xa5, 0xe4, 0x20, 0x94, 0x78, 0xf0, 0xea, 0x9a, 0x1d, 0x48, 0xa0, 0xdb,
	0x8d, 0x93, 0x8d, 0xa5, 0x7f, 0xc6, 0xdf, 0x2a, 0xfb, 0x65, 0x2a, 0x7a, 0xcb, 0x7b, 0xf3, 0x5e,
	0xde, 0xce, 0x3c, 0x98, 0xb5, 0x84, 0xb2, 0xa9, 0x4c, 0xa3, 0xf7, 0xbc, 0x25, 0x6d, 0x34, 0x4b,
	0x95, 0x96, 0xb8, 0x5b, 0x4e, 0x09, 0x3f, 0x7a, 0xec, 0x8c, 0x67, 0xf3, 0x3d, 0xc0, 0xf6, 0x47,
	0xc9, 0x16, 0x30, 0xe9, 0x3b, 0xa4, 0x42, 0x66, 0xc9, 0x2a, 0x59, 0x5f, 0x96, 0x01, 0xb1, 0x0c,
	0x2e, 0x94, 0x30, 0x55, 0x5d, 0xc8, 0xec, 0xcc, 0x0d, 0x22, 0xb4, 0x8e, 0x7a, 0xa3, 0xc5, 0xae,
	0xcb, 0xce, 0x57, 0xc9, 0x3a, 0x2d, 0x03, 0xb2, 0xbc, 0xf0, 0xfc, 0xd8, 0xf3, 0x1e, 0xe5, 0x6f,
	0xb0, 0x78, 0xb6, 0xd6, 0x21, 0xf4, 0xa5, 0x57, 0x4a, 0xd0, 0xd1, 0x66, 0xd4, 0x5a, 0xe1, 0x6b,
	0xb3, 0x77, 0xe1, 0x69, 0x19, 0x21, 0x63, 0x30, 0x96, 0x24, 0x0e, 0x2e, 0x3a, 0x2d, 0xdd, 0xb7,
	0x55, 0x8b, 0x83, 0x38, 0x5a, 0xb5, 0x0f, 0x8e, 0xf0, 0xee, 0x2b, 0x81, 0xf9, 0xc9, 0xdf, 0x91,
	0x3e, 0x9b, 0x0a, 0xd9, 0x23, 0x4c, 0x37, 0x68, 0x4e, 0x56, 0xcd, 0xb8, 0xbb, 0x07, 0x1f, 0xa8,
	0xd2, 0x1f, 0x66, 0x39, 0xff, 0x33, 0xc9, 0x47, 0xac, 0x80, 0xab, 0x5f, 0xfe, 0xf8, 0xea, 0x59,
	0x10, 0x17, 0x32, 0xda, 0xaf, 0x03, 0xf3, 0xff, 0x9a, 0xf9, 0xe8, 0xe9, 0x16, 0x6e, 0x2a, 0xad,
	0x78, 0xd5, 0xd5, 0xd8, 0x3e, 0xf0, 0x96, 0x50, 0x35, 0x48, 0xa1, 0x2e, 0x4d, 0x7c, 0x28, 0x6e,
	0x9b, 0xbc, 0x4f, 0x5c, 0x4b, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x84, 0x91, 0x18,
	0xcf, 0x01, 0x00, 0x00,
}