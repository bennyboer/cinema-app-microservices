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

// Client API for Presentation service

type PresentationService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	FindForCinema(ctx context.Context, in *FindForCinemaRequest, opts ...client.CallOption) (*FindForCinemaResponse, error)
	FindForMovie(ctx context.Context, in *FindForMovieRequest, opts ...client.CallOption) (*FindForMovieResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	DeleteForCinemas(ctx context.Context, in *DeleteForCinemasRequest, opts ...client.CallOption) (*DeleteForCinemasResponse, error)
	DeleteForMovies(ctx context.Context, in *DeleteForMoviesRequest, opts ...client.CallOption) (*DeleteForMoviesResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error)
}

type presentationService struct {
	c    client.Client
	name string
}

func NewPresentationService(name string, c client.Client) PresentationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &presentationService{
		c:    c,
		name: name,
	}
}

func (c *presentationService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) FindForCinema(ctx context.Context, in *FindForCinemaRequest, opts ...client.CallOption) (*FindForCinemaResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.FindForCinema", in)
	out := new(FindForCinemaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) FindForMovie(ctx context.Context, in *FindForMovieRequest, opts ...client.CallOption) (*FindForMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.FindForMovie", in)
	out := new(FindForMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.ReadAll", in)
	out := new(ReadAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) DeleteForCinemas(ctx context.Context, in *DeleteForCinemasRequest, opts ...client.CallOption) (*DeleteForCinemasResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.DeleteForCinemas", in)
	out := new(DeleteForCinemasResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) DeleteForMovies(ctx context.Context, in *DeleteForMoviesRequest, opts ...client.CallOption) (*DeleteForMoviesResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.DeleteForMovies", in)
	out := new(DeleteForMoviesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presentationService) Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error) {
	req := c.c.NewRequest(c.name, "Presentation.Clear", in)
	out := new(ClearResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Presentation service

type PresentationHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	FindForCinema(context.Context, *FindForCinemaRequest, *FindForCinemaResponse) error
	FindForMovie(context.Context, *FindForMovieRequest, *FindForMovieResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	ReadAll(context.Context, *ReadAllRequest, *ReadAllResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	DeleteForCinemas(context.Context, *DeleteForCinemasRequest, *DeleteForCinemasResponse) error
	DeleteForMovies(context.Context, *DeleteForMoviesRequest, *DeleteForMoviesResponse) error
	Clear(context.Context, *ClearRequest, *ClearResponse) error
}

func RegisterPresentationHandler(s server.Server, hdlr PresentationHandler, opts ...server.HandlerOption) error {
	type presentation interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		FindForCinema(ctx context.Context, in *FindForCinemaRequest, out *FindForCinemaResponse) error
		FindForMovie(ctx context.Context, in *FindForMovieRequest, out *FindForMovieResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		DeleteForCinemas(ctx context.Context, in *DeleteForCinemasRequest, out *DeleteForCinemasResponse) error
		DeleteForMovies(ctx context.Context, in *DeleteForMoviesRequest, out *DeleteForMoviesResponse) error
		Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error
	}
	type Presentation struct {
		presentation
	}
	h := &presentationHandler{hdlr}
	return s.Handle(s.NewHandler(&Presentation{h}, opts...))
}

type presentationHandler struct {
	PresentationHandler
}

func (h *presentationHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.PresentationHandler.Create(ctx, in, out)
}

func (h *presentationHandler) FindForCinema(ctx context.Context, in *FindForCinemaRequest, out *FindForCinemaResponse) error {
	return h.PresentationHandler.FindForCinema(ctx, in, out)
}

func (h *presentationHandler) FindForMovie(ctx context.Context, in *FindForMovieRequest, out *FindForMovieResponse) error {
	return h.PresentationHandler.FindForMovie(ctx, in, out)
}

func (h *presentationHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.PresentationHandler.Read(ctx, in, out)
}

func (h *presentationHandler) ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error {
	return h.PresentationHandler.ReadAll(ctx, in, out)
}

func (h *presentationHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.PresentationHandler.Delete(ctx, in, out)
}

func (h *presentationHandler) DeleteForCinemas(ctx context.Context, in *DeleteForCinemasRequest, out *DeleteForCinemasResponse) error {
	return h.PresentationHandler.DeleteForCinemas(ctx, in, out)
}

func (h *presentationHandler) DeleteForMovies(ctx context.Context, in *DeleteForMoviesRequest, out *DeleteForMoviesResponse) error {
	return h.PresentationHandler.DeleteForMovies(ctx, in, out)
}

func (h *presentationHandler) Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error {
	return h.PresentationHandler.Clear(ctx, in, out)
}
