// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.0
// source: http.proto

package conf

import (
	conf "github.com/go-lynx/lynx/app/conf"
	boot "github.com/go-lynx/lynx/boot"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Lynx    *boot.Lynx           `protobuf:"bytes,4,opt,name=lynx,proto3" json:"lynx,omitempty"`
	Tls     *conf.TlsService     `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

func (x *Http) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Http) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Http) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Http) GetLynx() *boot.Lynx {
	if x != nil {
		return x.Lynx
	}
	return nil
}

func (x *Http) GetTls() *conf.TlsService {
	if x != nil {
		return x.Tls
	}
	return nil
}

var File_http_proto protoreflect.FileDescriptor

var file_http_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x79,
	0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x79, 0x6e, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6c, 0x79, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x79, 0x6e, 0x78, 0x52, 0x04, 0x6c, 0x79, 0x6e, 0x78, 0x12, 0x33, 0x0a, 0x03, 0x74,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x79, 0x6e, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x03, 0x74, 0x6c, 0x73,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x6c, 0x79, 0x6e, 0x78, 0x2f, 0x6c, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_proto_rawDescOnce sync.Once
	file_http_proto_rawDescData = file_http_proto_rawDesc
)

func file_http_proto_rawDescGZIP() []byte {
	file_http_proto_rawDescOnce.Do(func() {
		file_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_proto_rawDescData)
	})
	return file_http_proto_rawDescData
}

var file_http_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_http_proto_goTypes = []interface{}{
	(*Http)(nil),                // 0: lynx.protobuf.http.http
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
	(*boot.Lynx)(nil),           // 2: lynx.protobuf.Lynx
	(*conf.TlsService)(nil),     // 3: lynx.protobuf.service.TlsService
}
var file_http_proto_depIdxs = []int32{
	1, // 0: lynx.protobuf.http.http.timeout:type_name -> google.protobuf.Duration
	2, // 1: lynx.protobuf.http.http.lynx:type_name -> lynx.protobuf.Lynx
	3, // 2: lynx.protobuf.http.http.tls:type_name -> lynx.protobuf.service.TlsService
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_http_proto_init() }
func file_http_proto_init() {
	if File_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
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
			RawDescriptor: file_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_proto_goTypes,
		DependencyIndexes: file_http_proto_depIdxs,
		MessageInfos:      file_http_proto_msgTypes,
	}.Build()
	File_http_proto = out.File
	file_http_proto_rawDesc = nil
	file_http_proto_goTypes = nil
	file_http_proto_depIdxs = nil
}
