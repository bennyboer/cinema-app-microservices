package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
)

type MockCinemaService struct{}

func (s *MockCinemaService) Create(ctx context.Context, in *proto.CreateRequest, opts ...client.CallOption) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{
		Data: &proto.CinemaData{
			Id:   1,
			Name: in.Name,
		},
	}, nil
}

func (s *MockCinemaService) Delete(ctx context.Context, in *proto.DeleteRequest, opts ...client.CallOption) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}
