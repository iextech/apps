// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/ansibleapp.proto

package ansibleapp

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

type AnsibleAppOutput_FileType int32

const (
	AnsibleAppOutput_TEXT AnsibleAppOutput_FileType = 0
	AnsibleAppOutput_TAR  AnsibleAppOutput_FileType = 1
	AnsibleAppOutput_TGZ  AnsibleAppOutput_FileType = 2
)

// Enum value maps for AnsibleAppOutput_FileType.
var (
	AnsibleAppOutput_FileType_name = map[int32]string{
		0: "TEXT",
		1: "TAR",
		2: "TGZ",
	}
	AnsibleAppOutput_FileType_value = map[string]int32{
		"TEXT": 0,
		"TAR":  1,
		"TGZ":  2,
	}
)

func (x AnsibleAppOutput_FileType) Enum() *AnsibleAppOutput_FileType {
	p := new(AnsibleAppOutput_FileType)
	*p = x
	return p
}

func (x AnsibleAppOutput_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnsibleAppOutput_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ansibleapp_proto_enumTypes[0].Descriptor()
}

func (AnsibleAppOutput_FileType) Type() protoreflect.EnumType {
	return &file_proto_ansibleapp_proto_enumTypes[0]
}

func (x AnsibleAppOutput_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnsibleAppOutput_FileType.Descriptor instead.
func (AnsibleAppOutput_FileType) EnumDescriptor() ([]byte, []int) {
	return file_proto_ansibleapp_proto_rawDescGZIP(), []int{1, 0}
}

// AnsibleApp
type AnsibleAppInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *AnsibleAppInput) Reset() {
	*x = AnsibleAppInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ansibleapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnsibleAppInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnsibleAppInput) ProtoMessage() {}

func (x *AnsibleAppInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ansibleapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnsibleAppInput.ProtoReflect.Descriptor instead.
func (*AnsibleAppInput) Descriptor() ([]byte, []int) {
	return file_proto_ansibleapp_proto_rawDescGZIP(), []int{0}
}

func (x *AnsibleAppInput) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type AnsibleAppOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskFileContent   string                    `protobuf:"bytes,1,opt,name=task_file_content,json=taskFileContent,proto3" json:"task_file_content,omitempty"`
	BinaryFileContent []byte                    `protobuf:"bytes,2,opt,name=binary_file_content,json=binaryFileContent,proto3" json:"binary_file_content,omitempty"`
	FileType          AnsibleAppOutput_FileType `protobuf:"varint,10,opt,name=file_type,json=fileType,proto3,enum=ansibleapp.AnsibleAppOutput_FileType" json:"file_type,omitempty"`
}

func (x *AnsibleAppOutput) Reset() {
	*x = AnsibleAppOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ansibleapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnsibleAppOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnsibleAppOutput) ProtoMessage() {}

func (x *AnsibleAppOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ansibleapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnsibleAppOutput.ProtoReflect.Descriptor instead.
func (*AnsibleAppOutput) Descriptor() ([]byte, []int) {
	return file_proto_ansibleapp_proto_rawDescGZIP(), []int{1}
}

func (x *AnsibleAppOutput) GetTaskFileContent() string {
	if x != nil {
		return x.TaskFileContent
	}
	return ""
}

func (x *AnsibleAppOutput) GetBinaryFileContent() []byte {
	if x != nil {
		return x.BinaryFileContent
	}
	return nil
}

func (x *AnsibleAppOutput) GetFileType() AnsibleAppOutput_FileType {
	if x != nil {
		return x.FileType
	}
	return AnsibleAppOutput_TEXT
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ansibleapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ansibleapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_proto_ansibleapp_proto_rawDescGZIP(), []int{2}
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type HelpText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *HelpText) Reset() {
	*x = HelpText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ansibleapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelpText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelpText) ProtoMessage() {}

func (x *HelpText) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ansibleapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelpText.ProtoReflect.Descriptor instead.
func (*HelpText) Descriptor() ([]byte, []int) {
	return file_proto_ansibleapp_proto_rawDescGZIP(), []int{3}
}

func (x *HelpText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_proto_ansibleapp_proto protoreflect.FileDescriptor

var file_proto_ansibleapp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x61, 0x70, 0x70, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0xda, 0x01, 0x0a, 0x10,
	0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x73,
	0x6b, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x26, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x41, 0x52, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x47, 0x5a, 0x10, 0x02, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x70, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x87, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e,
	0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2f, 0x0a, 0x04, 0x48, 0x65, 0x6c, 0x70, 0x12, 0x11, 0x2e, 0x61, 0x6e, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x14, 0x2e, 0x61, 0x6e,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x54, 0x65, 0x78,
	0x74, 0x42, 0x57, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x42, 0x13, 0x41, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_ansibleapp_proto_rawDescOnce sync.Once
	file_proto_ansibleapp_proto_rawDescData = file_proto_ansibleapp_proto_rawDesc
)

func file_proto_ansibleapp_proto_rawDescGZIP() []byte {
	file_proto_ansibleapp_proto_rawDescOnce.Do(func() {
		file_proto_ansibleapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ansibleapp_proto_rawDescData)
	})
	return file_proto_ansibleapp_proto_rawDescData
}

var file_proto_ansibleapp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ansibleapp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_ansibleapp_proto_goTypes = []interface{}{
	(AnsibleAppOutput_FileType)(0), // 0: ansibleapp.AnsibleAppOutput.FileType
	(*AnsibleAppInput)(nil),        // 1: ansibleapp.AnsibleAppInput
	(*AnsibleAppOutput)(nil),       // 2: ansibleapp.AnsibleAppOutput
	(*Topic)(nil),                  // 3: ansibleapp.Topic
	(*HelpText)(nil),               // 4: ansibleapp.HelpText
}
var file_proto_ansibleapp_proto_depIdxs = []int32{
	0, // 0: ansibleapp.AnsibleAppOutput.file_type:type_name -> ansibleapp.AnsibleAppOutput.FileType
	1, // 1: ansibleapp.AnsibleAppService.Exec:input_type -> ansibleapp.AnsibleAppInput
	3, // 2: ansibleapp.AnsibleAppService.Help:input_type -> ansibleapp.Topic
	2, // 3: ansibleapp.AnsibleAppService.Exec:output_type -> ansibleapp.AnsibleAppOutput
	4, // 4: ansibleapp.AnsibleAppService.Help:output_type -> ansibleapp.HelpText
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ansibleapp_proto_init() }
func file_proto_ansibleapp_proto_init() {
	if File_proto_ansibleapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ansibleapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnsibleAppInput); i {
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
		file_proto_ansibleapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnsibleAppOutput); i {
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
		file_proto_ansibleapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_proto_ansibleapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelpText); i {
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
			RawDescriptor: file_proto_ansibleapp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ansibleapp_proto_goTypes,
		DependencyIndexes: file_proto_ansibleapp_proto_depIdxs,
		EnumInfos:         file_proto_ansibleapp_proto_enumTypes,
		MessageInfos:      file_proto_ansibleapp_proto_msgTypes,
	}.Build()
	File_proto_ansibleapp_proto = out.File
	file_proto_ansibleapp_proto_rawDesc = nil
	file_proto_ansibleapp_proto_goTypes = nil
	file_proto_ansibleapp_proto_depIdxs = nil
}
