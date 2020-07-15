// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: board.proto

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

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tiles        *Tiles    `protobuf:"bytes,1,opt,name=tiles,proto3" json:"tiles,omitempty"`
	Entities     *Entities `protobuf:"bytes,2,opt,name=entities,proto3" json:"entities,omitempty"`
	OverlayTiles []*Tile   `protobuf:"bytes,3,rep,name=overlayTiles,proto3" json:"overlayTiles,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetTiles() *Tiles {
	if x != nil {
		return x.Tiles
	}
	return nil
}

func (x *Board) GetEntities() *Entities {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *Board) GetOverlayTiles() []*Tile {
	if x != nil {
		return x.OverlayTiles
	}
	return nil
}

var File_board_proto protoreflect.FileDescriptor

var file_board_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44,
	0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x0b, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x74, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8f, 0x01, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61,
	0x6e, 0x74, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x31,
	0x0a, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x54,
	0x69, 0x6c, 0x65, 0x52, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6c, 0x65,
	0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_proto_rawDescOnce sync.Once
	file_board_proto_rawDescData = file_board_proto_rawDesc
)

func file_board_proto_rawDescGZIP() []byte {
	file_board_proto_rawDescOnce.Do(func() {
		file_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_proto_rawDescData)
	})
	return file_board_proto_rawDescData
}

var file_board_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_board_proto_goTypes = []interface{}{
	(*Board)(nil),    // 0: Deviant.Board
	(*Tiles)(nil),    // 1: Deviant.Tiles
	(*Entities)(nil), // 2: Deviant.Entities
	(*Tile)(nil),     // 3: Deviant.Tile
}
var file_board_proto_depIdxs = []int32{
	1, // 0: Deviant.Board.tiles:type_name -> Deviant.Tiles
	2, // 1: Deviant.Board.entities:type_name -> Deviant.Entities
	3, // 2: Deviant.Board.overlayTiles:type_name -> Deviant.Tile
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_board_proto_init() }
func file_board_proto_init() {
	if File_board_proto != nil {
		return
	}
	file_tiles_proto_init()
	file_entities_proto_init()
	file_tile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
			RawDescriptor: file_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_board_proto_goTypes,
		DependencyIndexes: file_board_proto_depIdxs,
		MessageInfos:      file_board_proto_msgTypes,
	}.Build()
	File_board_proto = out.File
	file_board_proto_rawDesc = nil
	file_board_proto_goTypes = nil
	file_board_proto_depIdxs = nil
}