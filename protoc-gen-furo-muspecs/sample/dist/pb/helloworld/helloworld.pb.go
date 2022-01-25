// Code generated by furo-proto-gen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: helloworld/helloworld.proto

package helloworldpb

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

// Reply ...
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The response message containing the greetings.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// nested type is not imported at the moment
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetResults() []*SearchResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// Request ...
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request message containing the user's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Stuuuuufff ...
type Stuff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request message containing the user's name.
	Xxx []string `protobuf:"bytes,1,rep,name=xxx,proto3" json:"xxx,omitempty"`
	// Types that are assignable to TestOneof:
	//	*Stuff_Name
	//	*Stuff_SubMessage
	TestOneof isStuff_TestOneof `protobuf_oneof:"test_oneof"`
	// Types that are assignable to OgherOneof:
	//	*Stuff_Oname
	//	*Stuff_OsubMessage
	OgherOneof isStuff_OgherOneof `protobuf_oneof:"ogher_oneof"`
}

func (x *Stuff) Reset() {
	*x = Stuff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stuff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stuff) ProtoMessage() {}

func (x *Stuff) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stuff.ProtoReflect.Descriptor instead.
func (*Stuff) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *Stuff) GetXxx() []string {
	if x != nil {
		return x.Xxx
	}
	return nil
}

func (m *Stuff) GetTestOneof() isStuff_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *Stuff) GetName() string {
	if x, ok := x.GetTestOneof().(*Stuff_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Stuff) GetSubMessage() *Request {
	if x, ok := x.GetTestOneof().(*Stuff_SubMessage); ok {
		return x.SubMessage
	}
	return nil
}

func (m *Stuff) GetOgherOneof() isStuff_OgherOneof {
	if m != nil {
		return m.OgherOneof
	}
	return nil
}

func (x *Stuff) GetOname() string {
	if x, ok := x.GetOgherOneof().(*Stuff_Oname); ok {
		return x.Oname
	}
	return ""
}

func (x *Stuff) GetOsubMessage() *Reply {
	if x, ok := x.GetOgherOneof().(*Stuff_OsubMessage); ok {
		return x.OsubMessage
	}
	return nil
}

type isStuff_TestOneof interface {
	isStuff_TestOneof()
}

type Stuff_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"`
}

type Stuff_SubMessage struct {
	SubMessage *Request `protobuf:"bytes,9,opt,name=sub_message,json=subMessage,proto3,oneof"`
}

func (*Stuff_Name) isStuff_TestOneof() {}

func (*Stuff_SubMessage) isStuff_TestOneof() {}

type isStuff_OgherOneof interface {
	isStuff_OgherOneof()
}

type Stuff_Oname struct {
	// davor
	Oname string `protobuf:"bytes,2,opt,name=oname,proto3,oneof"` // Da geht noch was
}

type Stuff_OsubMessage struct {
	OsubMessage *Reply `protobuf:"bytes,24,opt,name=osub_message,json=osubMessage,proto3,oneof"`
}

func (*Stuff_Oname) isStuff_OgherOneof() {}

func (*Stuff_OsubMessage) isStuff_OgherOneof() {}

// subtype
type SearchResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string                           `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Subresult *SearchResponse_Result_SubResult `protobuf:"bytes,8,opt,name=subresult,proto3" json:"subresult,omitempty"`
	Title     string                           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets  []string                         `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *SearchResponse_Result) Reset() {
	*x = SearchResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_Result) ProtoMessage() {}

func (x *SearchResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_Result.ProtoReflect.Descriptor instead.
func (*SearchResponse_Result) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SearchResponse_Result) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SearchResponse_Result) GetSubresult() *SearchResponse_Result_SubResult {
	if x != nil {
		return x.Subresult
	}
	return nil
}

func (x *SearchResponse_Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchResponse_Result) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

// subtype
type SearchResponse_Result_SubResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SearchResponse_Result_SubResult) Reset() {
	*x = SearchResponse_Result_SubResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_Result_SubResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_Result_SubResult) ProtoMessage() {}

func (x *SearchResponse_Result_SubResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_Result_SubResult.ProtoReflect.Descriptor instead.
func (*SearchResponse_Result_SubResult) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *SearchResponse_Result_SubResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x21, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x86, 0x02, 0x0a,
	0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xb6, 0x01, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x49, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x1a, 0x1d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x10,
	0x0a, 0x03, 0x78, 0x78, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x78, 0x78, 0x78,
	0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x6f, 0x73, 0x75, 0x62, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48,
	0x01, 0x52, 0x0b, 0x6f, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x0d, 0x0a, 0x0b,
	0x6f, 0x67, 0x68, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x6f, 0x0a, 0x1f, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65,
	0x69, 0x74, 0x68, 0x2f, 0x70, 0x75, 0x72, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x3b,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_helloworld_proto_rawDescData = file_helloworld_helloworld_proto_rawDesc
)

func file_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_helloworld_proto_rawDescData)
	})
	return file_helloworld_helloworld_proto_rawDescData
}

var file_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_helloworld_helloworld_proto_goTypes = []interface{}{
	(*Reply)(nil),                           // 0: helloworld.Reply
	(*SearchResponse)(nil),                  // 1: helloworld.SearchResponse
	(*Request)(nil),                         // 2: helloworld.Request
	(*Stuff)(nil),                           // 3: helloworld.Stuff
	(*SearchResponse_Result)(nil),           // 4: helloworld.SearchResponse.Result
	(*SearchResponse_Result_SubResult)(nil), // 5: helloworld.SearchResponse.Result.SubResult
}
var file_helloworld_helloworld_proto_depIdxs = []int32{
	4, // 0: helloworld.SearchResponse.results:type_name -> helloworld.SearchResponse.Result
	2, // 1: helloworld.Stuff.sub_message:type_name -> helloworld.Request
	0, // 2: helloworld.Stuff.osub_message:type_name -> helloworld.Reply
	5, // 3: helloworld.SearchResponse.Result.subresult:type_name -> helloworld.SearchResponse.Result.SubResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_helloworld_helloworld_proto_init() }
func file_helloworld_helloworld_proto_init() {
	if File_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_helloworld_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stuff); i {
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
		file_helloworld_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse_Result); i {
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
		file_helloworld_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse_Result_SubResult); i {
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
	file_helloworld_helloworld_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Stuff_Name)(nil),
		(*Stuff_SubMessage)(nil),
		(*Stuff_Oname)(nil),
		(*Stuff_OsubMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_helloworld_proto = out.File
	file_helloworld_helloworld_proto_rawDesc = nil
	file_helloworld_helloworld_proto_goTypes = nil
	file_helloworld_helloworld_proto_depIdxs = nil
}
