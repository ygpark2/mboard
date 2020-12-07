// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: service/post/proto/entities/entities.proto

package entities

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	types "github.com/infobloxopen/protoc-gen-gorm/types"
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

// Post Entity
type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *types.UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt      *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt      *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	BoardId        *wrappers.StringValue `protobuf:"bytes,5,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Title          *wrappers.StringValue `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Category       *wrappers.StringValue `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	Password       *wrappers.StringValue `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Content        string                `protobuf:"bytes,9,opt,name=content,proto3" json:"content,omitempty"`
	Email          string                `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Writer         string                `protobuf:"bytes,11,opt,name=writer,proto3" json:"writer,omitempty"`
	CommentCounts  uint32                `protobuf:"varint,12,opt,name=comment_counts,json=commentCounts,proto3" json:"comment_counts,omitempty"`
	ReadCounts     uint32                `protobuf:"varint,13,opt,name=read_counts,json=readCounts,proto3" json:"read_counts,omitempty"`
	UpVoteCounts   uint32                `protobuf:"varint,14,opt,name=up_vote_counts,json=upVoteCounts,proto3" json:"up_vote_counts,omitempty"`
	DownVoteCounts uint32                `protobuf:"varint,15,opt,name=down_vote_counts,json=downVoteCounts,proto3" json:"down_vote_counts,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_post_proto_entities_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_service_post_proto_entities_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_service_post_proto_entities_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Post) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Post) GetBoardId() *wrappers.StringValue {
	if x != nil {
		return x.BoardId
	}
	return nil
}

func (x *Post) GetTitle() *wrappers.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Post) GetCategory() *wrappers.StringValue {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Post) GetPassword() *wrappers.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Post) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *Post) GetCommentCounts() uint32 {
	if x != nil {
		return x.CommentCounts
	}
	return 0
}

func (x *Post) GetReadCounts() uint32 {
	if x != nil {
		return x.ReadCounts
	}
	return 0
}

func (x *Post) GetUpVoteCounts() uint32 {
	if x != nil {
		return x.UpVoteCounts
	}
	return 0
}

func (x *Post) GetDownVoteCounts() uint32 {
	if x != nil {
		return x.DownVoteCounts
	}
	return 0
}

var File_service_post_proto_entities_entities_proto protoreflect.FileDescriptor

var file_service_post_proto_entities_entities_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6d, 0x6b,
	0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x06, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x57, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x18, 0x0a, 0x16, 0x52, 0x14, 0x69,
	0x64, 0x78, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x52,
	0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x19,
	0xba, 0xb9, 0x19, 0x15, 0x0a, 0x13, 0x40, 0x01, 0x52, 0x0f, 0x69, 0x64, 0x78, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x42, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x00, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x06,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x2d,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0e, 0x75, 0x70, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x75, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3a, 0x09, 0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08,
	0x01, 0x42, 0x58, 0x0a, 0x1d, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x67, 0x70, 0x61, 0x72, 0x6b, 0x32, 0x2f, 0x6d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_post_proto_entities_entities_proto_rawDescOnce sync.Once
	file_service_post_proto_entities_entities_proto_rawDescData = file_service_post_proto_entities_entities_proto_rawDesc
)

func file_service_post_proto_entities_entities_proto_rawDescGZIP() []byte {
	file_service_post_proto_entities_entities_proto_rawDescOnce.Do(func() {
		file_service_post_proto_entities_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_post_proto_entities_entities_proto_rawDescData)
	})
	return file_service_post_proto_entities_entities_proto_rawDescData
}

var file_service_post_proto_entities_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_post_proto_entities_entities_proto_goTypes = []interface{}{
	(*Post)(nil),                 // 0: mkit.service.post.entities.v1.Post
	(*types.UUID)(nil),           // 1: gorm.types.UUID
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_service_post_proto_entities_entities_proto_depIdxs = []int32{
	1, // 0: mkit.service.post.entities.v1.Post.id:type_name -> gorm.types.UUID
	2, // 1: mkit.service.post.entities.v1.Post.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: mkit.service.post.entities.v1.Post.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: mkit.service.post.entities.v1.Post.deleted_at:type_name -> google.protobuf.Timestamp
	3, // 4: mkit.service.post.entities.v1.Post.board_id:type_name -> google.protobuf.StringValue
	3, // 5: mkit.service.post.entities.v1.Post.title:type_name -> google.protobuf.StringValue
	3, // 6: mkit.service.post.entities.v1.Post.category:type_name -> google.protobuf.StringValue
	3, // 7: mkit.service.post.entities.v1.Post.password:type_name -> google.protobuf.StringValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_service_post_proto_entities_entities_proto_init() }
func file_service_post_proto_entities_entities_proto_init() {
	if File_service_post_proto_entities_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_post_proto_entities_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
			RawDescriptor: file_service_post_proto_entities_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_post_proto_entities_entities_proto_goTypes,
		DependencyIndexes: file_service_post_proto_entities_entities_proto_depIdxs,
		MessageInfos:      file_service_post_proto_entities_entities_proto_msgTypes,
	}.Build()
	File_service_post_proto_entities_entities_proto = out.File
	file_service_post_proto_entities_entities_proto_rawDesc = nil
	file_service_post_proto_entities_entities_proto_goTypes = nil
	file_service_post_proto_entities_entities_proto_depIdxs = nil
}
