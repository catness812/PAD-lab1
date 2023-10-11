// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: journal_svc.proto

package pb

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_journal_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_journal_svc_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RegisterEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *RegisterEntryRequest) Reset() {
	*x = RegisterEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterEntryRequest) ProtoMessage() {}

func (x *RegisterEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journal_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterEntryRequest.ProtoReflect.Descriptor instead.
func (*RegisterEntryRequest) Descriptor() ([]byte, []int) {
	return file_journal_svc_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterEntryRequest) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type RegisterEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterEntryResponse) Reset() {
	*x = RegisterEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterEntryResponse) ProtoMessage() {}

func (x *RegisterEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_journal_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterEntryResponse.ProtoReflect.Descriptor instead.
func (*RegisterEntryResponse) Descriptor() ([]byte, []int) {
	return file_journal_svc_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterEntryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_journal_svc_proto protoreflect.FileDescriptor

var file_journal_svc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x3a, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x31, 0x0a, 0x15, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5c,
	0x0a, 0x0e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_journal_svc_proto_rawDescOnce sync.Once
	file_journal_svc_proto_rawDescData = file_journal_svc_proto_rawDesc
)

func file_journal_svc_proto_rawDescGZIP() []byte {
	file_journal_svc_proto_rawDescOnce.Do(func() {
		file_journal_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_journal_svc_proto_rawDescData)
	})
	return file_journal_svc_proto_rawDescData
}

var file_journal_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_journal_svc_proto_goTypes = []interface{}{
	(*Entry)(nil),                 // 0: proto.Entry
	(*RegisterEntryRequest)(nil),  // 1: proto.RegisterEntryRequest
	(*RegisterEntryResponse)(nil), // 2: proto.RegisterEntryResponse
}
var file_journal_svc_proto_depIdxs = []int32{
	0, // 0: proto.RegisterEntryRequest.entry:type_name -> proto.Entry
	1, // 1: proto.JournalService.RegisterEntry:input_type -> proto.RegisterEntryRequest
	2, // 2: proto.JournalService.RegisterEntry:output_type -> proto.RegisterEntryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_journal_svc_proto_init() }
func file_journal_svc_proto_init() {
	if File_journal_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_journal_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_journal_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterEntryRequest); i {
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
		file_journal_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterEntryResponse); i {
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
			RawDescriptor: file_journal_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_journal_svc_proto_goTypes,
		DependencyIndexes: file_journal_svc_proto_depIdxs,
		MessageInfos:      file_journal_svc_proto_msgTypes,
	}.Build()
	File_journal_svc_proto = out.File
	file_journal_svc_proto_rawDesc = nil
	file_journal_svc_proto_goTypes = nil
	file_journal_svc_proto_depIdxs = nil
}