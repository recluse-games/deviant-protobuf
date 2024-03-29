// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: card_action.proto

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

type CardAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          CardType                 `protobuf:"varint,2,opt,name=type,proto3,enum=Deviant.CardType" json:"type,omitempty"`
	Status        CardActionStatusTypes    `protobuf:"varint,3,opt,name=status,proto3,enum=Deviant.CardActionStatusTypes" json:"status,omitempty"`
	TargetingType CardActionTargetingTypes `protobuf:"varint,4,opt,name=targeting_type,json=targetingType,proto3,enum=Deviant.CardActionTargetingTypes" json:"targeting_type,omitempty"`
	Up            []*CardAction            `protobuf:"bytes,5,rep,name=up,proto3" json:"up,omitempty"`
	Down          []*CardAction            `protobuf:"bytes,6,rep,name=down,proto3" json:"down,omitempty"`
	Left          []*CardAction            `protobuf:"bytes,7,rep,name=left,proto3" json:"left,omitempty"`
	Right         []*CardAction            `protobuf:"bytes,8,rep,name=right,proto3" json:"right,omitempty"`
	UpLeft        []*CardAction            `protobuf:"bytes,9,rep,name=up_left,json=upLeft,proto3" json:"up_left,omitempty"`
	UpRight       []*CardAction            `protobuf:"bytes,10,rep,name=up_right,json=upRight,proto3" json:"up_right,omitempty"`
	DownLeft      []*CardAction            `protobuf:"bytes,11,rep,name=down_left,json=downLeft,proto3" json:"down_left,omitempty"`
	DownRight     []*CardAction            `protobuf:"bytes,12,rep,name=down_right,json=downRight,proto3" json:"down_right,omitempty"`
	Origin        bool                     `protobuf:"varint,13,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *CardAction) Reset() {
	*x = CardAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardAction) ProtoMessage() {}

func (x *CardAction) ProtoReflect() protoreflect.Message {
	mi := &file_card_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardAction.ProtoReflect.Descriptor instead.
func (*CardAction) Descriptor() ([]byte, []int) {
	return file_card_action_proto_rawDescGZIP(), []int{0}
}

func (x *CardAction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CardAction) GetType() CardType {
	if x != nil {
		return x.Type
	}
	return CardType_ATTACK
}

func (x *CardAction) GetStatus() CardActionStatusTypes {
	if x != nil {
		return x.Status
	}
	return CardActionStatusTypes_EMPTY
}

func (x *CardAction) GetTargetingType() CardActionTargetingTypes {
	if x != nil {
		return x.TargetingType
	}
	return CardActionTargetingTypes_PATH
}

func (x *CardAction) GetUp() []*CardAction {
	if x != nil {
		return x.Up
	}
	return nil
}

func (x *CardAction) GetDown() []*CardAction {
	if x != nil {
		return x.Down
	}
	return nil
}

func (x *CardAction) GetLeft() []*CardAction {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *CardAction) GetRight() []*CardAction {
	if x != nil {
		return x.Right
	}
	return nil
}

func (x *CardAction) GetUpLeft() []*CardAction {
	if x != nil {
		return x.UpLeft
	}
	return nil
}

func (x *CardAction) GetUpRight() []*CardAction {
	if x != nil {
		return x.UpRight
	}
	return nil
}

func (x *CardAction) GetDownLeft() []*CardAction {
	if x != nil {
		return x.DownLeft
	}
	return nil
}

func (x *CardAction) GetDownRight() []*CardAction {
	if x != nil {
		return x.DownRight
	}
	return nil
}

func (x *CardAction) GetOrigin() bool {
	if x != nil {
		return x.Origin
	}
	return false
}

var File_card_action_proto protoreflect.FileDescriptor

var file_card_action_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x1e, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc3, 0x04, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48,
	0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x02, 0x75, 0x70, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x75, 0x70, 0x12, 0x27, 0x0a,
	0x04, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x75, 0x70,
	0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x75, 0x70, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x70, 0x5f, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x75, 0x70, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x64, 0x6f, 0x77, 0x6e,
	0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x6f,
	0x77, 0x6e, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_action_proto_rawDescOnce sync.Once
	file_card_action_proto_rawDescData = file_card_action_proto_rawDesc
)

func file_card_action_proto_rawDescGZIP() []byte {
	file_card_action_proto_rawDescOnce.Do(func() {
		file_card_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_action_proto_rawDescData)
	})
	return file_card_action_proto_rawDescData
}

var file_card_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_card_action_proto_goTypes = []interface{}{
	(*CardAction)(nil),            // 0: Deviant.CardAction
	(CardType)(0),                 // 1: Deviant.CardType
	(CardActionStatusTypes)(0),    // 2: Deviant.CardActionStatusTypes
	(CardActionTargetingTypes)(0), // 3: Deviant.CardActionTargetingTypes
}
var file_card_action_proto_depIdxs = []int32{
	1,  // 0: Deviant.CardAction.type:type_name -> Deviant.CardType
	2,  // 1: Deviant.CardAction.status:type_name -> Deviant.CardActionStatusTypes
	3,  // 2: Deviant.CardAction.targeting_type:type_name -> Deviant.CardActionTargetingTypes
	0,  // 3: Deviant.CardAction.up:type_name -> Deviant.CardAction
	0,  // 4: Deviant.CardAction.down:type_name -> Deviant.CardAction
	0,  // 5: Deviant.CardAction.left:type_name -> Deviant.CardAction
	0,  // 6: Deviant.CardAction.right:type_name -> Deviant.CardAction
	0,  // 7: Deviant.CardAction.up_left:type_name -> Deviant.CardAction
	0,  // 8: Deviant.CardAction.up_right:type_name -> Deviant.CardAction
	0,  // 9: Deviant.CardAction.down_left:type_name -> Deviant.CardAction
	0,  // 10: Deviant.CardAction.down_right:type_name -> Deviant.CardAction
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_card_action_proto_init() }
func file_card_action_proto_init() {
	if File_card_action_proto != nil {
		return
	}
	file_card_action_status_types_proto_init()
	file_card_action_targeting_types_proto_init()
	file_card_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_card_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardAction); i {
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
			RawDescriptor: file_card_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_action_proto_goTypes,
		DependencyIndexes: file_card_action_proto_depIdxs,
		MessageInfos:      file_card_action_proto_msgTypes,
	}.Build()
	File_card_action_proto = out.File
	file_card_action_proto_rawDesc = nil
	file_card_action_proto_goTypes = nil
	file_card_action_proto_depIdxs = nil
}
