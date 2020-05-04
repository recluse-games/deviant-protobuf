// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: Card.proto

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceId  string      `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	Cost        int32       `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	Title       string      `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Flavor      string      `protobuf:"bytes,5,opt,name=flavor,proto3" json:"flavor,omitempty"`
	Description string      `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Type        CardType    `protobuf:"varint,7,opt,name=type,proto3,enum=Deviant.CardType" json:"type,omitempty"`
	Action      *CardAction `protobuf:"bytes,8,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_Card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_Card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Card) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Card) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Card) GetFlavor() string {
	if x != nil {
		return x.Flavor
	}
	return ""
}

func (x *Card) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Card) GetType() CardType {
	if x != nil {
		return x.Type
	}
	return CardType_ATTACK
}

func (x *Card) GetAction() *CardAction {
	if x != nil {
		return x.Action
	}
	return nil
}

var File_Card_proto protoreflect.FileDescriptor

var file_Card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x10, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c,
	0x61, 0x76, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x76,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x64, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Card_proto_rawDescOnce sync.Once
	file_Card_proto_rawDescData = file_Card_proto_rawDesc
)

func file_Card_proto_rawDescGZIP() []byte {
	file_Card_proto_rawDescOnce.Do(func() {
		file_Card_proto_rawDescData = protoimpl.X.CompressGZIP(file_Card_proto_rawDescData)
	})
	return file_Card_proto_rawDescData
}

var file_Card_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Card_proto_goTypes = []interface{}{
	(*Card)(nil),       // 0: Deviant.Card
	(CardType)(0),      // 1: Deviant.CardType
	(*CardAction)(nil), // 2: Deviant.CardAction
}
var file_Card_proto_depIdxs = []int32{
	1, // 0: Deviant.Card.type:type_name -> Deviant.CardType
	2, // 1: Deviant.Card.action:type_name -> Deviant.CardAction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Card_proto_init() }
func file_Card_proto_init() {
	if File_Card_proto != nil {
		return
	}
	file_CardAction_proto_init()
	file_CardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
			RawDescriptor: file_Card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Card_proto_goTypes,
		DependencyIndexes: file_Card_proto_depIdxs,
		MessageInfos:      file_Card_proto_msgTypes,
	}.Build()
	File_Card_proto = out.File
	file_Card_proto_rawDesc = nil
	file_Card_proto_goTypes = nil
	file_Card_proto_depIdxs = nil
}
