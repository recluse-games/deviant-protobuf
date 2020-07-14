// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: Conditions.proto

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

type Conditions int32

const (
	Conditions_SLOW Conditions = 0
)

// Enum value maps for Conditions.
var (
	Conditions_name = map[int32]string{
		0: "SLOW",
	}
	Conditions_value = map[string]int32{
		"SLOW": 0,
	}
)

func (x Conditions) Enum() *Conditions {
	p := new(Conditions)
	*p = x
	return p
}

func (x Conditions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Conditions) Descriptor() protoreflect.EnumDescriptor {
	return file_Conditions_proto_enumTypes[0].Descriptor()
}

func (Conditions) Type() protoreflect.EnumType {
	return &file_Conditions_proto_enumTypes[0]
}

func (x Conditions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Conditions.Descriptor instead.
func (Conditions) EnumDescriptor() ([]byte, []int) {
	return file_Conditions_proto_rawDescGZIP(), []int{0}
}

var File_Conditions_proto protoreflect.FileDescriptor

var file_Conditions_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x16, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c, 0x4f,
	0x57, 0x10, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Conditions_proto_rawDescOnce sync.Once
	file_Conditions_proto_rawDescData = file_Conditions_proto_rawDesc
)

func file_Conditions_proto_rawDescGZIP() []byte {
	file_Conditions_proto_rawDescOnce.Do(func() {
		file_Conditions_proto_rawDescData = protoimpl.X.CompressGZIP(file_Conditions_proto_rawDescData)
	})
	return file_Conditions_proto_rawDescData
}

var file_Conditions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Conditions_proto_goTypes = []interface{}{
	(Conditions)(0), // 0: Deviant.Conditions
}
var file_Conditions_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Conditions_proto_init() }
func file_Conditions_proto_init() {
	if File_Conditions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Conditions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Conditions_proto_goTypes,
		DependencyIndexes: file_Conditions_proto_depIdxs,
		EnumInfos:         file_Conditions_proto_enumTypes,
	}.Build()
	File_Conditions_proto = out.File
	file_Conditions_proto_rawDesc = nil
	file_Conditions_proto_goTypes = nil
	file_Conditions_proto_depIdxs = nil
}
