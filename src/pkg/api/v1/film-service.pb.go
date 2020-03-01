// Code generated by protoc-gen-go. DO NOT EDIT.
// source: film-service.proto

package v1

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

type FilmEmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilmEmptyRequest) Reset()         { *m = FilmEmptyRequest{} }
func (m *FilmEmptyRequest) String() string { return proto.CompactTextString(m) }
func (*FilmEmptyRequest) ProtoMessage()    {}
func (*FilmEmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34fd3b300f068c9c, []int{0}
}

func (m *FilmEmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilmEmptyRequest.Unmarshal(m, b)
}
func (m *FilmEmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilmEmptyRequest.Marshal(b, m, deterministic)
}
func (m *FilmEmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilmEmptyRequest.Merge(m, src)
}
func (m *FilmEmptyRequest) XXX_Size() int {
	return xxx_messageInfo_FilmEmptyRequest.Size(m)
}
func (m *FilmEmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilmEmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilmEmptyRequest proto.InternalMessageInfo

type FilmCategory struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilmCategory) Reset()         { *m = FilmCategory{} }
func (m *FilmCategory) String() string { return proto.CompactTextString(m) }
func (*FilmCategory) ProtoMessage()    {}
func (*FilmCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_34fd3b300f068c9c, []int{1}
}

func (m *FilmCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilmCategory.Unmarshal(m, b)
}
func (m *FilmCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilmCategory.Marshal(b, m, deterministic)
}
func (m *FilmCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilmCategory.Merge(m, src)
}
func (m *FilmCategory) XXX_Size() int {
	return xxx_messageInfo_FilmCategory.Size(m)
}
func (m *FilmCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_FilmCategory.DiscardUnknown(m)
}

var xxx_messageInfo_FilmCategory proto.InternalMessageInfo

func (m *FilmCategory) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FilmCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FilmCategory) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FilmCategory) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*FilmEmptyRequest)(nil), "v1.FilmEmptyRequest")
	proto.RegisterType((*FilmCategory)(nil), "v1.FilmCategory")
}

func init() { proto.RegisterFile("film-service.proto", fileDescriptor_34fd3b300f068c9c) }

var fileDescriptor_34fd3b300f068c9c = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x46, 0x4d, 0x5c, 0xc5, 0x1d, 0x45, 0x96, 0x41, 0x21, 0xe8, 0xa5, 0xf4, 0xb4, 0x17, 0x8b,
	0xab, 0x3f, 0x61, 0x71, 0x4f, 0x9e, 0xba, 0x78, 0x96, 0xd8, 0x8e, 0x75, 0x20, 0x69, 0xda, 0x24,
	0x2d, 0xf4, 0xdf, 0x8b, 0xa9, 0x05, 0xf1, 0x36, 0xbc, 0x77, 0x98, 0xf7, 0x01, 0x7e, 0xb2, 0xb1,
	0x0f, 0x81, 0xfc, 0xc8, 0x15, 0x15, 0x9d, 0x77, 0xd1, 0xa1, 0x1c, 0x77, 0x39, 0xc2, 0xe6, 0xc0,
	0xc6, 0xbe, 0xd8, 0x2e, 0x4e, 0x25, 0xf5, 0x03, 0x85, 0x98, 0xf7, 0x70, 0xf5, 0xc3, 0xf6, 0x3a,
	0x52, 0xe3, 0xfc, 0x84, 0xd7, 0x20, 0xb9, 0x56, 0x22, 0x13, 0xdb, 0xb3, 0x52, 0x72, 0x8d, 0x08,
	0xab, 0x56, 0x5b, 0x52, 0x32, 0x13, 0xdb, 0x75, 0x99, 0x6e, 0xcc, 0xe0, 0xb2, 0xa6, 0x50, 0x79,
	0xee, 0x22, 0xbb, 0x56, 0x9d, 0x26, 0xf5, 0x17, 0xe1, 0x3d, 0xac, 0xd9, 0xea, 0x86, 0xde, 0x07,
	0x6f, 0xd4, 0x2a, 0xf9, 0x8b, 0x04, 0xde, 0xbc, 0x79, 0x3a, 0xce, 0x2f, 0xc3, 0x71, 0x0e, 0xc4,
	0x3d, 0xdc, 0xbe, 0x72, 0x88, 0x07, 0x1d, 0xbe, 0xd8, 0xb5, 0xbf, 0x25, 0x4c, 0x01, 0x6f, 0x8a,
	0x71, 0x57, 0xfc, 0x2f, 0xbe, 0xdb, 0x2c, 0x74, 0x69, 0xce, 0x4f, 0x1e, 0xc5, 0xc7, 0x79, 0x9a,
	0xf9, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x05, 0x5a, 0xfc, 0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FilmsServiceClient is the client API for FilmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilmsServiceClient interface {
	ListFashionCategories(ctx context.Context, in *FilmEmptyRequest, opts ...grpc.CallOption) (FilmsService_ListFashionCategoriesClient, error)
}

type filmsServiceClient struct {
	cc *grpc.ClientConn
}

func NewFilmsServiceClient(cc *grpc.ClientConn) FilmsServiceClient {
	return &filmsServiceClient{cc}
}

func (c *filmsServiceClient) ListFashionCategories(ctx context.Context, in *FilmEmptyRequest, opts ...grpc.CallOption) (FilmsService_ListFashionCategoriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FilmsService_serviceDesc.Streams[0], "/v1.FilmsService/ListFashionCategories", opts...)
	if err != nil {
		return nil, err
	}
	x := &filmsServiceListFashionCategoriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FilmsService_ListFashionCategoriesClient interface {
	Recv() (*FilmCategory, error)
	grpc.ClientStream
}

type filmsServiceListFashionCategoriesClient struct {
	grpc.ClientStream
}

func (x *filmsServiceListFashionCategoriesClient) Recv() (*FilmCategory, error) {
	m := new(FilmCategory)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilmsServiceServer is the server API for FilmsService service.
type FilmsServiceServer interface {
	ListFashionCategories(*FilmEmptyRequest, FilmsService_ListFashionCategoriesServer) error
}

// UnimplementedFilmsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFilmsServiceServer struct {
}

func (*UnimplementedFilmsServiceServer) ListFashionCategories(req *FilmEmptyRequest, srv FilmsService_ListFashionCategoriesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFashionCategories not implemented")
}

func RegisterFilmsServiceServer(s *grpc.Server, srv FilmsServiceServer) {
	s.RegisterService(&_FilmsService_serviceDesc, srv)
}

func _FilmsService_ListFashionCategories_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilmEmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilmsServiceServer).ListFashionCategories(m, &filmsServiceListFashionCategoriesServer{stream})
}

type FilmsService_ListFashionCategoriesServer interface {
	Send(*FilmCategory) error
	grpc.ServerStream
}

type filmsServiceListFashionCategoriesServer struct {
	grpc.ServerStream
}

func (x *filmsServiceListFashionCategoriesServer) Send(m *FilmCategory) error {
	return x.ServerStream.SendMsg(m)
}

var _FilmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FilmsService",
	HandlerType: (*FilmsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFashionCategories",
			Handler:       _FilmsService_ListFashionCategories_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "film-service.proto",
}
