// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: chittychat.proto

package disysminiproject2

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

type MessageWithLamport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *MessageWithLamport) Reset() {
	*x = MessageWithLamport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithLamport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithLamport) ProtoMessage() {}

func (x *MessageWithLamport) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithLamport.ProtoReflect.Descriptor instead.
func (*MessageWithLamport) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithLamport) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageWithLamport) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *MessageWithLamport) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{1}
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Lamport int32  `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectionRequest) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Lamport int32 `protobuf:"varint,2,opt,name=Lamport,proto3" json:"Lamport,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaveRequest) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

var File_chittychat_proto protoreflect.FileDescriptor

var file_chittychat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x22, 0x58,
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x41, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x38, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xe9,
	0x02, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3e, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x13, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x69, 0x65, 0x76, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x11, 0x2e,
	0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x68,
	0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f,
	0x64, 0x69, 0x73, 0x79, 0x73, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chittychat_proto_rawDescOnce sync.Once
	file_chittychat_proto_rawDescData = file_chittychat_proto_rawDesc
)

func file_chittychat_proto_rawDescGZIP() []byte {
	file_chittychat_proto_rawDescOnce.Do(func() {
		file_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chittychat_proto_rawDescData)
	})
	return file_chittychat_proto_rawDescData
}

var file_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chittychat_proto_goTypes = []interface{}{
	(*MessageWithLamport)(nil), // 0: chittychat.MessageWithLamport
	(*Empty)(nil),              // 1: chittychat.Empty
	(*ConnectionRequest)(nil),  // 2: chittychat.ConnectionRequest
	(*LeaveRequest)(nil),       // 3: chittychat.LeaveRequest
}
var file_chittychat_proto_depIdxs = []int32{
	0, // 0: chittychat.ChittyChat.Publish:input_type -> chittychat.MessageWithLamport
	0, // 1: chittychat.ChittyChat.Broadcast:input_type -> chittychat.MessageWithLamport
	2, // 2: chittychat.ChittyChat.EstablishConnection:input_type -> chittychat.ConnectionRequest
	0, // 3: chittychat.ChittyChat.RecieveBroadcast:input_type -> chittychat.MessageWithLamport
	3, // 4: chittychat.ChittyChat.Leave:input_type -> chittychat.LeaveRequest
	1, // 5: chittychat.ChittyChat.Publish:output_type -> chittychat.Empty
	1, // 6: chittychat.ChittyChat.Broadcast:output_type -> chittychat.Empty
	0, // 7: chittychat.ChittyChat.EstablishConnection:output_type -> chittychat.MessageWithLamport
	1, // 8: chittychat.ChittyChat.RecieveBroadcast:output_type -> chittychat.Empty
	1, // 9: chittychat.ChittyChat.Leave:output_type -> chittychat.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chittychat_proto_init() }
func file_chittychat_proto_init() {
	if File_chittychat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chittychat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithLamport); i {
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
		file_chittychat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_chittychat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_chittychat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
			RawDescriptor: file_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chittychat_proto_goTypes,
		DependencyIndexes: file_chittychat_proto_depIdxs,
		MessageInfos:      file_chittychat_proto_msgTypes,
	}.Build()
	File_chittychat_proto = out.File
	file_chittychat_proto_rawDesc = nil
	file_chittychat_proto_goTypes = nil
	file_chittychat_proto_depIdxs = nil
}
