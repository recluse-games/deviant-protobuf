// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: EntityActionNames.proto

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

type EntityActionNames int32

const (
	EntityActionNames_PLAY    EntityActionNames = 0
	EntityActionNames_DRAW    EntityActionNames = 1
	EntityActionNames_DISCARD EntityActionNames = 2
	EntityActionNames_NOTHING EntityActionNames = 3
)

// Enum value maps for EntityActionNames.
var (
	EntityActionNames_name = map[int32]string{
		0: "PLAY",
		1: "DRAW",
		2: "DISCARD",
		3: "NOTHING",
	}
	EntityActionNames_value = map[string]int32{
		"PLAY":    0,
		"DRAW":    1,
		"DISCARD": 2,
		"NOTHING": 3,
	}
)

func (x EntityActionNames) Enum() *EntityActionNames {
	p := new(EntityActionNames)
	*p = x
	return p
}

func (x EntityActionNames) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityActionNames) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityActionNames_proto_enumTypes[0].Descriptor()
}

func (EntityActionNames) Type() protoreflect.EnumType {
	return &file_EntityActionNames_proto_enumTypes[0]
}

func (x EntityActionNames) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityActionNames.Descriptor instead.
func (EntityActionNames) EnumDescriptor() ([]byte, []int) {
	return file_EntityActionNames_proto_rawDescGZIP(), []int{0}
}

var File_EntityActionNames_proto protoreflect.FileDescriptor

var file_EntityActionNames_proto_rawDesc = []byte{
	0x0a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61,
	0x6e, 0x74, 0x2a, 0x41, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x41, 0x59, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x48,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityActionNames_proto_rawDescOnce sync.Once
	file_EntityActionNames_proto_rawDescData = file_EntityActionNames_proto_rawDesc
)

func file_EntityActionNames_proto_rawDescGZIP() []byte {
	file_EntityActionNames_proto_rawDescOnce.Do(func() {
		file_EntityActionNames_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityActionNames_proto_rawDescData)
	})
	return file_EntityActionNames_proto_rawDescData
}

var file_EntityActionNames_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityActionNames_proto_goTypes = []interface{}{
	(EntityActionNames)(0), // 0: Deviant.EntityActionNames
}
var file_EntityActionNames_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EntityActionNames_proto_init() }
func file_EntityActionNames_proto_init() {
	if File_EntityActionNames_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EntityActionNames_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityActionNames_proto_goTypes,
		DependencyIndexes: file_EntityActionNames_proto_depIdxs,
		EnumInfos:         file_EntityActionNames_proto_enumTypes,
	}.Build()
	File_EntityActionNames_proto = out.File
	file_EntityActionNames_proto_rawDesc = nil
	file_EntityActionNames_proto_goTypes = nil
	file_EntityActionNames_proto_depIdxs = nil
}
