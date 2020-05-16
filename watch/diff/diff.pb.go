// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: watch/diff/diff.proto

package diff

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message
type DiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DiffRequest) Reset() {
	*x = DiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_diff_diff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffRequest) ProtoMessage() {}

func (x *DiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watch_diff_diff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffRequest.ProtoReflect.Descriptor instead.
func (*DiffRequest) Descriptor() ([]byte, []int) {
	return file_watch_diff_diff_proto_rawDescGZIP(), []int{0}
}

func (x *DiffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the diff
type DiffCommits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Commits   string `protobuf:"bytes,2,opt,name=commits,proto3" json:"commits,omitempty"`
	Salt      []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	Encrypted bool   `protobuf:"varint,4,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *DiffCommits) Reset() {
	*x = DiffCommits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_diff_diff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffCommits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffCommits) ProtoMessage() {}

func (x *DiffCommits) ProtoReflect() protoreflect.Message {
	mi := &file_watch_diff_diff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffCommits.ProtoReflect.Descriptor instead.
func (*DiffCommits) Descriptor() ([]byte, []int) {
	return file_watch_diff_diff_proto_rawDescGZIP(), []int{1}
}

func (x *DiffCommits) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiffCommits) GetCommits() string {
	if x != nil {
		return x.Commits
	}
	return ""
}

func (x *DiffCommits) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *DiffCommits) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

type DiffConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiffConfirm) Reset() {
	*x = DiffConfirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_diff_diff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffConfirm) ProtoMessage() {}

func (x *DiffConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_watch_diff_diff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffConfirm.ProtoReflect.Descriptor instead.
func (*DiffConfirm) Descriptor() ([]byte, []int) {
	return file_watch_diff_diff_proto_rawDescGZIP(), []int{2}
}

var File_watch_diff_diff_proto protoreflect.FileDescriptor

var file_watch_diff_diff_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x69, 0x66, 0x66, 0x2f, 0x64, 0x69, 0x66,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x69, 0x66, 0x66, 0x22, 0x21, 0x0a,
	0x0b, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6d, 0x0a, 0x0b, 0x44, 0x69, 0x66, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x22,
	0x0d, 0x0a, 0x0b, 0x44, 0x69, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x32, 0x70,
	0x0a, 0x04, 0x44, 0x69, 0x66, 0x66, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69,
	0x66, 0x66, 0x12, 0x11, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x66,
	0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x66, 0x66, 0x12, 0x11, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e,
	0x44, 0x69, 0x66, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x11, 0x2e, 0x64, 0x69,
	0x66, 0x66, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x69, 0x66, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watch_diff_diff_proto_rawDescOnce sync.Once
	file_watch_diff_diff_proto_rawDescData = file_watch_diff_diff_proto_rawDesc
)

func file_watch_diff_diff_proto_rawDescGZIP() []byte {
	file_watch_diff_diff_proto_rawDescOnce.Do(func() {
		file_watch_diff_diff_proto_rawDescData = protoimpl.X.CompressGZIP(file_watch_diff_diff_proto_rawDescData)
	})
	return file_watch_diff_diff_proto_rawDescData
}

var file_watch_diff_diff_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_watch_diff_diff_proto_goTypes = []interface{}{
	(*DiffRequest)(nil), // 0: diff.DiffRequest
	(*DiffCommits)(nil), // 1: diff.DiffCommits
	(*DiffConfirm)(nil), // 2: diff.DiffConfirm
}
var file_watch_diff_diff_proto_depIdxs = []int32{
	0, // 0: diff.Diff.SendDiff:input_type -> diff.DiffRequest
	1, // 1: diff.Diff.UploadDiff:input_type -> diff.DiffCommits
	1, // 2: diff.Diff.SendDiff:output_type -> diff.DiffCommits
	2, // 3: diff.Diff.UploadDiff:output_type -> diff.DiffConfirm
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_watch_diff_diff_proto_init() }
func file_watch_diff_diff_proto_init() {
	if File_watch_diff_diff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_watch_diff_diff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffRequest); i {
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
		file_watch_diff_diff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffCommits); i {
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
		file_watch_diff_diff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffConfirm); i {
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
			RawDescriptor: file_watch_diff_diff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watch_diff_diff_proto_goTypes,
		DependencyIndexes: file_watch_diff_diff_proto_depIdxs,
		MessageInfos:      file_watch_diff_diff_proto_msgTypes,
	}.Build()
	File_watch_diff_diff_proto = out.File
	file_watch_diff_diff_proto_rawDesc = nil
	file_watch_diff_diff_proto_goTypes = nil
	file_watch_diff_diff_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiffClient is the client API for Diff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiffClient interface {
	// Sends a diff patch to be applied
	SendDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffCommits, error)
	// Upload diff patches
	UploadDiff(ctx context.Context, in *DiffCommits, opts ...grpc.CallOption) (*DiffConfirm, error)
}

type diffClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffClient(cc grpc.ClientConnInterface) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) SendDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffCommits, error) {
	out := new(DiffCommits)
	err := c.cc.Invoke(ctx, "/diff.Diff/SendDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffClient) UploadDiff(ctx context.Context, in *DiffCommits, opts ...grpc.CallOption) (*DiffConfirm, error) {
	out := new(DiffConfirm)
	err := c.cc.Invoke(ctx, "/diff.Diff/UploadDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServer is the server API for Diff service.
type DiffServer interface {
	// Sends a diff patch to be applied
	SendDiff(context.Context, *DiffRequest) (*DiffCommits, error)
	// Upload diff patches
	UploadDiff(context.Context, *DiffCommits) (*DiffConfirm, error)
}

// UnimplementedDiffServer can be embedded to have forward compatible implementations.
type UnimplementedDiffServer struct {
}

func (*UnimplementedDiffServer) SendDiff(context.Context, *DiffRequest) (*DiffCommits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDiff not implemented")
}
func (*UnimplementedDiffServer) UploadDiff(context.Context, *DiffCommits) (*DiffConfirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDiff not implemented")
}

func RegisterDiffServer(s *grpc.Server, srv DiffServer) {
	s.RegisterService(&_Diff_serviceDesc, srv)
}

func _Diff_SendDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).SendDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diff.Diff/SendDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).SendDiff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diff_UploadDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffCommits)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).UploadDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diff.Diff/UploadDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).UploadDiff(ctx, req.(*DiffCommits))
	}
	return interceptor(ctx, in, info, handler)
}

var _Diff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "diff.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDiff",
			Handler:    _Diff_SendDiff_Handler,
		},
		{
			MethodName: "UploadDiff",
			Handler:    _Diff_UploadDiff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watch/diff/diff.proto",
}
