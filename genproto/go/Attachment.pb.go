// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: Attachment.proto

package deviant

import (
	proto "github.com/golang/protobuf/proto"
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

type Attachment int32

const (
	Attachment_ID Attachment = 0
)

// Enum value maps for Attachment.
var (
	Attachment_name = map[int32]string{
		0: "ID",
	}
	Attachment_value = map[string]int32{
		"ID": 0,
	}
)

func (x Attachment) Enum() *Attachment {
	p := new(Attachment)
	*p = x
	return p
}

func (x Attachment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attachment) Descriptor() protoreflect.EnumDescriptor {
	return file_Attachment_proto_enumTypes[0].Descriptor()
}

func (Attachment) Type() protoreflect.EnumType {
	return &file_Attachment_proto_enumTypes[0]
}

func (x Attachment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attachment.Descriptor instead.
func (Attachment) EnumDescriptor() ([]byte, []int) {
	return file_Attachment_proto_rawDescGZIP(), []int{0}
}

var File_Attachment_proto protoreflect.FileDescriptor

var file_Attachment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x14, 0x0a, 0x0a, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Attachment_proto_rawDescOnce sync.Once
	file_Attachment_proto_rawDescData = file_Attachment_proto_rawDesc
)

func file_Attachment_proto_rawDescGZIP() []byte {
	file_Attachment_proto_rawDescOnce.Do(func() {
		file_Attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_Attachment_proto_rawDescData)
	})
	return file_Attachment_proto_rawDescData
}

var file_Attachment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Attachment_proto_goTypes = []interface{}{
	(Attachment)(0), // 0: Deviant.Attachment
}
var file_Attachment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Attachment_proto_init() }
func file_Attachment_proto_init() {
	if File_Attachment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Attachment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Attachment_proto_goTypes,
		DependencyIndexes: file_Attachment_proto_depIdxs,
		EnumInfos:         file_Attachment_proto_enumTypes,
	}.Build()
	File_Attachment_proto = out.File
	file_Attachment_proto_rawDesc = nil
	file_Attachment_proto_goTypes = nil
	file_Attachment_proto_depIdxs = nil
}
