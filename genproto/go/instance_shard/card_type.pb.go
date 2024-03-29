// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: card_type.proto

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

type CardType int32

const (
	CardType_ATTACK   CardType = 0
	CardType_BUFF     CardType = 1
	CardType_DEBUFF   CardType = 2
	CardType_SUMMON   CardType = 3
	CardType_MOVEMENT CardType = 4
	CardType_HEAL     CardType = 5
	CardType_BLOCK    CardType = 6
)

// Enum value maps for CardType.
var (
	CardType_name = map[int32]string{
		0: "ATTACK",
		1: "BUFF",
		2: "DEBUFF",
		3: "SUMMON",
		4: "MOVEMENT",
		5: "HEAL",
		6: "BLOCK",
	}
	CardType_value = map[string]int32{
		"ATTACK":   0,
		"BUFF":     1,
		"DEBUFF":   2,
		"SUMMON":   3,
		"MOVEMENT": 4,
		"HEAL":     5,
		"BLOCK":    6,
	}
)

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}

func (x CardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardType) Descriptor() protoreflect.EnumDescriptor {
	return file_card_type_proto_enumTypes[0].Descriptor()
}

func (CardType) Type() protoreflect.EnumType {
	return &file_card_type_proto_enumTypes[0]
}

func (x CardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardType.Descriptor instead.
func (CardType) EnumDescriptor() ([]byte, []int) {
	return file_card_type_proto_rawDescGZIP(), []int{0}
}

var File_card_type_proto protoreflect.FileDescriptor

var file_card_type_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0x5b, 0x0a, 0x08, 0x43, 0x61,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x46, 0x46, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x42, 0x55, 0x46, 0x46, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4d, 0x4d,
	0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x4f, 0x56, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x06, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x76,
	0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_type_proto_rawDescOnce sync.Once
	file_card_type_proto_rawDescData = file_card_type_proto_rawDesc
)

func file_card_type_proto_rawDescGZIP() []byte {
	file_card_type_proto_rawDescOnce.Do(func() {
		file_card_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_type_proto_rawDescData)
	})
	return file_card_type_proto_rawDescData
}

var file_card_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_card_type_proto_goTypes = []interface{}{
	(CardType)(0), // 0: Deviant.CardType
}
var file_card_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_card_type_proto_init() }
func file_card_type_proto_init() {
	if File_card_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_card_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_type_proto_goTypes,
		DependencyIndexes: file_card_type_proto_depIdxs,
		EnumInfos:         file_card_type_proto_enumTypes,
	}.Build()
	File_card_type_proto = out.File
	file_card_type_proto_rawDesc = nil
	file_card_type_proto_goTypes = nil
	file_card_type_proto_depIdxs = nil
}
