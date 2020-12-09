// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/content/proto/content/content_service.proto

package contents

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for ContentService service

func NewContentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ContentService service

type ContentService interface {
	Save(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type contentService struct {
	c    client.Client
	name string
}

func NewContentService(name string, c client.Client) ContentService {
	return &contentService{
		c:    c,
		name: name,
	}
}

func (c *contentService) Save(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ContentService.Save", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentService service

type ContentServiceHandler interface {
	Save(context.Context, *Request, *Response) error
}

func RegisterContentServiceHandler(s server.Server, hdlr ContentServiceHandler, opts ...server.HandlerOption) error {
	type contentService interface {
		Save(ctx context.Context, in *Request, out *Response) error
	}
	type ContentService struct {
		contentService
	}
	h := &contentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ContentService{h}, opts...))
}

type contentServiceHandler struct {
	ContentServiceHandler
}

func (h *contentServiceHandler) Save(ctx context.Context, in *Request, out *Response) error {
	return h.ContentServiceHandler.Save(ctx, in, out)
}
