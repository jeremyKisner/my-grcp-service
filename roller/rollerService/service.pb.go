// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: roller/rollerService/service.proto

package rollerService

import (
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

type RollerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RollerRequest) Reset() {
	*x = RollerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roller_rollerService_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollerRequest) ProtoMessage() {}

func (x *RollerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_roller_rollerService_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollerRequest.ProtoReflect.Descriptor instead.
func (*RollerRequest) Descriptor() ([]byte, []int) {
	return file_roller_rollerService_service_proto_rawDescGZIP(), []int{0}
}

func (x *RollerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RollerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Total       string `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	PlayerScore int64  `protobuf:"varint,3,opt,name=playerScore,proto3" json:"playerScore,omitempty"`
	AiScore     int64  `protobuf:"varint,4,opt,name=aiScore,proto3" json:"aiScore,omitempty"`
	PlayerWins  bool   `protobuf:"varint,5,opt,name=playerWins,proto3" json:"playerWins,omitempty"`
}

func (x *RollerReply) Reset() {
	*x = RollerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roller_rollerService_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollerReply) ProtoMessage() {}

func (x *RollerReply) ProtoReflect() protoreflect.Message {
	mi := &file_roller_rollerService_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollerReply.ProtoReflect.Descriptor instead.
func (*RollerReply) Descriptor() ([]byte, []int) {
	return file_roller_rollerService_service_proto_rawDescGZIP(), []int{1}
}

func (x *RollerReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RollerReply) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *RollerReply) GetPlayerScore() int64 {
	if x != nil {
		return x.PlayerScore
	}
	return 0
}

func (x *RollerReply) GetAiScore() int64 {
	if x != nil {
		return x.AiScore
	}
	return 0
}

func (x *RollerReply) GetPlayerWins() bool {
	if x != nil {
		return x.PlayerWins
	}
	return false
}

var File_roller_rollerService_service_proto protoreflect.FileDescriptor

var file_roller_rollerService_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x69,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x69, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x69,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x57, 0x69, 0x6e, 0x73, 0x32, 0x4c, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x19, 0x50, 0x01, 0x5a, 0x15, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roller_rollerService_service_proto_rawDescOnce sync.Once
	file_roller_rollerService_service_proto_rawDescData = file_roller_rollerService_service_proto_rawDesc
)

func file_roller_rollerService_service_proto_rawDescGZIP() []byte {
	file_roller_rollerService_service_proto_rawDescOnce.Do(func() {
		file_roller_rollerService_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_roller_rollerService_service_proto_rawDescData)
	})
	return file_roller_rollerService_service_proto_rawDescData
}

var file_roller_rollerService_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_roller_rollerService_service_proto_goTypes = []interface{}{
	(*RollerRequest)(nil), // 0: rollerService.RollerRequest
	(*RollerReply)(nil),   // 1: rollerService.RollerReply
}
var file_roller_rollerService_service_proto_depIdxs = []int32{
	0, // 0: rollerService.Roller.Roll:input_type -> rollerService.RollerRequest
	1, // 1: rollerService.Roller.Roll:output_type -> rollerService.RollerReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_roller_rollerService_service_proto_init() }
func file_roller_rollerService_service_proto_init() {
	if File_roller_rollerService_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roller_rollerService_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollerRequest); i {
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
		file_roller_rollerService_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollerReply); i {
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
			RawDescriptor: file_roller_rollerService_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roller_rollerService_service_proto_goTypes,
		DependencyIndexes: file_roller_rollerService_service_proto_depIdxs,
		MessageInfos:      file_roller_rollerService_service_proto_msgTypes,
	}.Build()
	File_roller_rollerService_service_proto = out.File
	file_roller_rollerService_service_proto_rawDesc = nil
	file_roller_rollerService_service_proto_goTypes = nil
	file_roller_rollerService_service_proto_depIdxs = nil
}
