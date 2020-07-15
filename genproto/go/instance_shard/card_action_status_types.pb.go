// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: card_action_status_types.proto

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

type CardActionStatusTypes int32

const (
	CardActionStatusTypes_EMPTY     CardActionStatusTypes = 0
	CardActionStatusTypes_UNBLOCKED CardActionStatusTypes = 1
	CardActionStatusTypes_BLOCKED   CardActionStatusTypes = 2
)

// Enum value maps for CardActionStatusTypes.
var (
	CardActionStatusTypes_name = map[int32]string{
		0: "EMPTY",
		1: "UNBLOCKED",
		2: "BLOCKED",
	}
	CardActionStatusTypes_value = map[string]int32{
		"EMPTY":     0,
		"UNBLOCKED": 1,
		"BLOCKED":   2,
	}
)

func (x CardActionStatusTypes) Enum() *CardActionStatusTypes {
	p := new(CardActionStatusTypes)
	*p = x
	return p
}

func (x CardActionStatusTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardActionStatusTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_card_action_status_types_proto_enumTypes[0].Descriptor()
}

func (CardActionStatusTypes) Type() protoreflect.EnumType {
	return &file_card_action_status_types_proto_enumTypes[0]
}

func (x CardActionStatusTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardActionStatusTypes.Descriptor instead.
func (CardActionStatusTypes) EnumDescriptor() ([]byte, []int) {
	return file_card_action_status_types_proto_rawDescGZIP(), []int{0}
}

var File_card_action_status_types_proto protoreflect.FileDescriptor

var file_card_action_status_types_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x3e, 0x0a, 0x15, 0x43, 0x61, 0x72,
	0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64,
	0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_action_status_types_proto_rawDescOnce sync.Once
	file_card_action_status_types_proto_rawDescData = file_card_action_status_types_proto_rawDesc
)

func file_card_action_status_types_proto_rawDescGZIP() []byte {
	file_card_action_status_types_proto_rawDescOnce.Do(func() {
		file_card_action_status_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_action_status_types_proto_rawDescData)
	})
	return file_card_action_status_types_proto_rawDescData
}

var file_card_action_status_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_card_action_status_types_proto_goTypes = []interface{}{
	(CardActionStatusTypes)(0), // 0: Deviant.CardActionStatusTypes
}
var file_card_action_status_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_card_action_status_types_proto_init() }
func file_card_action_status_types_proto_init() {
	if File_card_action_status_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_card_action_status_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_action_status_types_proto_goTypes,
		DependencyIndexes: file_card_action_status_types_proto_depIdxs,
		EnumInfos:         file_card_action_status_types_proto_enumTypes,
	}.Build()
	File_card_action_status_types_proto = out.File
	file_card_action_status_types_proto_rawDesc = nil
	file_card_action_status_types_proto_goTypes = nil
	file_card_action_status_types_proto_depIdxs = nil
}
