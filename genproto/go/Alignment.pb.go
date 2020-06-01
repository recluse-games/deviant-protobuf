// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: Alignment.proto

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

type Alignment int32

const (
	Alignment_FRIENDLY   Alignment = 0
	Alignment_UNFRIENDLY Alignment = 1
	Alignment_NEUTRAL    Alignment = 2
)

// Enum value maps for Alignment.
var (
	Alignment_name = map[int32]string{
		0: "FRIENDLY",
		1: "UNFRIENDLY",
		2: "NEUTRAL",
	}
	Alignment_value = map[string]int32{
		"FRIENDLY":   0,
		"UNFRIENDLY": 1,
		"NEUTRAL":    2,
	}
)

func (x Alignment) Enum() *Alignment {
	p := new(Alignment)
	*p = x
	return p
}

func (x Alignment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Alignment) Descriptor() protoreflect.EnumDescriptor {
	return file_Alignment_proto_enumTypes[0].Descriptor()
}

func (Alignment) Type() protoreflect.EnumType {
	return &file_Alignment_proto_enumTypes[0]
}

func (x Alignment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Alignment.Descriptor instead.
func (Alignment) EnumDescriptor() ([]byte, []int) {
	return file_Alignment_proto_rawDescGZIP(), []int{0}
}

var File_Alignment_proto protoreflect.FileDescriptor

var file_Alignment_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x36, 0x0a, 0x09, 0x41, 0x6c,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x55, 0x54, 0x52, 0x41, 0x4c,
	0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Alignment_proto_rawDescOnce sync.Once
	file_Alignment_proto_rawDescData = file_Alignment_proto_rawDesc
)

func file_Alignment_proto_rawDescGZIP() []byte {
	file_Alignment_proto_rawDescOnce.Do(func() {
		file_Alignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_Alignment_proto_rawDescData)
	})
	return file_Alignment_proto_rawDescData
}

var file_Alignment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Alignment_proto_goTypes = []interface{}{
	(Alignment)(0), // 0: Deviant.Alignment
}
var file_Alignment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Alignment_proto_init() }
func file_Alignment_proto_init() {
	if File_Alignment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Alignment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Alignment_proto_goTypes,
		DependencyIndexes: file_Alignment_proto_depIdxs,
		EnumInfos:         file_Alignment_proto_enumTypes,
	}.Build()
	File_Alignment_proto = out.File
	file_Alignment_proto_rawDesc = nil
	file_Alignment_proto_goTypes = nil
	file_Alignment_proto_depIdxs = nil
}
