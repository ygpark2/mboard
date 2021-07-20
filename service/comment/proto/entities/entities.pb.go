// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: service/comment/proto/entities/entities.proto

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

// Comment Entity
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *types.UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt     *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	BoardId       *wrappers.StringValue `protobuf:"bytes,5,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	PostId        *wrappers.StringValue `protobuf:"bytes,6,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ContentId     *wrappers.StringValue `protobuf:"bytes,7,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	Userid        *wrappers.StringValue `protobuf:"bytes,8,opt,name=userid,proto3" json:"userid,omitempty"`
	Username      *wrappers.StringValue `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	Nickname      *wrappers.StringValue `protobuf:"bytes,10,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email         *wrappers.StringValue `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	Password      *wrappers.StringValue `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	Url           *wrappers.StringValue `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
	UseHtml       *wrappers.BoolValue   `protobuf:"bytes,14,opt,name=use_html,json=useHtml,proto3" json:"use_html,omitempty"`
	UseSecret     *wrappers.BoolValue   `protobuf:"bytes,15,opt,name=use_secret,json=useSecret,proto3" json:"use_secret,omitempty"`
	UpVoteCount   uint32                `protobuf:"varint,16,opt,name=up_vote_count,json=upVoteCount,proto3" json:"up_vote_count,omitempty"`
	DownVoteCount uint32                `protobuf:"varint,17,opt,name=down_vote_count,json=downVoteCount,proto3" json:"down_vote_count,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_comment_proto_entities_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_service_comment_proto_entities_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_service_comment_proto_entities_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Comment) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Comment) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Comment) GetBoardId() *wrappers.StringValue {
	if x != nil {
		return x.BoardId
	}
	return nil
}

func (x *Comment) GetPostId() *wrappers.StringValue {
	if x != nil {
		return x.PostId
	}
	return nil
}

func (x *Comment) GetContentId() *wrappers.StringValue {
	if x != nil {
		return x.ContentId
	}
	return nil
}

func (x *Comment) GetUserid() *wrappers.StringValue {
	if x != nil {
		return x.Userid
	}
	return nil
}

func (x *Comment) GetUsername() *wrappers.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *Comment) GetNickname() *wrappers.StringValue {
	if x != nil {
		return x.Nickname
	}
	return nil
}

func (x *Comment) GetEmail() *wrappers.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *Comment) GetPassword() *wrappers.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *Comment) GetUrl() *wrappers.StringValue {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Comment) GetUseHtml() *wrappers.BoolValue {
	if x != nil {
		return x.UseHtml
	}
	return nil
}

func (x *Comment) GetUseSecret() *wrappers.BoolValue {
	if x != nil {
		return x.UseSecret
	}
	return nil
}

func (x *Comment) GetUpVoteCount() uint32 {
	if x != nil {
		return x.UpVoteCount
	}
	return 0
}

func (x *Comment) GetDownVoteCount() uint32 {
	if x != nil {
		return x.DownVoteCount
	}
	return 0
}

var File_service_comment_proto_entities_entities_proto protoreflect.FileDescriptor

var file_service_comment_proto_entities_entities_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x6d, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x92, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9,
	0x19, 0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5a,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1f,
	0xba, 0xb9, 0x19, 0x1b, 0x0a, 0x19, 0x52, 0x17, 0x69, 0x64, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f,
	0x0a, 0x0d, 0x40, 0x01, 0x52, 0x09, 0x69, 0x64, 0x78, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x40,
	0x01, 0x52, 0x08, 0x69, 0x64, 0x78, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x40, 0x01, 0x52,
	0x0b, 0x69, 0x64, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x18, 0x64, 0x40,
	0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a,
	0x05, 0x18, 0xc8, 0x01, 0x40, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x45, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x18, 0xc8, 0x01, 0x40, 0x01, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x18, 0xfa, 0x01, 0x40,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40,
	0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x74,
	0x6d, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x43, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40,
	0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x0d,
	0x75, 0x70, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x75,
	0x70, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x6f,
	0x77, 0x6e, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0d, 0x64,
	0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x09, 0xf8, 0x42,
	0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x5e, 0x0a, 0x20, 0x6d, 0x6b, 0x69, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x67, 0x70, 0x61, 0x72, 0x6b,
	0x32, 0x2f, 0x6d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_comment_proto_entities_entities_proto_rawDescOnce sync.Once
	file_service_comment_proto_entities_entities_proto_rawDescData = file_service_comment_proto_entities_entities_proto_rawDesc
)

func file_service_comment_proto_entities_entities_proto_rawDescGZIP() []byte {
	file_service_comment_proto_entities_entities_proto_rawDescOnce.Do(func() {
		file_service_comment_proto_entities_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_comment_proto_entities_entities_proto_rawDescData)
	})
	return file_service_comment_proto_entities_entities_proto_rawDescData
}

var file_service_comment_proto_entities_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_comment_proto_entities_entities_proto_goTypes = []interface{}{
	(*Comment)(nil),              // 0: mkit.service.comment.entities.v1.Comment
	(*types.UUID)(nil),           // 1: gorm.types.UUID
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrappers.BoolValue)(nil),   // 4: google.protobuf.BoolValue
}
var file_service_comment_proto_entities_entities_proto_depIdxs = []int32{
	1,  // 0: mkit.service.comment.entities.v1.Comment.id:type_name -> gorm.types.UUID
	2,  // 1: mkit.service.comment.entities.v1.Comment.created_at:type_name -> google.protobuf.Timestamp
	2,  // 2: mkit.service.comment.entities.v1.Comment.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 3: mkit.service.comment.entities.v1.Comment.deleted_at:type_name -> google.protobuf.Timestamp
	3,  // 4: mkit.service.comment.entities.v1.Comment.board_id:type_name -> google.protobuf.StringValue
	3,  // 5: mkit.service.comment.entities.v1.Comment.post_id:type_name -> google.protobuf.StringValue
	3,  // 6: mkit.service.comment.entities.v1.Comment.content_id:type_name -> google.protobuf.StringValue
	3,  // 7: mkit.service.comment.entities.v1.Comment.userid:type_name -> google.protobuf.StringValue
	3,  // 8: mkit.service.comment.entities.v1.Comment.username:type_name -> google.protobuf.StringValue
	3,  // 9: mkit.service.comment.entities.v1.Comment.nickname:type_name -> google.protobuf.StringValue
	3,  // 10: mkit.service.comment.entities.v1.Comment.email:type_name -> google.protobuf.StringValue
	3,  // 11: mkit.service.comment.entities.v1.Comment.password:type_name -> google.protobuf.StringValue
	3,  // 12: mkit.service.comment.entities.v1.Comment.url:type_name -> google.protobuf.StringValue
	4,  // 13: mkit.service.comment.entities.v1.Comment.use_html:type_name -> google.protobuf.BoolValue
	4,  // 14: mkit.service.comment.entities.v1.Comment.use_secret:type_name -> google.protobuf.BoolValue
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_service_comment_proto_entities_entities_proto_init() }
func file_service_comment_proto_entities_entities_proto_init() {
	if File_service_comment_proto_entities_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_comment_proto_entities_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_service_comment_proto_entities_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_comment_proto_entities_entities_proto_goTypes,
		DependencyIndexes: file_service_comment_proto_entities_entities_proto_depIdxs,
		MessageInfos:      file_service_comment_proto_entities_entities_proto_msgTypes,
	}.Build()
	File_service_comment_proto_entities_entities_proto = out.File
	file_service_comment_proto_entities_entities_proto_rawDesc = nil
	file_service_comment_proto_entities_entities_proto_goTypes = nil
	file_service_comment_proto_entities_entities_proto_depIdxs = nil
}
