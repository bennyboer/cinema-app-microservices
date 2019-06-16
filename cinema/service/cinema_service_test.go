package service

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	rs "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/service"
	"testing"
)

func getHandler() *CinemaServiceHandler {
	return NewCinemaServiceHandler(
		&CinemaServiceDependencies{
			PresentationService: func() presentation.PresentationService {
				return &rs.MockPresentationService{}
			},
		},
	)
}

func TestCinemaService_Create(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test",
		Row:   2,
		Seats: 2,
	}, rsp)

	if err != nil {
		t.Fatalf("create request returned unexpected error. Error:\n%s", err.Error())
	}

	if rsp.Data.Id < 0 {
		t.Errorf("expected created cinema id to be non-negative")
	}
}

func TestCinemaService_Create_EmptyName(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Name: "",
	}, rsp)

	if err == nil {
		t.Fatalf("expected search request to be unsuccessful")
	}
}

func TestCinemaService_Read(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name: "Test",
	}, createRsp)

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: createRsp.Data.Id,
	}, rsp)

	if err != nil {
		t.Fatalf("expected search request to be successful")
	}

	if createRsp.Data.Id != rsp.Data.Id {
		t.Fatalf("expected ids to be the same")
	}

	if !rsp.Success {
		t.Fatalf("expected search to be successful")
	}
}
