// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/entrypoint.proto

package entrypoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level   string                 `protobuf:"bytes,1,opt,name=Level,proto3" json:"Level,omitempty"`
	Prefix  string                 `protobuf:"bytes,2,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	IP      string                 `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	Agent   string                 `protobuf:"bytes,4,opt,name=Agent,proto3" json:"Agent,omitempty"`
	Message string                 `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	Trace   []*Frame               `protobuf:"bytes,6,rep,name=Trace,proto3" json:"Trace,omitempty"`
	Source  *FrameWithCode         `protobuf:"bytes,7,opt,name=Source,proto3" json:"Source,omitempty"`
	Time    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=Time,proto3" json:"Time,omitempty"`
	Token   string                 `protobuf:"bytes,9,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *PayloadRequest) Reset() {
	*x = PayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_entrypoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadRequest) ProtoMessage() {}

func (x *PayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_entrypoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadRequest.ProtoReflect.Descriptor instead.
func (*PayloadRequest) Descriptor() ([]byte, []int) {
	return file_proto_entrypoint_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *PayloadRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *PayloadRequest) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *PayloadRequest) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *PayloadRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PayloadRequest) GetTrace() []*Frame {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *PayloadRequest) GetSource() *FrameWithCode {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *PayloadRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *PayloadRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Line uint32 `protobuf:"varint,2,opt,name=Line,proto3" json:"Line,omitempty"`
	Func string `protobuf:"bytes,3,opt,name=Func,proto3" json:"Func,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_entrypoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_proto_entrypoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_proto_entrypoint_proto_rawDescGZIP(), []int{1}
}

func (x *Frame) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Frame) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Frame) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

type FrameWithCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Line uint32 `protobuf:"varint,2,opt,name=Line,proto3" json:"Line,omitempty"`
	Func string `protobuf:"bytes,3,opt,name=Func,proto3" json:"Func,omitempty"`
	Code string `protobuf:"bytes,4,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *FrameWithCode) Reset() {
	*x = FrameWithCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_entrypoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameWithCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameWithCode) ProtoMessage() {}

func (x *FrameWithCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_entrypoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameWithCode.ProtoReflect.Descriptor instead.
func (*FrameWithCode) Descriptor() ([]byte, []int) {
	return file_proto_entrypoint_proto_rawDescGZIP(), []int{2}
}

func (x *FrameWithCode) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FrameWithCode) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *FrameWithCode) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

func (x *FrameWithCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type PayloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PayloadReply) Reset() {
	*x = PayloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_entrypoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadReply) ProtoMessage() {}

func (x *PayloadReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_entrypoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadReply.ProtoReflect.Descriptor instead.
func (*PayloadReply) Descriptor() ([]byte, []int) {
	return file_proto_entrypoint_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_entrypoint_proto protoreflect.FileDescriptor

var file_proto_entrypoint_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x21, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x75, 0x6e, 0x63, 0x22, 0x5f, 0x0a, 0x0d, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x0c,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x41, 0x73, 0x79,
	0x6e, 0x63, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x11,
	0x5a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_entrypoint_proto_rawDescOnce sync.Once
	file_proto_entrypoint_proto_rawDescData = file_proto_entrypoint_proto_rawDesc
)

func file_proto_entrypoint_proto_rawDescGZIP() []byte {
	file_proto_entrypoint_proto_rawDescOnce.Do(func() {
		file_proto_entrypoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_entrypoint_proto_rawDescData)
	})
	return file_proto_entrypoint_proto_rawDescData
}

var file_proto_entrypoint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_entrypoint_proto_goTypes = []interface{}{
	(*PayloadRequest)(nil),        // 0: grpc.PayloadRequest
	(*Frame)(nil),                 // 1: grpc.Frame
	(*FrameWithCode)(nil),         // 2: grpc.FrameWithCode
	(*PayloadReply)(nil),          // 3: grpc.PayloadReply
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_entrypoint_proto_depIdxs = []int32{
	1, // 0: grpc.PayloadRequest.Trace:type_name -> grpc.Frame
	2, // 1: grpc.PayloadRequest.Source:type_name -> grpc.FrameWithCode
	4, // 2: grpc.PayloadRequest.Time:type_name -> google.protobuf.Timestamp
	0, // 3: grpc.Entrypoint.WriteLogSync:input_type -> grpc.PayloadRequest
	0, // 4: grpc.Entrypoint.WriteLogAsync:input_type -> grpc.PayloadRequest
	3, // 5: grpc.Entrypoint.WriteLogSync:output_type -> grpc.PayloadReply
	3, // 6: grpc.Entrypoint.WriteLogAsync:output_type -> grpc.PayloadReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_entrypoint_proto_init() }
func file_proto_entrypoint_proto_init() {
	if File_proto_entrypoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_entrypoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadRequest); i {
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
		file_proto_entrypoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_proto_entrypoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameWithCode); i {
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
		file_proto_entrypoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadReply); i {
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
			RawDescriptor: file_proto_entrypoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_entrypoint_proto_goTypes,
		DependencyIndexes: file_proto_entrypoint_proto_depIdxs,
		MessageInfos:      file_proto_entrypoint_proto_msgTypes,
	}.Build()
	File_proto_entrypoint_proto = out.File
	file_proto_entrypoint_proto_rawDesc = nil
	file_proto_entrypoint_proto_goTypes = nil
	file_proto_entrypoint_proto_depIdxs = nil
}
