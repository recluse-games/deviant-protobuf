// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: EntityTargetAction.proto

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

type EntityTargetAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tiles []*Tile `protobuf:"bytes,2,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *EntityTargetAction) Reset() {
	*x = EntityTargetAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityTargetAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityTargetAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityTargetAction) ProtoMessage() {}

func (x *EntityTargetAction) ProtoReflect() protoreflect.Message {
	mi := &file_EntityTargetAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityTargetAction.ProtoReflect.Descriptor instead.
func (*EntityTargetAction) Descriptor() ([]byte, []int) {
	return file_EntityTargetAction_proto_rawDescGZIP(), []int{0}
}

func (x *EntityTargetAction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntityTargetAction) GetTiles() []*Tile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

var File_EntityTargetAction_proto protoreflect.FileDescriptor

var file_EntityTargetAction_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x1a, 0x0a, 0x54, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x54,
	0x69, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b,
	0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityTargetAction_proto_rawDescOnce sync.Once
	file_EntityTargetAction_proto_rawDescData = file_EntityTargetAction_proto_rawDesc
)

func file_EntityTargetAction_proto_rawDescGZIP() []byte {
	file_EntityTargetAction_proto_rawDescOnce.Do(func() {
		file_EntityTargetAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityTargetAction_proto_rawDescData)
	})
	return file_EntityTargetAction_proto_rawDescData
}

var file_EntityTargetAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityTargetAction_proto_goTypes = []interface{}{
	(*EntityTargetAction)(nil), // 0: Deviant.EntityTargetAction
	(*Tile)(nil),               // 1: Deviant.Tile
}
var file_EntityTargetAction_proto_depIdxs = []int32{
	1, // 0: Deviant.EntityTargetAction.tiles:type_name -> Deviant.Tile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityTargetAction_proto_init() }
func file_EntityTargetAction_proto_init() {
	if File_EntityTargetAction_proto != nil {
		return
	}
	file_Tile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityTargetAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityTargetAction); i {
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
			RawDescriptor: file_EntityTargetAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityTargetAction_proto_goTypes,
		DependencyIndexes: file_EntityTargetAction_proto_depIdxs,
		MessageInfos:      file_EntityTargetAction_proto_msgTypes,
	}.Build()
	File_EntityTargetAction_proto = out.File
	file_EntityTargetAction_proto_rawDesc = nil
	file_EntityTargetAction_proto_goTypes = nil
	file_EntityTargetAction_proto_depIdxs = nil
}
