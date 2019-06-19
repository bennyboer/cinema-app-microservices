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

func TestCinemaService_Delete(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name: "Test",
	}, createRsp)

	rsp := &proto.DeleteResponse{}
	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: createRsp.Data.Id,
	}, rsp)

	if err != nil {
		t.Fatalf("expected search request to be successful")
	}

	if !rsp.Success {
		t.Fatalf("expected to be successful")
	}
}

func TestCinemaService_Occupy(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test",
		Row:   4,
		Seats: 4,
	}, createRsp)

	seats := []*proto.SeatData{
		{Row: 1, Seat: 1},
		{Row: 2, Seat: 2},
		{Row: 3, Seat: 3},
	}

	occupyRsp := &proto.OccupiedResponse{}
	err := handler.Occupy(context.TODO(), &proto.OccupiedRequest{
		Seats: seats,
		Id:    createRsp.Data.Id,
	}, occupyRsp)
	if err != nil {
		t.Fatalf("expected no error")
	}

	availableRsp := &proto.AvailableResponse{}
	err = handler.AreAvailable(context.TODO(), &proto.AvailableRequest{
		Id:    createRsp.Data.Id,
		Seats: seats,
	}, availableRsp)
	if err != nil {
		t.Fatalf("expected no error")
	}

	if availableRsp.Available {
		t.Errorf("expected seats to be unavailable")
	}
}

func TestCinemaService_Free(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test",
		Row:   4,
		Seats: 4,
	}, createRsp)

	seats := []*proto.SeatData{
		{Row: 1, Seat: 1},
		{Row: 2, Seat: 2},
		{Row: 3, Seat: 3},
	}

	occupyRsp := &proto.OccupiedResponse{}
	err := handler.Occupy(context.TODO(), &proto.OccupiedRequest{
		Seats: seats,
		Id:    createRsp.Data.Id,
	}, occupyRsp)
	if err != nil {
		t.Fatalf("expected no error")
	}

	err = handler.Free(context.TODO(), &proto.OccupiedRequest{
		Id:    createRsp.Data.Id,
		Seats: seats,
	}, &proto.OccupiedResponse{})
	if err != nil {
		t.Fatalf("expected no error")
	}

	availableRsp := &proto.AvailableResponse{}
	err = handler.AreAvailable(context.TODO(), &proto.AvailableRequest{
		Id:    createRsp.Data.Id,
		Seats: seats,
	}, availableRsp)
	if err != nil {
		t.Fatalf("expected no error")
	}

	if !availableRsp.Available {
		t.Errorf("expected seats to be available")
	}
}

func TestCinemaService_List(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 1",
		Row:   4,
		Seats: 4,
	}, createRsp)

	createRsp2 := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 2",
		Row:   4,
		Seats: 4,
	}, createRsp2)

	createRsp3 := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 3",
		Row:   4,
		Seats: 4,
	}, createRsp3)

	listRsp := &proto.ListResponse{}
	err := handler.List(context.TODO(), &proto.ListRequest{}, listRsp)
	if err != nil {
		t.Fatalf("expected no error")
	}

	if len(listRsp.Data) != 3 {
		t.Errorf("expected to list 3 cinemas; got %d", len(listRsp.Data))
	}
}

func TestCinemaService_Clear(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 1",
		Row:   4,
		Seats: 4,
	}, createRsp)

	createRsp2 := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 2",
		Row:   4,
		Seats: 4,
	}, createRsp2)

	createRsp3 := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Name:  "Test 3",
		Row:   4,
		Seats: 4,
	}, createRsp3)

	err := handler.Clear(context.TODO(), &proto.ClearRequest{}, &proto.ClearResponse{})
	if err != nil {
		t.Fatalf("expected no error; got %s", err.Error())
	}

	readRsp := &proto.ListResponse{}
	err = handler.List(context.TODO(), &proto.ListRequest{}, readRsp)
	if err != nil {
		t.Fatalf("expected no error; got %s", err.Error())
	}

	if len(readRsp.Data) != 0 {
		t.Errorf("expected service to have no more data; got %d dates", len(readRsp.Data))
	}
}
