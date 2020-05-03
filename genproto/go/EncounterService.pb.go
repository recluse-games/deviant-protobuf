// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: EncounterService.proto

package deviant

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type EncounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId         string            `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Encounter        *Encounter        `protobuf:"bytes,2,opt,name=encounter,proto3" json:"encounter,omitempty"`
	EntityActionName EntityActionNames `protobuf:"varint,3,opt,name=entityActionName,proto3,enum=Deviant.EntityActionNames" json:"entityActionName,omitempty"`
}

func (x *EncounterRequest) Reset() {
	*x = EncounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EncounterService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterRequest) ProtoMessage() {}

func (x *EncounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EncounterService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterRequest.ProtoReflect.Descriptor instead.
func (*EncounterRequest) Descriptor() ([]byte, []int) {
	return file_EncounterService_proto_rawDescGZIP(), []int{0}
}

func (x *EncounterRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *EncounterRequest) GetEncounter() *Encounter {
	if x != nil {
		return x.Encounter
	}
	return nil
}

func (x *EncounterRequest) GetEntityActionName() EntityActionNames {
	if x != nil {
		return x.EntityActionName
	}
	return EntityActionNames_PLAY
}

// The response message containing the greetings
type EncounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string     `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Encounter *Encounter `protobuf:"bytes,2,opt,name=encounter,proto3" json:"encounter,omitempty"`
}

func (x *EncounterResponse) Reset() {
	*x = EncounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EncounterService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterResponse) ProtoMessage() {}

func (x *EncounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EncounterService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterResponse.ProtoReflect.Descriptor instead.
func (*EncounterResponse) Descriptor() ([]byte, []int) {
	return file_EncounterService_proto_rawDescGZIP(), []int{1}
}

func (x *EncounterResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *EncounterResponse) GetEncounter() *Encounter {
	if x != nil {
		return x.Encounter
	}
	return nil
}

var File_EncounterService_proto protoreflect.FileDescriptor

var file_EncounterService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e,
	0x74, 0x1a, 0x0f, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x10,
	0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x09,
	0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x46,
	0x0a, 0x10, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61,
	0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x10, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x11, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x09,
	0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x32, 0xfd, 0x01, 0x0a, 0x10, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EncounterService_proto_rawDescOnce sync.Once
	file_EncounterService_proto_rawDescData = file_EncounterService_proto_rawDesc
)

func file_EncounterService_proto_rawDescGZIP() []byte {
	file_EncounterService_proto_rawDescOnce.Do(func() {
		file_EncounterService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EncounterService_proto_rawDescData)
	})
	return file_EncounterService_proto_rawDescData
}

var file_EncounterService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_EncounterService_proto_goTypes = []interface{}{
	(*EncounterRequest)(nil),  // 0: Deviant.EncounterRequest
	(*EncounterResponse)(nil), // 1: Deviant.EncounterResponse
	(*Encounter)(nil),         // 2: Deviant.Encounter
	(EntityActionNames)(0),    // 3: Deviant.EntityActionNames
}
var file_EncounterService_proto_depIdxs = []int32{
	2, // 0: Deviant.EncounterRequest.encounter:type_name -> Deviant.Encounter
	3, // 1: Deviant.EncounterRequest.entityActionName:type_name -> Deviant.EntityActionNames
	2, // 2: Deviant.EncounterResponse.encounter:type_name -> Deviant.Encounter
	0, // 3: Deviant.EncounterService.StartEncounter:input_type -> Deviant.EncounterRequest
	0, // 4: Deviant.EncounterService.UpdateEncounter:input_type -> Deviant.EncounterRequest
	0, // 5: Deviant.EncounterService.CompleteEncounter:input_type -> Deviant.EncounterRequest
	1, // 6: Deviant.EncounterService.StartEncounter:output_type -> Deviant.EncounterResponse
	1, // 7: Deviant.EncounterService.UpdateEncounter:output_type -> Deviant.EncounterResponse
	1, // 8: Deviant.EncounterService.CompleteEncounter:output_type -> Deviant.EncounterResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_EncounterService_proto_init() }
func file_EncounterService_proto_init() {
	if File_EncounterService_proto != nil {
		return
	}
	file_Encounter_proto_init()
	file_EntityActionNames_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EncounterService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterRequest); i {
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
		file_EncounterService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterResponse); i {
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
			RawDescriptor: file_EncounterService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_EncounterService_proto_goTypes,
		DependencyIndexes: file_EncounterService_proto_depIdxs,
		MessageInfos:      file_EncounterService_proto_msgTypes,
	}.Build()
	File_EncounterService_proto = out.File
	file_EncounterService_proto_rawDesc = nil
	file_EncounterService_proto_goTypes = nil
	file_EncounterService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EncounterServiceClient is the client API for EncounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EncounterServiceClient interface {
	StartEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_StartEncounterClient, error)
	UpdateEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_UpdateEncounterClient, error)
	CompleteEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_CompleteEncounterClient, error)
}

type encounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterServiceClient(cc grpc.ClientConnInterface) EncounterServiceClient {
	return &encounterServiceClient{cc}
}

func (c *encounterServiceClient) StartEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_StartEncounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EncounterService_serviceDesc.Streams[0], "/Deviant.EncounterService/StartEncounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &encounterServiceStartEncounterClient{stream}
	return x, nil
}

type EncounterService_StartEncounterClient interface {
	Send(*EncounterRequest) error
	Recv() (*EncounterResponse, error)
	grpc.ClientStream
}

type encounterServiceStartEncounterClient struct {
	grpc.ClientStream
}

func (x *encounterServiceStartEncounterClient) Send(m *EncounterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *encounterServiceStartEncounterClient) Recv() (*EncounterResponse, error) {
	m := new(EncounterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *encounterServiceClient) UpdateEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_UpdateEncounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EncounterService_serviceDesc.Streams[1], "/Deviant.EncounterService/UpdateEncounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &encounterServiceUpdateEncounterClient{stream}
	return x, nil
}

type EncounterService_UpdateEncounterClient interface {
	Send(*EncounterRequest) error
	Recv() (*EncounterResponse, error)
	grpc.ClientStream
}

type encounterServiceUpdateEncounterClient struct {
	grpc.ClientStream
}

func (x *encounterServiceUpdateEncounterClient) Send(m *EncounterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *encounterServiceUpdateEncounterClient) Recv() (*EncounterResponse, error) {
	m := new(EncounterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *encounterServiceClient) CompleteEncounter(ctx context.Context, opts ...grpc.CallOption) (EncounterService_CompleteEncounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EncounterService_serviceDesc.Streams[2], "/Deviant.EncounterService/CompleteEncounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &encounterServiceCompleteEncounterClient{stream}
	return x, nil
}

type EncounterService_CompleteEncounterClient interface {
	Send(*EncounterRequest) error
	Recv() (*EncounterResponse, error)
	grpc.ClientStream
}

type encounterServiceCompleteEncounterClient struct {
	grpc.ClientStream
}

func (x *encounterServiceCompleteEncounterClient) Send(m *EncounterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *encounterServiceCompleteEncounterClient) Recv() (*EncounterResponse, error) {
	m := new(EncounterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EncounterServiceServer is the server API for EncounterService service.
type EncounterServiceServer interface {
	StartEncounter(EncounterService_StartEncounterServer) error
	UpdateEncounter(EncounterService_UpdateEncounterServer) error
	CompleteEncounter(EncounterService_CompleteEncounterServer) error
}

// UnimplementedEncounterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (*UnimplementedEncounterServiceServer) StartEncounter(EncounterService_StartEncounterServer) error {
	return status.Errorf(codes.Unimplemented, "method StartEncounter not implemented")
}
func (*UnimplementedEncounterServiceServer) UpdateEncounter(EncounterService_UpdateEncounterServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateEncounter not implemented")
}
func (*UnimplementedEncounterServiceServer) CompleteEncounter(EncounterService_CompleteEncounterServer) error {
	return status.Errorf(codes.Unimplemented, "method CompleteEncounter not implemented")
}

func RegisterEncounterServiceServer(s *grpc.Server, srv EncounterServiceServer) {
	s.RegisterService(&_EncounterService_serviceDesc, srv)
}

func _EncounterService_StartEncounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EncounterServiceServer).StartEncounter(&encounterServiceStartEncounterServer{stream})
}

type EncounterService_StartEncounterServer interface {
	Send(*EncounterResponse) error
	Recv() (*EncounterRequest, error)
	grpc.ServerStream
}

type encounterServiceStartEncounterServer struct {
	grpc.ServerStream
}

func (x *encounterServiceStartEncounterServer) Send(m *EncounterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *encounterServiceStartEncounterServer) Recv() (*EncounterRequest, error) {
	m := new(EncounterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EncounterService_UpdateEncounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EncounterServiceServer).UpdateEncounter(&encounterServiceUpdateEncounterServer{stream})
}

type EncounterService_UpdateEncounterServer interface {
	Send(*EncounterResponse) error
	Recv() (*EncounterRequest, error)
	grpc.ServerStream
}

type encounterServiceUpdateEncounterServer struct {
	grpc.ServerStream
}

func (x *encounterServiceUpdateEncounterServer) Send(m *EncounterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *encounterServiceUpdateEncounterServer) Recv() (*EncounterRequest, error) {
	m := new(EncounterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EncounterService_CompleteEncounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EncounterServiceServer).CompleteEncounter(&encounterServiceCompleteEncounterServer{stream})
}

type EncounterService_CompleteEncounterServer interface {
	Send(*EncounterResponse) error
	Recv() (*EncounterRequest, error)
	grpc.ServerStream
}

type encounterServiceCompleteEncounterServer struct {
	grpc.ServerStream
}

func (x *encounterServiceCompleteEncounterServer) Send(m *EncounterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *encounterServiceCompleteEncounterServer) Recv() (*EncounterRequest, error) {
	m := new(EncounterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EncounterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Deviant.EncounterService",
	HandlerType: (*EncounterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartEncounter",
			Handler:       _EncounterService_StartEncounter_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateEncounter",
			Handler:       _EncounterService_UpdateEncounter_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CompleteEncounter",
			Handler:       _EncounterService_CompleteEncounter_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "EncounterService.proto",
}
