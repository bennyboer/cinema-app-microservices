package service

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
)

type MovieServiceHandler struct{}

func NewMovieServiceHandler() *MovieServiceHandler {
	return &MovieServiceHandler{}
}

func (h *MovieServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	panic("implement me")
}

func (h *MovieServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	panic("implement me")
}

func (h *MovieServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	panic("implement me")
}

func (h *MovieServiceHandler) Update(context context.Context, request *proto.UpdateRequest, response *proto.UpdateResponse) error {
	panic("implement me")
}

func (h *MovieServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	panic("implement me")
}
