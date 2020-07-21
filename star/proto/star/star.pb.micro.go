// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: star/proto/star/star.proto

package go_micro_service_star

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Star service

type StarService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type starService struct {
	c    client.Client
	name string
}

func NewStarService(name string, c client.Client) StarService {
	return &starService{
		c:    c,
		name: name,
	}
}

func (c *starService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Star.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Star service

type StarHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterStarHandler(s server.Server, hdlr StarHandler, opts ...server.HandlerOption) error {
	type star interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Star struct {
		star
	}
	h := &starHandler{hdlr}
	return s.Handle(s.NewHandler(&Star{h}, opts...))
}

type starHandler struct {
	StarHandler
}

func (h *starHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.StarHandler.Call(ctx, in, out)
}
