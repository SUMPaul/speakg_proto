// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: stream_proxy/chat.proto

package stream_gateway

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

type LearnConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceSpeed float32 `protobuf:"fixed32,1,opt,name=voice_speed,json=voiceSpeed,proto3" json:"voice_speed,omitempty"`
	GoogleName string  `protobuf:"bytes,2,opt,name=google_name,json=googleName,proto3" json:"google_name,omitempty"`
	Pitch      float32 `protobuf:"fixed32,3,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Language   string  `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Gender     string  `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *LearnConfig) Reset() {
	*x = LearnConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proxy_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LearnConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LearnConfig) ProtoMessage() {}

func (x *LearnConfig) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proxy_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LearnConfig.ProtoReflect.Descriptor instead.
func (*LearnConfig) Descriptor() ([]byte, []int) {
	return file_stream_proxy_chat_proto_rawDescGZIP(), []int{0}
}

func (x *LearnConfig) GetVoiceSpeed() float32 {
	if x != nil {
		return x.VoiceSpeed
	}
	return 0
}

func (x *LearnConfig) GetGoogleName() string {
	if x != nil {
		return x.GoogleName
	}
	return ""
}

func (x *LearnConfig) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *LearnConfig) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *LearnConfig) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type ProxyChatGPTStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content     string       `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	LearnConfig *LearnConfig `protobuf:"bytes,2,opt,name=learn_config,json=learnConfig,proto3" json:"learn_config,omitempty"`
}

func (x *ProxyChatGPTStreamReq) Reset() {
	*x = ProxyChatGPTStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proxy_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyChatGPTStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyChatGPTStreamReq) ProtoMessage() {}

func (x *ProxyChatGPTStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proxy_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyChatGPTStreamReq.ProtoReflect.Descriptor instead.
func (*ProxyChatGPTStreamReq) Descriptor() ([]byte, []int) {
	return file_stream_proxy_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyChatGPTStreamReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ProxyChatGPTStreamReq) GetLearnConfig() *LearnConfig {
	if x != nil {
		return x.LearnConfig
	}
	return nil
}

type ProxyChatGPTStreamResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content          string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AudioContent     string `protobuf:"bytes,3,opt,name=audioContent,proto3" json:"audioContent,omitempty"`
	Url              string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	FullContent      string `protobuf:"bytes,5,opt,name=full_content,json=fullContent,proto3" json:"full_content,omitempty"`
	FullAudioContent string `protobuf:"bytes,6,opt,name=full_audioContent,json=fullAudioContent,proto3" json:"full_audioContent,omitempty"`
	FullUrl          string `protobuf:"bytes,7,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"` // Add more fields as needed
}

func (x *ProxyChatGPTStreamResp) Reset() {
	*x = ProxyChatGPTStreamResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proxy_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyChatGPTStreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyChatGPTStreamResp) ProtoMessage() {}

func (x *ProxyChatGPTStreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proxy_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyChatGPTStreamResp.ProtoReflect.Descriptor instead.
func (*ProxyChatGPTStreamResp) Descriptor() ([]byte, []int) {
	return file_stream_proxy_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyChatGPTStreamResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetAudioContent() string {
	if x != nil {
		return x.AudioContent
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetFullContent() string {
	if x != nil {
		return x.FullContent
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetFullAudioContent() string {
	if x != nil {
		return x.FullAudioContent
	}
	return ""
}

func (x *ProxyChatGPTStreamResp) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

var File_stream_proxy_chat_proto protoreflect.FileDescriptor

var file_stream_proxy_chat_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x67, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x99,
	0x01, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x15, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x43, 0x68, 0x61, 0x74, 0x47, 0x50, 0x54, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a,
	0x0c, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x67, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xe3, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x68, 0x61, 0x74,
	0x47, 0x50, 0x54, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x75, 0x6c,
	0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0x87, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x71, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x68, 0x61, 0x74, 0x47, 0x50, 0x54, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x67, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x43, 0x68, 0x61, 0x74, 0x47, 0x50, 0x54, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x1a, 0x2b, 0x2e, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x47, 0x50, 0x54, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x55, 0x4d, 0x50, 0x61, 0x75, 0x6c, 0x2f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x33, 0x2e,
	0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proxy_chat_proto_rawDescOnce sync.Once
	file_stream_proxy_chat_proto_rawDescData = file_stream_proxy_chat_proto_rawDesc
)

func file_stream_proxy_chat_proto_rawDescGZIP() []byte {
	file_stream_proxy_chat_proto_rawDescOnce.Do(func() {
		file_stream_proxy_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proxy_chat_proto_rawDescData)
	})
	return file_stream_proxy_chat_proto_rawDescData
}

var file_stream_proxy_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stream_proxy_chat_proto_goTypes = []interface{}{
	(*LearnConfig)(nil),            // 0: speakg.stream_proxy.LearnConfig
	(*ProxyChatGPTStreamReq)(nil),  // 1: speakg.stream_proxy.ProxyChatGPTStreamReq
	(*ProxyChatGPTStreamResp)(nil), // 2: speakg.stream_proxy.ProxyChatGPTStreamResp
}
var file_stream_proxy_chat_proto_depIdxs = []int32{
	0, // 0: speakg.stream_proxy.ProxyChatGPTStreamReq.learn_config:type_name -> speakg.stream_proxy.LearnConfig
	1, // 1: speakg.stream_proxy.StreamProxyService.ProxyChatGPTStream:input_type -> speakg.stream_proxy.ProxyChatGPTStreamReq
	2, // 2: speakg.stream_proxy.StreamProxyService.ProxyChatGPTStream:output_type -> speakg.stream_proxy.ProxyChatGPTStreamResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stream_proxy_chat_proto_init() }
func file_stream_proxy_chat_proto_init() {
	if File_stream_proxy_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proxy_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LearnConfig); i {
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
		file_stream_proxy_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyChatGPTStreamReq); i {
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
		file_stream_proxy_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyChatGPTStreamResp); i {
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
			RawDescriptor: file_stream_proxy_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proxy_chat_proto_goTypes,
		DependencyIndexes: file_stream_proxy_chat_proto_depIdxs,
		MessageInfos:      file_stream_proxy_chat_proto_msgTypes,
	}.Build()
	File_stream_proxy_chat_proto = out.File
	file_stream_proxy_chat_proto_rawDesc = nil
	file_stream_proxy_chat_proto_goTypes = nil
	file_stream_proxy_chat_proto_depIdxs = nil
}
