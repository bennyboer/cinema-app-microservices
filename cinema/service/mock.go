package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
)

type MockCinemaService struct{}

func (s *MockCinemaService) Create(ctx context.Context, in *proto.CreateRequest, opts ...client.CallOption) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{
		Id: 1,
	}, nil
}

func (s *MockCinemaService) Search(ctx context.Context, in *proto.SearchRequest, opts ...client.CallOption) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{
		Success: true,
		Data: &proto.CinemaData{
			Name: in.Name,
			Id:   1,
		},
	}, nil
}

func (s *MockCinemaService) Delete(ctx context.Context, in *proto.DeleteRequest, opts ...client.CallOption) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}
