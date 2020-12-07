// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/post/proto/post/post_service.proto

package post

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/ygpark2/mboard/service/post/proto/entities"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PostService service

func NewPostServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PostService service

type PostService interface {
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type postService struct {
	c    client.Client
	name string
}

func NewPostService(name string, c client.Client) PostService {
	return &postService{
		c:    c,
		name: name,
	}
}

func (c *postService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.Exist", in)
	out := new(ExistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "PostService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostService service

type PostServiceHandler interface {
	Exist(context.Context, *ExistRequest, *ExistResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterPostServiceHandler(s server.Server, hdlr PostServiceHandler, opts ...server.HandlerOption) error {
	type postService interface {
		Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type PostService struct {
		postService
	}
	h := &postServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PostService{h}, opts...))
}

type postServiceHandler struct {
	PostServiceHandler
}

func (h *postServiceHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error {
	return h.PostServiceHandler.Exist(ctx, in, out)
}

func (h *postServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.PostServiceHandler.List(ctx, in, out)
}

func (h *postServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.PostServiceHandler.Get(ctx, in, out)
}

func (h *postServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.PostServiceHandler.Create(ctx, in, out)
}

func (h *postServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.PostServiceHandler.Update(ctx, in, out)
}

func (h *postServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.PostServiceHandler.Delete(ctx, in, out)
}
