// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "User.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error) {
	req := c.c.NewRequest(c.name, "User.ReadAll", in)
	out := new(ReadAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "User.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "User.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error) {
	req := c.c.NewRequest(c.name, "User.Clear", in)
	out := new(ClearResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	ReadAll(context.Context, *ReadAllRequest, *ReadAllResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Clear(context.Context, *ClearRequest, *ClearResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.UserHandler.Read(ctx, in, out)
}

func (h *userHandler) ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error {
	return h.UserHandler.ReadAll(ctx, in, out)
}

func (h *userHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UserHandler.Update(ctx, in, out)
}

func (h *userHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UserHandler.Delete(ctx, in, out)
}

func (h *userHandler) Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error {
	return h.UserHandler.Clear(ctx, in, out)
}
