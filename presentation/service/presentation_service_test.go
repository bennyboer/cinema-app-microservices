package service

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"testing"
)

func TestPresentationService_Create(t *testing.T) {
	handler := NewPresentationServiceHandler()

	err := handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	if err != nil {
		t.Fatalf("expected no error")
	}
}

func TestPresentationService_Create_NegativeId(t *testing.T) {
	handler := NewPresentationServiceHandler()

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
	handler := NewPresentationServiceHandler()

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
	handler := NewPresentationServiceHandler()

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

	if len(rsp.Dates) != 2 {
		t.Fatalf("expected 2 presentations; got %d", len(rsp.Dates))
	}

	if rsp.Dates[0].MovieId != 2 || rsp.Dates[1].MovieId != 6 {
		t.Errorf("expected movies to have the ids 2 and 6; got %d and %d", rsp.Dates[0].MovieId, rsp.Dates[1].MovieId)
	}
}

func TestPresentationService_FindForMovie(t *testing.T) {
	handler := NewPresentationServiceHandler()

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

	if len(rsp.Dates) != 1 {
		t.Fatalf("expected 1 presentations; got %d", len(rsp.Dates))
	}

	if rsp.Dates[0].CinemaId != 1 {
		t.Errorf("expected cinema to have the id 1; got %d", rsp.Dates[0].CinemaId)
	}
}

func TestPresentationService_ReadAll(t *testing.T) {
	handler := NewPresentationServiceHandler()

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

	if len(rsp.Dates) != 3 {
		t.Fatalf("expected 3 presentations; got %d", len(rsp.Dates))
	}
}

func TestPresentationService_Delete(t *testing.T) {
	handler := NewPresentationServiceHandler()

	_ = handler.Create(context.TODO(), &proto.CreateRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.CreateResponse{})

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.DeleteResponse{})

	if err != nil {
		t.Errorf("expected no error")
	}
}

func TestPresentationService_Delete_NotExistent(t *testing.T) {
	handler := NewPresentationServiceHandler()

	err := handler.Delete(context.TODO(), &proto.DeleteRequest{
		Data: &proto.PresentationData{
			CinemaId: 1,
			MovieId:  2,
		},
	}, &proto.DeleteResponse{})

	if err == nil {
		t.Errorf("expected error because of non-existent presentation")
	}
}
