// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/tag.proto

package tags

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the parent object
	ParentID string `protobuf:"bytes,1,opt,name=parentID,proto3" json:"parentID,omitempty"`
	// Type is useful for namespacing and listing across parents,
	// ie. list tags for posts, customers etc.
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Slug  string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Count int64  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *Tag) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Tag) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Tag) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Tag) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type IncreaseCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID string `protobuf:"bytes,1,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *IncreaseCountRequest) Reset() {
	*x = IncreaseCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseCountRequest) ProtoMessage() {}

func (x *IncreaseCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseCountRequest.ProtoReflect.Descriptor instead.
func (*IncreaseCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{1}
}

func (x *IncreaseCountRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *IncreaseCountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IncreaseCountRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type IncreaseCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncreaseCountResponse) Reset() {
	*x = IncreaseCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseCountResponse) ProtoMessage() {}

func (x *IncreaseCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseCountResponse.ProtoReflect.Descriptor instead.
func (*IncreaseCountResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{2}
}

type DecreaseCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID string `protobuf:"bytes,1,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DecreaseCountRequest) Reset() {
	*x = DecreaseCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseCountRequest) ProtoMessage() {}

func (x *DecreaseCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseCountRequest.ProtoReflect.Descriptor instead.
func (*DecreaseCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{3}
}

func (x *DecreaseCountRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *DecreaseCountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DecreaseCountRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DecreaseCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecreaseCountResponse) Reset() {
	*x = DecreaseCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseCountResponse) ProtoMessage() {}

func (x *DecreaseCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseCountResponse.ProtoReflect.Descriptor instead.
func (*DecreaseCountResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{4}
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID string `protobuf:"bytes,1,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *UpdateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{6}
}

// ListRequest: list either by parent id or type.
// Optionally filter by min or max count.
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID string `protobuf:"bytes,1,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	MinCount int64  `protobuf:"varint,3,opt,name=minCount,proto3" json:"minCount,omitempty"`
	MaxCount int64  `protobuf:"varint,4,opt,name=maxCount,proto3" json:"maxCount,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{7}
}

func (x *ListRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *ListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ListRequest) GetMinCount() int64 {
	if x != nil {
		return x.MinCount
	}
	return 0
}

func (x *ListRequest) GetMaxCount() int64 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{8}
}

func (x *ListResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_proto_tag_proto protoreflect.FileDescriptor

var file_proto_tag_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x75, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c,
	0x0a, 0x14, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x17, 0x0a, 0x15,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x74, 0x61, 0x67,
	0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0x86, 0x02, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x61, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74,
	0x61, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tag_proto_rawDescOnce sync.Once
	file_proto_tag_proto_rawDescData = file_proto_tag_proto_rawDesc
)

func file_proto_tag_proto_rawDescGZIP() []byte {
	file_proto_tag_proto_rawDescOnce.Do(func() {
		file_proto_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tag_proto_rawDescData)
	})
	return file_proto_tag_proto_rawDescData
}

var file_proto_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),                   // 0: tags.Tag
	(*IncreaseCountRequest)(nil),  // 1: tags.IncreaseCountRequest
	(*IncreaseCountResponse)(nil), // 2: tags.IncreaseCountResponse
	(*DecreaseCountRequest)(nil),  // 3: tags.DecreaseCountRequest
	(*DecreaseCountResponse)(nil), // 4: tags.DecreaseCountResponse
	(*UpdateRequest)(nil),         // 5: tags.UpdateRequest
	(*UpdateResponse)(nil),        // 6: tags.UpdateResponse
	(*ListRequest)(nil),           // 7: tags.ListRequest
	(*ListResponse)(nil),          // 8: tags.ListResponse
}
var file_proto_tag_proto_depIdxs = []int32{
	0, // 0: tags.ListResponse.tags:type_name -> tags.Tag
	1, // 1: tags.Tags.IncreaseCount:input_type -> tags.IncreaseCountRequest
	3, // 2: tags.Tags.DecreaseCount:input_type -> tags.DecreaseCountRequest
	7, // 3: tags.Tags.List:input_type -> tags.ListRequest
	5, // 4: tags.Tags.Update:input_type -> tags.UpdateRequest
	2, // 5: tags.Tags.IncreaseCount:output_type -> tags.IncreaseCountResponse
	4, // 6: tags.Tags.DecreaseCount:output_type -> tags.DecreaseCountResponse
	8, // 7: tags.Tags.List:output_type -> tags.ListResponse
	6, // 8: tags.Tags.Update:output_type -> tags.UpdateResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_tag_proto_init() }
func file_proto_tag_proto_init() {
	if File_proto_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_proto_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseCountRequest); i {
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
		file_proto_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseCountResponse); i {
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
		file_proto_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseCountRequest); i {
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
		file_proto_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseCountResponse); i {
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
		file_proto_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_proto_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_proto_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tag_proto_goTypes,
		DependencyIndexes: file_proto_tag_proto_depIdxs,
		MessageInfos:      file_proto_tag_proto_msgTypes,
	}.Build()
	File_proto_tag_proto = out.File
	file_proto_tag_proto_rawDesc = nil
	file_proto_tag_proto_goTypes = nil
	file_proto_tag_proto_depIdxs = nil
}
