package service

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	rs "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/service"
	"testing"
)

func getHandler() *PresentationServiceHandler {
	return NewPresentationServiceHandler(
		&PresentationServiceDependencies{
			ReservationService: func() reservation.ReservationService {
				return &rs.MockReservationService{}
			},
		},
	)
}

func TestPresentationService_Create(t *testing.T) {
	handler := getHandler()

	rsp := &proto.CreateResponse{}
	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, rsp)

	if err != nil {
		t.Fatalf("expected no error")
	}

	if rsp.CreatedId < 0 {
		t.Errorf("expected created presentation id to be non-negative")
	}
}

func TestPresentationService_Create_NegativeId(t *testing.T) {
	handler := getHandler()

	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: -5,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	if err == nil {
		t.Fatalf("expected error because of negative ids")
	}

	err = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 3,
			MovieId:  -6,
		},
	}, &proto.CreateResponse{})

	if err == nil {
		t.Fatalf("expected error because of negative ids")
	}
}

func TestPresentationService_Create_AlreadyExistent(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	if err == nil {
		t.Fatalf("expected error because of duplicate presentations")
	}
}

func TestPresentationService_FindForCinema(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  6,
		},
	}, &proto.CreateResponse{})

	rsp := &proto.FindForCinemaResponse{}
	err := handler.FindForCinema(context.TODO(), &proto.FindForCinemaRequest{
		CinemaId: 1,
	}, rsp)

	if err != nil {
		t.Fatalf("expected no error")
	}

	if len(rsp.Ids) != 2 {
		t.Fatalf("expected 2 presentation ids; got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 2 {
		t.Fatalf("expected 2 presentations; got %d", len(rsp.Dates))
	}
}

func TestPresentationService_FindForMovie(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  6,
		},
	}, &proto.CreateResponse{})

	rsp := &proto.FindForMovieResponse{}
	err := handler.FindForMovie(context.TODO(), &proto.FindForMovieRequest{
		MovieId: 2,
	}, rsp)

	if err != nil {
		t.Fatalf("expected no error")
	}

	if len(rsp.Ids) != 1 {
		t.Fatalf("expected 1 presentation id; got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 1 {
		t.Fatalf("expected 1 presentations; got %d", len(rsp.Dates))
	}

	if rsp.Dates[0].CinemaId != 1 {
		t.Errorf("expected cinema to have the id 1; got %d", rsp.Dates[0].CinemaId)
	}
}

func TestPresentationService_Read(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, createRsp)

	id := createRsp.CreatedId

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: id,
	}, rsp)

	if err != nil {
		t.Fatalf("expected no error")
	}

	if rsp.Data.CinemaId != 1 || rsp.Data.MovieId != 2 {
		t.Errorf("expected cinema id %d and movie id %d; got %d and %d", 1, 2, rsp.Data.CinemaId, rsp.Data.MovieId)
	}
}

func TestPresentationService_Read_NotFound(t *testing.T) {
	handler := getHandler()

	rsp := &proto.ReadResponse{}
	err := handler.Read(context.TODO(), &proto.ReadRequest{
		Id: 4,
	}, rsp)

	if err == nil {
		t.Fatalf("expected error because the presentation is unavailable")
	}
}

func TestPresentationService_ReadAll(t *testing.T) {
	handler := getHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  6,
		},
	}, &proto.CreateResponse{})

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 4,
			MovieId:  6,
		},
	}, &proto.CreateResponse{})

	rsp := &proto.ReadAllResponse{}
	err := handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, rsp)

	if err != nil {
		t.Fatalf("expected no error")
	}

	if len(rsp.Ids) != 3 {
		t.Fatalf("expected 3 presentation ids; got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 3 {
		t.Fatalf("expected 3 presentations; got %d", len(rsp.Dates))
	}
}

func TestPresentationService_Delete(t *testing.T) {
	handler := getHandler()

	createRsp := &proto.CreateResponse{}
	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, createRsp)

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: createRsp.CreatedId,
	}, &proto.DeleteResponse{})

	if err != nil {
		t.Errorf("expected no error")
	}
}

func TestPresentationService_Delete_NotExistent(t *testing.T) {
	handler := getHandler()

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Id: 36,
	}, &proto.DeleteResponse{})

	if err == nil {
		t.Errorf("expected error because of non-existent presentation")
	}
}
