package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
)

type MockCinemaService struct{}

func (s *MockCinemaService) Read(ctx context.Context, in *proto.ReadRequest, opts ...client.CallOption) (*proto.ReadResponse, error) {
	return &proto.ReadResponse{
		Data: &proto.CinemaData{
			Id:    in.Id,
			Seats: nil,
			Name:  "test",
		},
	}, nil
}

func (s *MockCinemaService) List(ctx context.Context, in *proto.ListRequest, opts ...client.CallOption) (*proto.ListResponse, error) {
	return &proto.ListResponse{
		Data: []*proto.CinemaData{
			{
				Name:  "test",
				Seats: nil,
				Id:    1,
			},
		},
	}, nil
}

func (s *MockCinemaService) Occupy(ctx context.Context, in *proto.OccupiedRequest, opts ...client.CallOption) (*proto.OccupiedResponse, error) {
	return &proto.OccupiedResponse{
		Seats: nil,
	}, nil
}

func (s *MockCinemaService) Free(ctx context.Context, in *proto.OccupiedRequest, opts ...client.CallOption) (*proto.OccupiedResponse, error) {
	return &proto.OccupiedResponse{
		Seats: nil,
	}, nil
}

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

func (s *MockCinemaService) AreAvailable(ctx context.Context, in *proto.AvailableRequest, opts ...client.CallOption) (*proto.AvailableResponse, error) {
	return &proto.AvailableResponse{
		Available: true,
	}, nil
}

func (s *MockCinemaService) Clear(ctx context.Context, in *proto.ClearRequest, opts ...client.CallOption) (*proto.ClearResponse, error) {
	return &proto.ClearResponse{}, nil
}
