package service

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
)

// User service used to test services properly.
type MockUserService struct{}

func (s *MockUserService) Create(ctx context.Context, in *proto.CreateRequest, opts ...client.CallOption) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{
		CreatedId: 1,
	}, nil
}

func (s *MockUserService) Read(ctx context.Context, in *proto.ReadRequest, opts ...client.CallOption) (*proto.ReadResponse, error) {
	return &proto.ReadResponse{
		Id: in.Id,
		Data: &proto.UserData{
			Name: "Test user",
		},
	}, nil
}

func (s *MockUserService) ReadAll(ctx context.Context, in *proto.ReadAllRequest, opts ...client.CallOption) (*proto.ReadAllResponse, error) {
	return &proto.ReadAllResponse{
		Ids: []int64{1, 2, 3},
		Dates: []*proto.UserData{
			{Name: "Test user 1"},
			{Name: "Test user 2"},
			{Name: "Test user 3"},
		},
	}, nil
}

func (s *MockUserService) Update(ctx context.Context, in *proto.UpdateRequest, opts ...client.CallOption) (*proto.UpdateResponse, error) {
	return &proto.UpdateResponse{}, nil
}

func (s *MockUserService) Delete(ctx context.Context, in *proto.DeleteRequest, opts ...client.CallOption) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}

func (s *MockUserService) Clear(ctx context.Context, in *proto.ClearRequest, opts ...client.CallOption) (*proto.ClearResponse, error) {
	return &proto.ClearResponse{}, nil
}
