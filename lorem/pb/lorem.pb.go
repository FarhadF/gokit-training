// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lorem.proto

package pb

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

type LoremRequest struct {
	RequestType          string   `protobuf:"bytes,1,opt,name=requestType" json:"requestType,omitempty"`
	Min                  int32    `protobuf:"varint,2,opt,name=min" json:"min,omitempty"`
	Max                  int32    `protobuf:"varint,3,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoremRequest) Reset()         { *m = LoremRequest{} }
func (m *LoremRequest) String() string { return proto.CompactTextString(m) }
func (*LoremRequest) ProtoMessage()    {}
func (*LoremRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lorem_296e33bcc78b35fb, []int{0}
}
func (m *LoremRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoremRequest.Unmarshal(m, b)
}
func (m *LoremRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoremRequest.Marshal(b, m, deterministic)
}
func (dst *LoremRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoremRequest.Merge(dst, src)
}
func (m *LoremRequest) XXX_Size() int {
	return xxx_messageInfo_LoremRequest.Size(m)
}
func (m *LoremRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoremRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoremRequest proto.InternalMessageInfo

func (m *LoremRequest) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *LoremRequest) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *LoremRequest) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type LoremResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoremResponse) Reset()         { *m = LoremResponse{} }
func (m *LoremResponse) String() string { return proto.CompactTextString(m) }
func (*LoremResponse) ProtoMessage()    {}
func (*LoremResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lorem_296e33bcc78b35fb, []int{1}
}
func (m *LoremResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoremResponse.Unmarshal(m, b)
}
func (m *LoremResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoremResponse.Marshal(b, m, deterministic)
}
func (dst *LoremResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoremResponse.Merge(dst, src)
}
func (m *LoremResponse) XXX_Size() int {
	return xxx_messageInfo_LoremResponse.Size(m)
}
func (m *LoremResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoremResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoremResponse proto.InternalMessageInfo

func (m *LoremResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoremResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*LoremRequest)(nil), "pb.LoremRequest")
	proto.RegisterType((*LoremResponse)(nil), "pb.LoremResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Lorem service

type LoremClient interface {
	Lorem(ctx context.Context, in *LoremRequest, opts ...grpc.CallOption) (*LoremResponse, error)
}

type loremClient struct {
	cc *grpc.ClientConn
}

func NewLoremClient(cc *grpc.ClientConn) LoremClient {
	return &loremClient{cc}
}

func (c *loremClient) Lorem(ctx context.Context, in *LoremRequest, opts ...grpc.CallOption) (*LoremResponse, error) {
	out := new(LoremResponse)
	err := grpc.Invoke(ctx, "/pb.Lorem/Lorem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lorem service

type LoremServer interface {
	Lorem(context.Context, *LoremRequest) (*LoremResponse, error)
}

func RegisterLoremServer(s *grpc.Server, srv LoremServer) {
	s.RegisterService(&_Lorem_serviceDesc, srv)
}

func _Lorem_Lorem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoremRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoremServer).Lorem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Lorem/Lorem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoremServer).Lorem(ctx, req.(*LoremRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lorem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Lorem",
	HandlerType: (*LoremServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lorem",
			Handler:    _Lorem_Lorem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorem.proto",
}

func init() { proto.RegisterFile("lorem.proto", fileDescriptor_lorem_296e33bcc78b35fb) }

var fileDescriptor_lorem_296e33bcc78b35fb = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x2f, 0x4a,
	0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x0a, 0xe1, 0xe2, 0xf1,
	0x01, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x29, 0x70, 0x71, 0x17, 0x41, 0x98,
	0x21, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x42, 0x02, 0x5c,
	0xcc, 0xb9, 0x99, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x20, 0x26, 0x58, 0x24, 0xb1,
	0x42, 0x82, 0x19, 0x2a, 0x92, 0x58, 0xa1, 0x64, 0xcd, 0xc5, 0x0b, 0x35, 0xb5, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0x66, 0x24, 0x8c,
	0x0b, 0xd2, 0x9c, 0x5a, 0x54, 0x04, 0x36, 0x8e, 0x33, 0x08, 0xc4, 0x34, 0x32, 0xe7, 0x62, 0x05,
	0x6b, 0x16, 0xd2, 0x83, 0x31, 0x04, 0xf4, 0x0a, 0x92, 0xf4, 0x90, 0x9d, 0x29, 0x25, 0x88, 0x24,
	0x02, 0xb1, 0x42, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x2d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x15, 0xb5, 0xf1, 0x5e, 0xe5, 0x00, 0x00, 0x00,
}
