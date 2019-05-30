package service

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	ps "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/service"
	"testing"
)

func getHandler() *MovieServiceHandler {
	return NewMovieServiceHandler(
		&MovieServiceDependencies{
			PresentationService: func() presentation.PresentationService {
				return &ps.MockPresentationService{}
			},
		},
	)
}

func TestMovieService_Create(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Inception",
		},
	}, rsp)

	if err != nil {
		t.Fatalf("create request returned unexpected error. Error:\n%s", err.Error())
	}

	if rsp.CreatedId < 0 {
		t.Errorf("expected created movie id to be non-negative")
	}
}

func TestMovieService_Create_EmptyTitle(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "",
		},
	}, rsp)

	if err == nil {
		t.Fatalf("expected create request to be unsuccessful")
	}
}

func TestMovieService_Read(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Inception",
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

	if rsp.Data.Title != "Inception" {
		t.Errorf("expected name of the read movie to be '%s' and not '%s'\n", "Inception", rsp.Data.Title)
	}
}

func TestMovieService_Read_Unsuccessful(t *testing.T) {
	handler := getHandler()

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: 5,
	}, rsp)

	if err == nil {
		t.Fatalf("expected read request to be unsuccessful")
	}
}

func TestMovieService_ReadAll(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Inception",
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Lord of the Rings",
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Titanic",
		},
	}, &proto.CreateResponse{})

	rsp := &proto.ReadAllResponse{}

	err := handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, rsp)
	if err != nil {
		t.Fatalf("read request returned unexpected error. Error:\n%s", err.Error())
	}

	if len(rsp.Ids) != 3 {
		t.Errorf("expected to find all 3 movie ids, got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 3 {
		t.Errorf("expected to find all 3 movie dates, got %d", len(rsp.Ids))
	}
}

func TestMovieService_Update(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Inception",
		},
	}, createRsp)

	id := createRsp.CreatedId

	err := handler.Update(context.TODO(), &proto.UpdateRequest{
		Id: id,
		Data: &proto.MovieData{
			Title: "Lord of the Rings",
		},
	}, &proto.UpdateResponse{})
	if err != nil {
		t.Fatalf("update request returned unexpected error. Error:\n%s", err.Error())
	}

	readRsp := &proto.ReadResponse{}
	_ = handler.Read(context.TODO(), &proto.ReadRequest{
		Id: id,
	}, readRsp)

	if readRsp.Data.Title != "Lord of the Rings" {
		t.Errorf("expected update to actually change the title")
	}
}

func TestMovieService_Update_Unsuccessful(t *testing.T) {
	handler := getHandler()

	err := handler.Update(context.TODO(), &proto.UpdateRequest{
		Id: 5,
		Data: &proto.MovieData{
			Title: "Another Title",
		},
	}, &proto.UpdateResponse{})

	if err == nil {
		t.Fatalf("expected update request to be unsuccessful")
	}
}

func TestMovieService_Delete(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.MovieData{
			Title: "Inception",
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
		t.Errorf("expected delete to actually delete the movie")
	}
}

func TestMovieService_Delete_Unsuccessful(t *testing.T) {
	handler := getHandler()

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: 6,
	}, &proto.DeleteResponse{})

	if err == nil {
		t.Fatalf("expected delete request to be unsuccessful")
	}
}
