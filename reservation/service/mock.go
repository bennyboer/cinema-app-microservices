package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
)

// Reservation service for tests.
type MockReservationService struct{}

func (s *MockReservationService) Reserve(ctx context.Context, in *proto.ReservationRequest, opts ...client.CallOption) (*proto.ReservationResponse, error) {
	return &proto.ReservationResponse{
		Available: true,
		CreatedId: 1,
	}, nil
}

func (s *MockReservationService) AcceptReservation(ctx context.Context, in *proto.AcceptReservationRequest, opts ...client.CallOption) (*proto.AcceptReservationResponse, error) {
	return &proto.AcceptReservationResponse{}, nil
}

func (s *MockReservationService) Cancel(ctx context.Context, in *proto.CancelReservationRequest, opts ...client.CallOption) (*proto.CancelReservationResponse, error) {
	return &proto.CancelReservationResponse{}, nil
}

func (s *MockReservationService) ReadAll(ctx context.Context, in *proto.ReadAllRequest, opts ...client.CallOption) (*proto.ReadAllResponse, error) {
	return &proto.ReadAllResponse{
		Ids: []int64{1},
		Dates: []*proto.ReservationData{
			{UserId: 1, PresentationId: 1, Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			}},
		},
	}, nil
}
