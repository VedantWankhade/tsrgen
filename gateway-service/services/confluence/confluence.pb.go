// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: confluence.proto

package confluence

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

type PageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConfluenceInstance string `protobuf:"bytes,2,opt,name=confluenceInstance,proto3" json:"confluenceInstance,omitempty"`
	ConfleunceUsername string `protobuf:"bytes,3,opt,name=confleunceUsername,proto3" json:"confleunceUsername,omitempty"`
	ConfluenceToken    string `protobuf:"bytes,4,opt,name=confluenceToken,proto3" json:"confluenceToken,omitempty"`
}

func (x *PageReq) Reset() {
	*x = PageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_confluence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageReq) ProtoMessage() {}

func (x *PageReq) ProtoReflect() protoreflect.Message {
	mi := &file_confluence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageReq.ProtoReflect.Descriptor instead.
func (*PageReq) Descriptor() ([]byte, []int) {
	return file_confluence_proto_rawDescGZIP(), []int{0}
}

func (x *PageReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PageReq) GetConfluenceInstance() string {
	if x != nil {
		return x.ConfluenceInstance
	}
	return ""
}

func (x *PageReq) GetConfleunceUsername() string {
	if x != nil {
		return x.ConfleunceUsername
	}
	return ""
}

func (x *PageReq) GetConfluenceToken() string {
	if x != nil {
		return x.ConfluenceToken
	}
	return ""
}

type PageCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HtmlContent        string `protobuf:"bytes,1,opt,name=htmlContent,proto3" json:"htmlContent,omitempty"`
	Title              string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SpaceId            string `protobuf:"bytes,3,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	ParentPageId       string `protobuf:"bytes,4,opt,name=parentPageId,proto3" json:"parentPageId,omitempty"`
	ConfluenceInstance string `protobuf:"bytes,5,opt,name=confluenceInstance,proto3" json:"confluenceInstance,omitempty"`
	ConfleunceUsername string `protobuf:"bytes,6,opt,name=confleunceUsername,proto3" json:"confleunceUsername,omitempty"`
	ConfluenceToken    string `protobuf:"bytes,7,opt,name=confluenceToken,proto3" json:"confluenceToken,omitempty"`
}

func (x *PageCreateReq) Reset() {
	*x = PageCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_confluence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageCreateReq) ProtoMessage() {}

func (x *PageCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_confluence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageCreateReq.ProtoReflect.Descriptor instead.
func (*PageCreateReq) Descriptor() ([]byte, []int) {
	return file_confluence_proto_rawDescGZIP(), []int{1}
}

func (x *PageCreateReq) GetHtmlContent() string {
	if x != nil {
		return x.HtmlContent
	}
	return ""
}

func (x *PageCreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PageCreateReq) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PageCreateReq) GetParentPageId() string {
	if x != nil {
		return x.ParentPageId
	}
	return ""
}

func (x *PageCreateReq) GetConfluenceInstance() string {
	if x != nil {
		return x.ConfluenceInstance
	}
	return ""
}

func (x *PageCreateReq) GetConfleunceUsername() string {
	if x != nil {
		return x.ConfleunceUsername
	}
	return ""
}

func (x *PageCreateReq) GetConfluenceToken() string {
	if x != nil {
		return x.ConfluenceToken
	}
	return ""
}

type PageRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Html  string `protobuf:"bytes,3,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *PageRes) Reset() {
	*x = PageRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_confluence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRes) ProtoMessage() {}

func (x *PageRes) ProtoReflect() protoreflect.Message {
	mi := &file_confluence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRes.ProtoReflect.Descriptor instead.
func (*PageRes) Descriptor() ([]byte, []int) {
	return file_confluence_proto_rawDescGZIP(), []int{2}
}

func (x *PageRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PageRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PageRes) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

var File_confluence_proto protoreflect.FileDescriptor

var file_confluence_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x65, 0x75, 0x6e, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x65, 0x75, 0x6e, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x74,
	0x6d, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x68, 0x74, 0x6d, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x65, 0x75, 0x6e, 0x63, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x66, 0x6c, 0x65, 0x75, 0x6e, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x6c,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x74, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x32,
	0x57, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x08, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x08, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x64, 0x61, 0x6e, 0x74, 0x77, 0x61, 0x6e,
	0x6b, 0x68, 0x61, 0x64, 0x65, 0x2f, 0x74, 0x73, 0x72, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_confluence_proto_rawDescOnce sync.Once
	file_confluence_proto_rawDescData = file_confluence_proto_rawDesc
)

func file_confluence_proto_rawDescGZIP() []byte {
	file_confluence_proto_rawDescOnce.Do(func() {
		file_confluence_proto_rawDescData = protoimpl.X.CompressGZIP(file_confluence_proto_rawDescData)
	})
	return file_confluence_proto_rawDescData
}

var file_confluence_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_confluence_proto_goTypes = []interface{}{
	(*PageReq)(nil),       // 0: PageReq
	(*PageCreateReq)(nil), // 1: PageCreateReq
	(*PageRes)(nil),       // 2: PageRes
}
var file_confluence_proto_depIdxs = []int32{
	0, // 0: Confluence.GetPage:input_type -> PageReq
	1, // 1: Confluence.CreatePage:input_type -> PageCreateReq
	2, // 2: Confluence.GetPage:output_type -> PageRes
	2, // 3: Confluence.CreatePage:output_type -> PageRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_confluence_proto_init() }
func file_confluence_proto_init() {
	if File_confluence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_confluence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageReq); i {
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
		file_confluence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageCreateReq); i {
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
		file_confluence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRes); i {
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
			RawDescriptor: file_confluence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_confluence_proto_goTypes,
		DependencyIndexes: file_confluence_proto_depIdxs,
		MessageInfos:      file_confluence_proto_msgTypes,
	}.Build()
	File_confluence_proto = out.File
	file_confluence_proto_rawDesc = nil
	file_confluence_proto_goTypes = nil
	file_confluence_proto_depIdxs = nil
}
