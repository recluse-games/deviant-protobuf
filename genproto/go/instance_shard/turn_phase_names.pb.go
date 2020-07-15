// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: turn_phase_names.proto

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

type TurnPhaseNames int32

const (
	TurnPhaseNames_PHASE_POINT   TurnPhaseNames = 0
	TurnPhaseNames_PHASE_EFFECT  TurnPhaseNames = 1
	TurnPhaseNames_PHASE_DRAW    TurnPhaseNames = 2
	TurnPhaseNames_PHASE_ACTION  TurnPhaseNames = 3
	TurnPhaseNames_PHASE_DISCARD TurnPhaseNames = 4
	TurnPhaseNames_PHASE_END     TurnPhaseNames = 5
)

// Enum value maps for TurnPhaseNames.
var (
	TurnPhaseNames_name = map[int32]string{
		0: "PHASE_POINT",
		1: "PHASE_EFFECT",
		2: "PHASE_DRAW",
		3: "PHASE_ACTION",
		4: "PHASE_DISCARD",
		5: "PHASE_END",
	}
	TurnPhaseNames_value = map[string]int32{
		"PHASE_POINT":   0,
		"PHASE_EFFECT":  1,
		"PHASE_DRAW":    2,
		"PHASE_ACTION":  3,
		"PHASE_DISCARD": 4,
		"PHASE_END":     5,
	}
)

func (x TurnPhaseNames) Enum() *TurnPhaseNames {
	p := new(TurnPhaseNames)
	*p = x
	return p
}

func (x TurnPhaseNames) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TurnPhaseNames) Descriptor() protoreflect.EnumDescriptor {
	return file_turn_phase_names_proto_enumTypes[0].Descriptor()
}

func (TurnPhaseNames) Type() protoreflect.EnumType {
	return &file_turn_phase_names_proto_enumTypes[0]
}

func (x TurnPhaseNames) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TurnPhaseNames.Descriptor instead.
func (TurnPhaseNames) EnumDescriptor() ([]byte, []int) {
	return file_turn_phase_names_proto_rawDescGZIP(), []int{0}
}

var File_turn_phase_names_proto protoreflect.FileDescriptor

var file_turn_phase_names_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e,
	0x74, 0x2a, 0x77, 0x0a, 0x0e, 0x54, 0x75, 0x72, 0x6e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x46,
	0x46, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f,
	0x44, 0x52, 0x41, 0x57, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x41, 0x53,
	0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x48, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x05, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b,
	0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_turn_phase_names_proto_rawDescOnce sync.Once
	file_turn_phase_names_proto_rawDescData = file_turn_phase_names_proto_rawDesc
)

func file_turn_phase_names_proto_rawDescGZIP() []byte {
	file_turn_phase_names_proto_rawDescOnce.Do(func() {
		file_turn_phase_names_proto_rawDescData = protoimpl.X.CompressGZIP(file_turn_phase_names_proto_rawDescData)
	})
	return file_turn_phase_names_proto_rawDescData
}

var file_turn_phase_names_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_turn_phase_names_proto_goTypes = []interface{}{
	(TurnPhaseNames)(0), // 0: Deviant.TurnPhaseNames
}
var file_turn_phase_names_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_turn_phase_names_proto_init() }
func file_turn_phase_names_proto_init() {
	if File_turn_phase_names_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_turn_phase_names_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_turn_phase_names_proto_goTypes,
		DependencyIndexes: file_turn_phase_names_proto_depIdxs,
		EnumInfos:         file_turn_phase_names_proto_enumTypes,
	}.Build()
	File_turn_phase_names_proto = out.File
	file_turn_phase_names_proto_rawDesc = nil
	file_turn_phase_names_proto_goTypes = nil
	file_turn_phase_names_proto_depIdxs = nil
}
