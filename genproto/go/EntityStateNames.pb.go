// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: EntityStateNames.proto

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

type EntityStateNames int32

const (
	EntityStateNames_IDLE       EntityStateNames = 0
	EntityStateNames_MOVING     EntityStateNames = 1
	EntityStateNames_ATTACKING  EntityStateNames = 2
	EntityStateNames_CASTING    EntityStateNames = 3
	EntityStateNames_RECOVERING EntityStateNames = 4
	EntityStateNames_RECOILING  EntityStateNames = 5
	EntityStateNames_DYING      EntityStateNames = 6
	EntityStateNames_DEAD       EntityStateNames = 7
)

// Enum value maps for EntityStateNames.
var (
	EntityStateNames_name = map[int32]string{
		0: "IDLE",
		1: "MOVING",
		2: "ATTACKING",
		3: "CASTING",
		4: "RECOVERING",
		5: "RECOILING",
		6: "DYING",
		7: "DEAD",
	}
	EntityStateNames_value = map[string]int32{
		"IDLE":       0,
		"MOVING":     1,
		"ATTACKING":  2,
		"CASTING":    3,
		"RECOVERING": 4,
		"RECOILING":  5,
		"DYING":      6,
		"DEAD":       7,
	}
)

func (x EntityStateNames) Enum() *EntityStateNames {
	p := new(EntityStateNames)
	*p = x
	return p
}

func (x EntityStateNames) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityStateNames) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityStateNames_proto_enumTypes[0].Descriptor()
}

func (EntityStateNames) Type() protoreflect.EnumType {
	return &file_EntityStateNames_proto_enumTypes[0]
}

func (x EntityStateNames) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityStateNames.Descriptor instead.
func (EntityStateNames) EnumDescriptor() ([]byte, []int) {
	return file_EntityStateNames_proto_rawDescGZIP(), []int{0}
}

var File_EntityStateNames_proto protoreflect.FileDescriptor

var file_EntityStateNames_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e,
	0x74, 0x2a, 0x78, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41,
	0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x43, 0x4f, 0x49,
	0x4c, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x59, 0x49, 0x4e, 0x47, 0x10,
	0x06, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x41, 0x44, 0x10, 0x07, 0x42, 0x09, 0x5a, 0x07, 0x64,
	0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityStateNames_proto_rawDescOnce sync.Once
	file_EntityStateNames_proto_rawDescData = file_EntityStateNames_proto_rawDesc
)

func file_EntityStateNames_proto_rawDescGZIP() []byte {
	file_EntityStateNames_proto_rawDescOnce.Do(func() {
		file_EntityStateNames_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityStateNames_proto_rawDescData)
	})
	return file_EntityStateNames_proto_rawDescData
}

var file_EntityStateNames_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityStateNames_proto_goTypes = []interface{}{
	(EntityStateNames)(0), // 0: Deviant.EntityStateNames
}
var file_EntityStateNames_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EntityStateNames_proto_init() }
func file_EntityStateNames_proto_init() {
	if File_EntityStateNames_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EntityStateNames_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityStateNames_proto_goTypes,
		DependencyIndexes: file_EntityStateNames_proto_depIdxs,
		EnumInfos:         file_EntityStateNames_proto_enumTypes,
	}.Build()
	File_EntityStateNames_proto = out.File
	file_EntityStateNames_proto_rawDesc = nil
	file_EntityStateNames_proto_goTypes = nil
	file_EntityStateNames_proto_depIdxs = nil
}
