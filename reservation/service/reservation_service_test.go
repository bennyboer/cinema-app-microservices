package service

import (
	"context"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	cs "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/service"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	ps "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/service"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	us "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/service"
	"testing"
)

func getHandler() *ReservationServiceHandler {
	return NewReservationServiceHandler(
		&ReservationServiceDependencies{
			UserService: func() user.UserService {
				return &us.MockUserService{}
			},
			PresentationService: func() presentation.PresentationService {
				return &ps.MockPresentationService{}
			},
			CinemaService: func() cinema.CinemaService {
				return &cs.MockCinemaService{}
			},
		},
	)
}

func TestReservationService_Reserve(t *testing.T) {
	handler := getHandler()

	rsp := &proto.ReservationResponse{}
	err := handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, rsp)

	if err != nil {
		t.Fatalf("expected no error | %s", err.Error())
	}

	if !rsp.Available {
		t.Errorf("expected seats to be available")
	}

	if rsp.CreatedId < 0 {
		t.Errorf("expected created ID to be non-negative")
	}
}

func TestReservationService_AcceptReservation(t *testing.T) {
	handler := getHandler()

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)

	if !reservationRsp.Available {
		t.Fatalf("expected seats to be available")
	}

	err := handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	if err != nil {
		t.Fatalf("expected no error")
	}
}

func TestReservationService_AcceptReservation_NotAvailable(t *testing.T) {
	handler := getHandler()

	err := handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: 5,
	}, &proto.AcceptReservationResponse{})

	if err == nil {
		t.Fatalf("expected an error because of a missing reservation")
	}
}

func TestReservationService_AcceptReservation_AlreadyAccepted(t *testing.T) {
	handler := getHandler()

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	err := handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	if err == nil {
		t.Errorf("expected error because the reservation has already been accepted")
	}
}

func TestReservationService_ReadAll(t *testing.T) {
	handler := getHandler()

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)
	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	rsp := &proto.ReadAllResponse{}
	err := handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, rsp)

	if err != nil {
		t.Fatalf("expected no error | %s", err.Error())
	}

	if len(rsp.Ids) != 1 {
		t.Fatalf("expected to have 1 reservation id; got %d", len(rsp.Ids))
	}

	if len(rsp.Dates) != 1 {
		t.Fatalf("expected to have 1 reservation; got %d", len(rsp.Dates))
	}
}

func TestReservationService_CancelReservation(t *testing.T) {
	handler := getHandler()

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	err := handler.Cancel(context.TODO(), &proto.CancelReservationRequest{
		ReservationId: reservationRsp.CreatedId,
	}, &proto.CancelReservationResponse{})

	if err != nil {
		t.Errorf("expected no error")
	}
}

func TestReservationService_CancelReservation_NotAvailable(t *testing.T) {
	handler := getHandler()

	err := handler.Cancel(context.TODO(), &proto.CancelReservationRequest{
		ReservationId: 5,
	}, &proto.CancelReservationResponse{})

	if err == nil {
		t.Errorf("expected error because of missing reservation")
	}
}

func TestReservationService_CancelForPresentations(t *testing.T) {
	handler := getHandler()

	// Make some reservations first

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	reservationRsp2 := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 2,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp2)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp2.CreatedId,
	}, &proto.AcceptReservationResponse{})

	reservationRsp3 := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 2,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp2)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp3.CreatedId,
	}, &proto.AcceptReservationResponse{})

	err := handler.CancelForPresentations(context.TODO(), &proto.CancelForPresentationsRequest{
		PresentationIds: []int64{2},
	}, &proto.CancelForPresentationsResponse{})

	if err != nil {
		t.Errorf("expected no error")
	}

	// Check that the reservation 1 is still there
	readRsp := &proto.ReadAllResponse{}
	err = handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, readRsp)
	if err != nil {
		t.Errorf("expected no error")
	}

	if len(readRsp.Ids) != 1 {
		t.Errorf("expected to have exactly one reservation left")
	}

	if readRsp.Ids[0] != reservationRsp.CreatedId {
		t.Errorf("expected the first reservation to be still there")
	}
}

func TestReservationService_CancelForUsers(t *testing.T) {
	handler := getHandler()

	// Make some reservations first

	reservationRsp := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         1,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	}, &proto.AcceptReservationResponse{})

	reservationRsp2 := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         2,
			PresentationId: 1,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp2)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp2.CreatedId,
	}, &proto.AcceptReservationResponse{})

	reservationRsp3 := &proto.ReservationResponse{}
	_ = handler.Reserve(context.TODO(), &proto.ReservationRequest{
		Data: &proto.ReservationData{
			UserId:         2,
			PresentationId: 2,
			Seats: []*proto.Seat{
				{Row: 1, Number: 1},
			},
		},
	}, reservationRsp2)

	_ = handler.AcceptReservation(context.TODO(), &proto.AcceptReservationRequest{
		Id: reservationRsp3.CreatedId,
	}, &proto.AcceptReservationResponse{})

	err := handler.CancelForUsers(context.TODO(), &proto.CancelForUsersRequest{
		UserIds: []int64{2},
	}, &proto.CancelForUsersResponse{})

	if err != nil {
		t.Errorf("expected no error")
	}

	// Check that the reservation 1 is still there
	readRsp := &proto.ReadAllResponse{}
	err = handler.ReadAll(context.TODO(), &proto.ReadAllRequest{}, readRsp)
	if err != nil {
		t.Errorf("expected no error")
	}

	if len(readRsp.Ids) != 1 {
		t.Errorf("expected to have exactly one reservation left")
	}

	if readRsp.Ids[0] != reservationRsp.CreatedId {
		t.Errorf("expected the first reservation to be still there")
	}
}
