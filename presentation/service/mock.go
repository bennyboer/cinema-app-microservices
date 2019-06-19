package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
)

// Presentation which can be injected during testing.
type MockPresentationService struct {
}

func (s *MockPresentationService) Create(ctx context.Context, in *proto.CreateRequest, opts ...client.CallOption) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{}, nil
}

func (s *MockPresentationService) FindForCinema(ctx context.Context, in *proto.FindForCinemaRequest, opts ...client.CallOption) (*proto.FindForCinemaResponse, error) {
	return &proto.FindForCinemaResponse{
		Dates: []*proto.PresentationData{
			{
				MovieId:  1,
				CinemaId: in.CinemaId,
			},
		},
	}, nil
}

func (s *MockPresentationService) FindForMovie(ctx context.Context, in *proto.FindForMovieRequest, opts ...client.CallOption) (*proto.FindForMovieResponse, error) {
	return &proto.FindForMovieResponse{
		Dates: []*proto.PresentationData{
			{
				MovieId:  in.MovieId,
				CinemaId: 1,
			},
		},
	}, nil
}

func (s *MockPresentationService) Read(ctx context.Context, in *proto.ReadRequest, opts ...client.CallOption) (*proto.ReadResponse, error) {
	return &proto.ReadResponse{
		Data: &proto.PresentationData{
			MovieId:  1,
			CinemaId: 1,
		},
	}, nil
}

func (s *MockPresentationService) ReadAll(ctx context.Context, in *proto.ReadAllRequest, opts ...client.CallOption) (*proto.ReadAllResponse, error) {
	return &proto.ReadAllResponse{
		Dates: []*proto.PresentationData{
			{
				MovieId:  1,
				CinemaId: 1,
			},
		},
	}, nil
}

func (s *MockPresentationService) Delete(ctx context.Context, in *proto.DeleteRequest, opts ...client.CallOption) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}

func (s *MockPresentationService) DeleteForCinemas(ctx context.Context, in *proto.DeleteForCinemasRequest, opts ...client.CallOption) (*proto.DeleteForCinemasResponse, error) {
	return &proto.DeleteForCinemasResponse{}, nil
}

func (s *MockPresentationService) DeleteForMovies(ctx context.Context, in *proto.DeleteForMoviesRequest, opts ...client.CallOption) (*proto.DeleteForMoviesResponse, error) {
	return &proto.DeleteForMoviesResponse{}, nil
}

func (s *MockPresentationService) Clear(ctx context.Context, in *proto.ClearRequest, opts ...client.CallOption) (*proto.ClearResponse, error) {
	return &proto.ClearResponse{}, nil
}

