// protos/did.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: protos/did.proto

package protos

import (
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

type Did struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Did) Reset() {
	*x = Did{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_did_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Did) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Did) ProtoMessage() {}

func (x *Did) ProtoReflect() protoreflect.Message {
	mi := &file_protos_did_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Did.ProtoReflect.Descriptor instead.
func (*Did) Descriptor() ([]byte, []int) {
	return file_protos_did_proto_rawDescGZIP(), []int{0}
}

func (x *Did) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Did) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var File_protos_did_proto protoreflect.FileDescriptor

var file_protos_did_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x64, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x03, 0x44, 0x69, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x32, 0x2b, 0x0a, 0x09, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x44, 0x69, 0x64, 0x12, 0x08, 0x2e, 0x64, 0x69, 0x64, 0x2e,
	0x44, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x44, 0x69, 0x64, 0x22, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x73, 0x73, 0x69, 0x6b, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_did_proto_rawDescOnce sync.Once
	file_protos_did_proto_rawDescData = file_protos_did_proto_rawDesc
)

func file_protos_did_proto_rawDescGZIP() []byte {
	file_protos_did_proto_rawDescOnce.Do(func() {
		file_protos_did_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_did_proto_rawDescData)
	})
	return file_protos_did_proto_rawDescData
}

var file_protos_did_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_did_proto_goTypes = []interface{}{
	(*Did)(nil), // 0: did.Did
}
var file_protos_did_proto_depIdxs = []int32{
	0, // 0: did.SimpleDid.GetDid:input_type -> did.Did
	0, // 1: did.SimpleDid.GetDid:output_type -> did.Did
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_did_proto_init() }
func file_protos_did_proto_init() {
	if File_protos_did_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_did_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Did); i {
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
			RawDescriptor: file_protos_did_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_did_proto_goTypes,
		DependencyIndexes: file_protos_did_proto_depIdxs,
		MessageInfos:      file_protos_did_proto_msgTypes,
	}.Build()
	File_protos_did_proto = out.File
	file_protos_did_proto_rawDesc = nil
	file_protos_did_proto_goTypes = nil
	file_protos_did_proto_depIdxs = nil
}
