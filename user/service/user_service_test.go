package service

import (
	"context"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	rs "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/service"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"testing"
)

func getHandler() *UserServiceHandler {
	return NewUserServiceHandler(
		&UserServiceDependencies{
			ReservationService: func() reservation.ReservationService {
				return &rs.MockReservationService{}
			},
		},
	)
}

func TestUserService_Create(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Max Mustermann",
		},
	}, rsp)

	if err != nil {
		t.Fatalf("create request returned unexpected error. Error:\n%s", err.Error())
	}

	if rsp.CreatedId < 0 {
		t.Errorf("expected created user id to be non-negative")
	}
}

func TestUserService_Create_EmptyName(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "",
		},
	}, rsp)

	if err == nil {
		t.Fatalf("expected create request to be unsuccessful")
	}
}

func TestUserService_Read(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Max Mustermann",
		},
	}, createRsp)

	id := createRsp.CreatedId

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: id,
	}, rsp)

	if err != nil {
		t.Fatalf("read request returned unexpected error. Error:\n%s", err.Error())
	}

	if rsp.Id != id {
		t.Errorf("expected id %d, got %d\n", id, rsp.Id)
	}

	if rsp.Data.Name != "Max Mustermann" {
		t.Errorf("expected name of the read user to be '%s' and not '%s'\n", "Max Mustermann", rsp.Data.Name)
	}
}

func TestUserService_Read_Unsuccessful(t *testing.T) {
	handler := getHandler()

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: 5,
	}, rsp)

	if err == nil {
		t.Fatalf("expected read request to be unsuccessful")
	}
}

func TestUserService_ReadAll(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Max Mustermann",
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Another Name",
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Hello World",
		},
	}, &proto.CreateResponse{})

	rsp := &proto.ReadAllResponse{}

	err := handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, rsp)
	if err != nil {
		t.Fatalf("read request returned unexpected error. Error:\n%s", err.Error())
	}

	if len(rsp.Ids) != 3 {
		t.Errorf("expected to find all 3 user ids, got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 3 {
		t.Errorf("expected to find all 3 user dates, got %d", len(rsp.Ids))
	}
}

func TestUserService_Update(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Max Mustermann",
		},
	}, createRsp)

	id := createRsp.CreatedId

	err := handler.Update(context.TODO(), &proto.UpdateRequest{
		Id: id,
		Data: &proto.UserData{
			Name: "Another Name",
		},
	}, &proto.UpdateResponse{})
	if err != nil {
		t.Fatalf("update request returned unexpected error. Error:\n%s", err.Error())
	}

	readRsp := &proto.ReadResponse{}
	_ = handler.Read(context.TODO(), &proto.ReadRequest{
		Id: id,
	}, readRsp)

	if readRsp.Data.Name != "Another Name" {
		t.Errorf("expected update to actually change the name")
	}
}

func TestUserService_Update_Unsuccessful(t *testing.T) {
	handler := getHandler()

	err := handler.Update(context.TODO(), &proto.UpdateRequest{
		Id: 5,
		Data: &proto.UserData{
			Name: "Another Name",
		},
	}, &proto.UpdateResponse{})

	if err == nil {
		t.Fatalf("expected update request to be unsuccessful")
	}
}

func TestUserService_Delete(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.UserData{
			Name: "Max Mustermann",
		},
	}, createRsp)

	id := createRsp.CreatedId

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: id,
	}, &proto.DeleteResponse{})

	if err != nil {
		t.Fatalf("delete request returned unexpected error. Error:\n%s", err.Error())
	}

	err = handler.Read(context.TODO(), &proto.ReadRequest{
		Id: id,
	}, &proto.ReadResponse{})

	if err == nil {
		t.Errorf("expected delete to actually delete the user")
	}
}

func TestUserService_Delete_Unsuccessful(t *testing.T) {
	handler := getHandler()

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: 6,
	}, &proto.DeleteResponse{})

	if err == nil {
		t.Fatalf("expected delete request to be unsuccessful")
	}
}
