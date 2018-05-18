// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movies.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type GetMoviesResponse struct {
	Movies               []*Movie `protobuf:"bytes,1,rep,name=movies" json:"movies,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMoviesResponse) Reset()         { *m = GetMoviesResponse{} }
func (m *GetMoviesResponse) String() string { return proto.CompactTextString(m) }
func (*GetMoviesResponse) ProtoMessage()    {}
func (*GetMoviesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{0}
}
func (m *GetMoviesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMoviesResponse.Unmarshal(m, b)
}
func (m *GetMoviesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMoviesResponse.Marshal(b, m, deterministic)
}
func (dst *GetMoviesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMoviesResponse.Merge(dst, src)
}
func (m *GetMoviesResponse) XXX_Size() int {
	return xxx_messageInfo_GetMoviesResponse.Size(m)
}
func (m *GetMoviesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMoviesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMoviesResponse proto.InternalMessageInfo

func (m *GetMoviesResponse) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

func (m *GetMoviesResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Movie struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Director             []*Director          `protobuf:"bytes,3,rep,name=director" json:"director,omitempty"`
	Year                 string               `protobuf:"bytes,4,opt,name=year" json:"year,omitempty"`
	Userid               string               `protobuf:"bytes,5,opt,name=userid" json:"userid,omitempty"`
	Createdon            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdon" json:"createdon,omitempty"`
	Updatedon            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updatedon" json:"updatedon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{1}
}
func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (dst *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(dst, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetDirector() []*Director {
	if m != nil {
		return m.Director
	}
	return nil
}

func (m *Movie) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Movie) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *Movie) GetCreatedon() *timestamp.Timestamp {
	if m != nil {
		return m.Createdon
	}
	return nil
}

func (m *Movie) GetUpdatedon() *timestamp.Timestamp {
	if m != nil {
		return m.Updatedon
	}
	return nil
}

type Director struct {
	Director             string   `protobuf:"bytes,1,opt,name=director" json:"director,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Director) Reset()         { *m = Director{} }
func (m *Director) String() string { return proto.CompactTextString(m) }
func (*Director) ProtoMessage()    {}
func (*Director) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{2}
}
func (m *Director) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Director.Unmarshal(m, b)
}
func (m *Director) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Director.Marshal(b, m, deterministic)
}
func (dst *Director) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Director.Merge(dst, src)
}
func (m *Director) XXX_Size() int {
	return xxx_messageInfo_Director.Size(m)
}
func (m *Director) XXX_DiscardUnknown() {
	xxx_messageInfo_Director.DiscardUnknown(m)
}

var xxx_messageInfo_Director proto.InternalMessageInfo

func (m *Director) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

type GetMovieByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMovieByIdRequest) Reset()         { *m = GetMovieByIdRequest{} }
func (m *GetMovieByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetMovieByIdRequest) ProtoMessage()    {}
func (*GetMovieByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{3}
}
func (m *GetMovieByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMovieByIdRequest.Unmarshal(m, b)
}
func (m *GetMovieByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMovieByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetMovieByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMovieByIdRequest.Merge(dst, src)
}
func (m *GetMovieByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetMovieByIdRequest.Size(m)
}
func (m *GetMovieByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMovieByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMovieByIdRequest proto.InternalMessageInfo

func (m *GetMovieByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetMovieByIdResponse struct {
	Movie                *Movie   `protobuf:"bytes,1,opt,name=movie" json:"movie,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMovieByIdResponse) Reset()         { *m = GetMovieByIdResponse{} }
func (m *GetMovieByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetMovieByIdResponse) ProtoMessage()    {}
func (*GetMovieByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{4}
}
func (m *GetMovieByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMovieByIdResponse.Unmarshal(m, b)
}
func (m *GetMovieByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMovieByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetMovieByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMovieByIdResponse.Merge(dst, src)
}
func (m *GetMovieByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetMovieByIdResponse.Size(m)
}
func (m *GetMovieByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMovieByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMovieByIdResponse proto.InternalMessageInfo

func (m *GetMovieByIdResponse) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *GetMovieByIdResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type NewMovieRequest struct {
	Title                string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Director             []*Director `protobuf:"bytes,2,rep,name=director" json:"director,omitempty"`
	Year                 string      `protobuf:"bytes,3,opt,name=year" json:"year,omitempty"`
	Userid               string      `protobuf:"bytes,4,opt,name=userid" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NewMovieRequest) Reset()         { *m = NewMovieRequest{} }
func (m *NewMovieRequest) String() string { return proto.CompactTextString(m) }
func (*NewMovieRequest) ProtoMessage()    {}
func (*NewMovieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{5}
}
func (m *NewMovieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMovieRequest.Unmarshal(m, b)
}
func (m *NewMovieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMovieRequest.Marshal(b, m, deterministic)
}
func (dst *NewMovieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMovieRequest.Merge(dst, src)
}
func (m *NewMovieRequest) XXX_Size() int {
	return xxx_messageInfo_NewMovieRequest.Size(m)
}
func (m *NewMovieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMovieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewMovieRequest proto.InternalMessageInfo

func (m *NewMovieRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewMovieRequest) GetDirector() []*Director {
	if m != nil {
		return m.Director
	}
	return nil
}

func (m *NewMovieRequest) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *NewMovieRequest) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type NewMovieResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMovieResponse) Reset()         { *m = NewMovieResponse{} }
func (m *NewMovieResponse) String() string { return proto.CompactTextString(m) }
func (*NewMovieResponse) ProtoMessage()    {}
func (*NewMovieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{6}
}
func (m *NewMovieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMovieResponse.Unmarshal(m, b)
}
func (m *NewMovieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMovieResponse.Marshal(b, m, deterministic)
}
func (dst *NewMovieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMovieResponse.Merge(dst, src)
}
func (m *NewMovieResponse) XXX_Size() int {
	return xxx_messageInfo_NewMovieResponse.Size(m)
}
func (m *NewMovieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMovieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewMovieResponse proto.InternalMessageInfo

func (m *NewMovieResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NewMovieResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_movies_29668afc7e8789eb, []int{7}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetMoviesResponse)(nil), "pb.GetMoviesResponse")
	proto.RegisterType((*Movie)(nil), "pb.Movie")
	proto.RegisterType((*Director)(nil), "pb.Director")
	proto.RegisterType((*GetMovieByIdRequest)(nil), "pb.GetMovieByIdRequest")
	proto.RegisterType((*GetMovieByIdResponse)(nil), "pb.GetMovieByIdResponse")
	proto.RegisterType((*NewMovieRequest)(nil), "pb.NewMovieRequest")
	proto.RegisterType((*NewMovieResponse)(nil), "pb.NewMovieResponse")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Movies service

type MoviesClient interface {
	GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMoviesResponse, error)
	GetMovieById(ctx context.Context, in *GetMovieByIdRequest, opts ...grpc.CallOption) (*GetMovieByIdResponse, error)
	NewMovie(ctx context.Context, in *NewMovieRequest, opts ...grpc.CallOption) (*NewMovieResponse, error)
}

type moviesClient struct {
	cc *grpc.ClientConn
}

func NewMoviesClient(cc *grpc.ClientConn) MoviesClient {
	return &moviesClient{cc}
}

func (c *moviesClient) GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMoviesResponse, error) {
	out := new(GetMoviesResponse)
	err := grpc.Invoke(ctx, "/pb.Movies/GetMovies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) GetMovieById(ctx context.Context, in *GetMovieByIdRequest, opts ...grpc.CallOption) (*GetMovieByIdResponse, error) {
	out := new(GetMovieByIdResponse)
	err := grpc.Invoke(ctx, "/pb.Movies/GetMovieById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) NewMovie(ctx context.Context, in *NewMovieRequest, opts ...grpc.CallOption) (*NewMovieResponse, error) {
	out := new(NewMovieResponse)
	err := grpc.Invoke(ctx, "/pb.Movies/NewMovie", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movies service

type MoviesServer interface {
	GetMovies(context.Context, *Empty) (*GetMoviesResponse, error)
	GetMovieById(context.Context, *GetMovieByIdRequest) (*GetMovieByIdResponse, error)
	NewMovie(context.Context, *NewMovieRequest) (*NewMovieResponse, error)
}

func RegisterMoviesServer(s *grpc.Server, srv MoviesServer) {
	s.RegisterService(&_Movies_serviceDesc, srv)
}

func _Movies_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Movies/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetMovies(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_GetMovieById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetMovieById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Movies/GetMovieById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetMovieById(ctx, req.(*GetMovieByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_NewMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).NewMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Movies/NewMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).NewMovie(ctx, req.(*NewMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Movies_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Movies",
	HandlerType: (*MoviesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovies",
			Handler:    _Movies_GetMovies_Handler,
		},
		{
			MethodName: "GetMovieById",
			Handler:    _Movies_GetMovieById_Handler,
		},
		{
			MethodName: "NewMovie",
			Handler:    _Movies_NewMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movies.proto",
}

func init() { proto.RegisterFile("movies.proto", fileDescriptor_movies_29668afc7e8789eb) }

var fileDescriptor_movies_29668afc7e8789eb = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0xef, 0x93, 0x40,
	0x10, 0xc6, 0x05, 0x0a, 0xff, 0x32, 0x6d, 0xb4, 0x4e, 0xab, 0x6e, 0xb8, 0xb4, 0x92, 0x68, 0x38,
	0xd1, 0xa4, 0x9a, 0xe8, 0xd9, 0x97, 0x68, 0x0f, 0x7a, 0x20, 0x7e, 0x81, 0x52, 0xc6, 0x86, 0xa4,
	0x74, 0x71, 0x59, 0x34, 0x3d, 0xf8, 0xb9, 0xfc, 0x68, 0x5e, 0x0d, 0xbb, 0x0b, 0xa5, 0x2f, 0x46,
	0x6f, 0xbb, 0x33, 0xbf, 0xd9, 0xd9, 0x79, 0x9e, 0x81, 0x71, 0xc1, 0xbf, 0xe7, 0x54, 0xc5, 0xa5,
	0xe0, 0x92, 0xa3, 0x5d, 0xa6, 0xc1, 0x7c, 0xc7, 0xf9, 0x6e, 0x4f, 0x4b, 0x15, 0x49, 0xeb, 0xaf,
	0x4b, 0x99, 0x17, 0x54, 0xc9, 0x4d, 0x51, 0x6a, 0x28, 0xfc, 0x08, 0x0f, 0x3f, 0x90, 0xfc, 0xa4,
	0xea, 0x12, 0xaa, 0x4a, 0x7e, 0xa8, 0x08, 0x9f, 0x82, 0xa7, 0x5f, 0x62, 0xd6, 0xc2, 0x89, 0x46,
	0x2b, 0x3f, 0x2e, 0xd3, 0x58, 0x31, 0x89, 0x49, 0xe0, 0x04, 0x1c, 0x12, 0x82, 0xd9, 0x0b, 0x2b,
	0xf2, 0x93, 0xe6, 0x18, 0xfe, 0xb6, 0xc0, 0x55, 0x0c, 0xde, 0x07, 0x3b, 0xcf, 0x98, 0xa5, 0x52,
	0x76, 0x9e, 0xe1, 0x0c, 0x5c, 0x99, 0xcb, 0x3d, 0x19, 0x5a, 0x5f, 0x30, 0x82, 0x61, 0x96, 0x0b,
	0xda, 0x4a, 0x2e, 0x98, 0xa3, 0xda, 0x8c, 0x9b, 0x36, 0xef, 0x4c, 0x2c, 0xe9, 0xb2, 0x88, 0x30,
	0x38, 0xd2, 0x46, 0xb0, 0x81, 0x2a, 0x57, 0x67, 0x7c, 0x0c, 0x5e, 0x5d, 0x91, 0xc8, 0x33, 0xe6,
	0xaa, 0xa8, 0xb9, 0xe1, 0x6b, 0xf0, 0xb7, 0x82, 0x36, 0x92, 0x32, 0x7e, 0x60, 0xde, 0xc2, 0x8a,
	0x46, 0xab, 0x20, 0xd6, 0x22, 0xc4, 0xad, 0x08, 0xf1, 0x97, 0x56, 0x84, 0xe4, 0x04, 0x37, 0x95,
	0x75, 0x99, 0x99, 0xca, 0xbb, 0x7f, 0x57, 0x76, 0x70, 0xf8, 0x1c, 0x86, 0xed, 0xaf, 0x31, 0xe8,
	0x4d, 0xa5, 0x15, 0xe8, 0xee, 0xe1, 0x33, 0x98, 0xb6, 0x5a, 0xbf, 0x39, 0xae, 0xb3, 0x84, 0xbe,
	0xd5, 0x54, 0xc9, 0x4b, 0xb9, 0xc2, 0x35, 0xcc, 0xce, 0x31, 0xe3, 0xca, 0x1c, 0x5c, 0x25, 0xbe,
	0x42, 0xcf, 0x4c, 0xd1, 0xf1, 0x1b, 0x9e, 0xfc, 0x84, 0x07, 0x9f, 0xe9, 0x87, 0x86, 0x4c, 0xb7,
	0xce, 0x0c, 0xeb, 0x6f, 0x66, 0xd8, 0xff, 0x65, 0x86, 0x73, 0xd3, 0x8c, 0x41, 0xdf, 0x8c, 0xf0,
	0x25, 0x4c, 0x4e, 0xed, 0xcd, 0x14, 0x97, 0xcb, 0x71, 0xfd, 0xe9, 0x3b, 0x70, 0xdf, 0x17, 0xa5,
	0x3c, 0xae, 0x7e, 0x59, 0xe0, 0xe9, 0xcd, 0xc4, 0x25, 0xf8, 0xdd, 0x9a, 0xa2, 0x9a, 0x5c, 0x21,
	0xc1, 0xa3, 0xe6, 0x78, 0xb5, 0xc0, 0xe1, 0x3d, 0x7c, 0x0b, 0xe3, 0xbe, 0x88, 0xf8, 0xa4, 0x0f,
	0xf6, 0xd4, 0x0f, 0xd8, 0x75, 0xa2, 0x7b, 0xe4, 0x15, 0x0c, 0xdb, 0xff, 0xe3, 0xb4, 0xe1, 0x2e,
	0xc4, 0x0c, 0x66, 0xe7, 0xc1, 0xb6, 0x30, 0xf5, 0xd4, 0xc2, 0xbc, 0xf8, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0xb5, 0xb3, 0x0e, 0x91, 0x03, 0x00, 0x00,
}
