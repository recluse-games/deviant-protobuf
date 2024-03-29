// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: entity_move_action.proto

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

type EntityMoveAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartXPosition int32 `protobuf:"varint,1,opt,name=startXPosition,proto3" json:"startXPosition,omitempty"`
	StartYPosition int32 `protobuf:"varint,2,opt,name=startYPosition,proto3" json:"startYPosition,omitempty"`
	FinalXPosition int32 `protobuf:"varint,3,opt,name=finalXPosition,proto3" json:"finalXPosition,omitempty"`
	FinalYPosition int32 `protobuf:"varint,4,opt,name=finalYPosition,proto3" json:"finalYPosition,omitempty"`
}

func (x *EntityMoveAction) Reset() {
	*x = EntityMoveAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_move_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityMoveAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityMoveAction) ProtoMessage() {}

func (x *EntityMoveAction) ProtoReflect() protoreflect.Message {
	mi := &file_entity_move_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityMoveAction.ProtoReflect.Descriptor instead.
func (*EntityMoveAction) Descriptor() ([]byte, []int) {
	return file_entity_move_action_proto_rawDescGZIP(), []int{0}
}

func (x *EntityMoveAction) GetStartXPosition() int32 {
	if x != nil {
		return x.StartXPosition
	}
	return 0
}

func (x *EntityMoveAction) GetStartYPosition() int32 {
	if x != nil {
		return x.StartYPosition
	}
	return 0
}

func (x *EntityMoveAction) GetFinalXPosition() int32 {
	if x != nil {
		return x.FinalXPosition
	}
	return 0
}

func (x *EntityMoveAction) GetFinalYPosition() int32 {
	if x != nil {
		return x.FinalYPosition
	}
	return 0
}

var File_entity_move_action_proto protoreflect.FileDescriptor

var file_entity_move_action_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f,
	0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x58, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x58, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x58, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x58, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x59, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x59,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_move_action_proto_rawDescOnce sync.Once
	file_entity_move_action_proto_rawDescData = file_entity_move_action_proto_rawDesc
)

func file_entity_move_action_proto_rawDescGZIP() []byte {
	file_entity_move_action_proto_rawDescOnce.Do(func() {
		file_entity_move_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_move_action_proto_rawDescData)
	})
	return file_entity_move_action_proto_rawDescData
}

var file_entity_move_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_move_action_proto_goTypes = []interface{}{
	(*EntityMoveAction)(nil), // 0: Deviant.EntityMoveAction
}
var file_entity_move_action_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_move_action_proto_init() }
func file_entity_move_action_proto_init() {
	if File_entity_move_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_move_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityMoveAction); i {
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
			RawDescriptor: file_entity_move_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_move_action_proto_goTypes,
		DependencyIndexes: file_entity_move_action_proto_depIdxs,
		MessageInfos:      file_entity_move_action_proto_msgTypes,
	}.Build()
	File_entity_move_action_proto = out.File
	file_entity_move_action_proto_rawDesc = nil
	file_entity_move_action_proto_goTypes = nil
	file_entity_move_action_proto_depIdxs = nil
}
