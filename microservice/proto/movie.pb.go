//*
//customer is a service which holds the functionalites of handling the customers who are registering the application

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: movie.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie string `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	Page  string `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListMovieRequest) Reset() {
	*x = ListMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMovieRequest) ProtoMessage() {}

func (x *ListMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMovieRequest.ProtoReflect.Descriptor instead.
func (*ListMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *ListMovieRequest) GetMovie() string {
	if x != nil {
		return x.Movie
	}
	return ""
}

func (x *ListMovieRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type MovieItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year     string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	MovieId  string `protobuf:"bytes,3,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	MovieUrl string `protobuf:"bytes,5,opt,name=movie_url,json=movieUrl,proto3" json:"movie_url,omitempty"`
}

func (x *MovieItem) Reset() {
	*x = MovieItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieItem) ProtoMessage() {}

func (x *MovieItem) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieItem.ProtoReflect.Descriptor instead.
func (*MovieItem) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieItem) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieItem) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *MovieItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieItem) GetMovieUrl() string {
	if x != nil {
		return x.MovieUrl
	}
	return ""
}

type MovieList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*MovieItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Page string       `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *MovieList) Reset() {
	*x = MovieList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieList) ProtoMessage() {}

func (x *MovieList) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieList.ProtoReflect.Descriptor instead.
func (*MovieList) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *MovieList) GetList() []*MovieItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *MovieList) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x45, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x49, 0x0a,
	0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x39, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x62, 0x69, 0x74, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_movie_proto_goTypes = []interface{}{
	(*ListMovieRequest)(nil), // 0: proto.ListMovieRequest
	(*MovieItem)(nil),        // 1: proto.MovieItem
	(*MovieList)(nil),        // 2: proto.MovieList
}
var file_movie_proto_depIdxs = []int32{
	1, // 0: proto.MovieList.list:type_name -> proto.MovieItem
	0, // 1: proto.MovieHandler.GetListMovie:input_type -> proto.ListMovieRequest
	2, // 2: proto.MovieHandler.GetListMovie:output_type -> proto.MovieList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMovieRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieHandlerClient is the client API for MovieHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieHandlerClient interface {
	GetListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (*MovieList, error)
}

type movieHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieHandlerClient(cc grpc.ClientConnInterface) MovieHandlerClient {
	return &movieHandlerClient{cc}
}

func (c *movieHandlerClient) GetListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (*MovieList, error) {
	out := new(MovieList)
	err := c.cc.Invoke(ctx, "/proto.MovieHandler/GetListMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieHandlerServer is the server API for MovieHandler service.
type MovieHandlerServer interface {
	GetListMovie(context.Context, *ListMovieRequest) (*MovieList, error)
}

// UnimplementedMovieHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedMovieHandlerServer struct {
}

func (*UnimplementedMovieHandlerServer) GetListMovie(context.Context, *ListMovieRequest) (*MovieList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListMovie not implemented")
}

func RegisterMovieHandlerServer(s *grpc.Server, srv MovieHandlerServer) {
	s.RegisterService(&_MovieHandler_serviceDesc, srv)
}

func _MovieHandler_GetListMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).GetListMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieHandler/GetListMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).GetListMovie(ctx, req.(*ListMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MovieHandler",
	HandlerType: (*MovieHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListMovie",
			Handler:    _MovieHandler_GetListMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
