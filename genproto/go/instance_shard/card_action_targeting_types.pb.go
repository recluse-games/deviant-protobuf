// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: card_action_targeting_types.proto

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

type CardActionTargetingTypes int32

const (
	CardActionTargetingTypes_PATH   CardActionTargetingTypes = 0
	CardActionTargetingTypes_GROUND CardActionTargetingTypes = 1
)

// Enum value maps for CardActionTargetingTypes.
var (
	CardActionTargetingTypes_name = map[int32]string{
		0: "PATH",
		1: "GROUND",
	}
	CardActionTargetingTypes_value = map[string]int32{
		"PATH":   0,
		"GROUND": 1,
	}
)

func (x CardActionTargetingTypes) Enum() *CardActionTargetingTypes {
	p := new(CardActionTargetingTypes)
	*p = x
	return p
}

func (x CardActionTargetingTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardActionTargetingTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_card_action_targeting_types_proto_enumTypes[0].Descriptor()
}

func (CardActionTargetingTypes) Type() protoreflect.EnumType {
	return &file_card_action_targeting_types_proto_enumTypes[0]
}

func (x CardActionTargetingTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardActionTargetingTypes.Descriptor instead.
func (CardActionTargetingTypes) EnumDescriptor() ([]byte, []int) {
	return file_card_action_targeting_types_proto_rawDescGZIP(), []int{0}
}

var File_card_action_targeting_types_proto protoreflect.FileDescriptor

var file_card_action_targeting_types_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x30, 0x0a, 0x18,
	0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x54, 0x48,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_card_action_targeting_types_proto_rawDescOnce sync.Once
	file_card_action_targeting_types_proto_rawDescData = file_card_action_targeting_types_proto_rawDesc
)

func file_card_action_targeting_types_proto_rawDescGZIP() []byte {
	file_card_action_targeting_types_proto_rawDescOnce.Do(func() {
		file_card_action_targeting_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_action_targeting_types_proto_rawDescData)
	})
	return file_card_action_targeting_types_proto_rawDescData
}

var file_card_action_targeting_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_card_action_targeting_types_proto_goTypes = []interface{}{
	(CardActionTargetingTypes)(0), // 0: Deviant.CardActionTargetingTypes
}
var file_card_action_targeting_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_card_action_targeting_types_proto_init() }
func file_card_action_targeting_types_proto_init() {
	if File_card_action_targeting_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_card_action_targeting_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_action_targeting_types_proto_goTypes,
		DependencyIndexes: file_card_action_targeting_types_proto_depIdxs,
		EnumInfos:         file_card_action_targeting_types_proto_enumTypes,
	}.Build()
	File_card_action_targeting_types_proto = out.File
	file_card_action_targeting_types_proto_rawDesc = nil
	file_card_action_targeting_types_proto_goTypes = nil
	file_card_action_targeting_types_proto_depIdxs = nil
}
