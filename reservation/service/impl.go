package service

import (
	"context"
	"fmt"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"sync"
)

// Implementation of the reservation service handler.
type ReservationServiceHandler struct {
	lastID       int64
	reservations map[int64]*proto.ReservationData
	dependencies ReservationServiceDependencies
	toAccept     map[int64]bool
	mux          sync.RWMutex
}

// Dependencies the reservation service depends on.
type ReservationServiceDependencies struct {
	PresentationService func() presentation.PresentationService
	UserService         func() user.UserService
}

func NewReservationServiceHandler(dependencies *ReservationServiceDependencies) *ReservationServiceHandler {
	reservations := make(map[int64]*proto.ReservationData)
	toAccept := make(map[int64]bool)

	return &ReservationServiceHandler{
		lastID:       0,
		reservations: reservations,
		toAccept:     toAccept,
		dependencies: *dependencies,
	}
}

// Find the cinema ID the passed presentation is assigned to.
func (h *ReservationServiceHandler) getCinemaIDForPresentation(context context.Context, presentationID int64) (int64, error) {
	presentationService := h.getPresentationService()

	rsp, err := presentationService.Read(context, &presentation.ReadRequest{
		Id: presentationID,
	})
	if err != nil {
		return -1, err
	}

	return rsp.Data.CinemaId, nil
}

// Check if the passed seats are still available in the passed presentation.
func (h *ReservationServiceHandler) checkSeatsAvailable(context context.Context, seats []*proto.Seat, presentationID int64) (bool, error) {
	cinemaID, err := h.getCinemaIDForPresentation(context, presentationID)
	if err != nil {
		return false, err
	}

	fmt.Println(cinemaID)
	fmt.Println(seats)
	// TODO Ask cinema if seats are still available

	return true, nil
}

// Try to mark the passed seats as available (or not if false is passed).
// Will fail if the seats are not available.
func (h *ReservationServiceHandler) markSeatsAsAvailable(available bool, context context.Context, seats []*proto.Seat, presentationID int64) error {
	cinemaID, err := h.getCinemaIDForPresentation(context, presentationID)
	if err != nil {
		return err
	}

	fmt.Println(cinemaID)
	fmt.Println(available)
	// TODO Ask cinema to mark seats to be available or fail

	return nil
}

// Check if the passed user is available.
func (h *ReservationServiceHandler) checkUserAvailable(context context.Context, userID int64) bool {
	userService := h.getUserService()

	rsp, err := userService.Read(context, &user.ReadRequest{
		Id: userID,
	})

	return err == nil && rsp.Id == userID
}

// Get an instance of the user service.
func (h *ReservationServiceHandler) getUserService() user.UserService {
	return h.dependencies.UserService()
}

// Get an instance of the presentation service.
func (h *ReservationServiceHandler) getPresentationService() presentation.PresentationService {
	return h.dependencies.PresentationService()
}

func (h *ReservationServiceHandler) Reserve(context context.Context, request *proto.ReservationRequest, response *proto.ReservationResponse) error {
	if !h.checkUserAvailable(context, request.Data.UserId) {
		return fmt.Errorf("could not find user with id %d to make reservation for", request.Data.UserId)
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	available, err := h.checkSeatsAvailable(context, request.Data.Seats, request.Data.PresentationId)
	if err != nil {
		return err
	}

	if !available {
		response.Available = false
		return nil
	}

	// Mark reservation as available
	response.Available = true

	h.lastID++
	id := h.lastID
	response.CreatedId = id

	h.reservations[id] = request.Data // Save for later
	h.toAccept[id] = true             // Mark to be accepted

	return nil
}

func (h *ReservationServiceHandler) AcceptReservation(context context.Context, request *proto.AcceptReservationRequest, response *proto.AcceptReservationResponse) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if reservation id is to be accepted
	_, ok := h.toAccept[request.Id]
	if !ok {
		return fmt.Errorf("the passed reservation id is not marked to be accepted")
	}

	// Fetch stored reservation
	reservation, ok := h.reservations[request.Id]
	if !ok {
		return fmt.Errorf("could not find reservation with the passed id")
	}

	// Try to mark seats as taken -> Will fail if seats have already been taken
	err := h.markSeatsAsAvailable(false, context, reservation.Seats, reservation.PresentationId)
	if err != nil {
		return err
	}

	// Remove from marked for acception set
	delete(h.toAccept, request.Id)

	return nil
}

func (h *ReservationServiceHandler) Cancel(context context.Context, request *proto.CancelReservationRequest, response *proto.CancelReservationResponse) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if reservation exists
	reservation, ok := h.reservations[request.ReservationId]
	if !ok {
		return fmt.Errorf("cannot cancel a non-existing reservation")
	}

	// Mark seats as available
	err := h.markSeatsAsAvailable(true, context, reservation.Seats, reservation.PresentationId)
	if err != nil {
		return err
	}

	// Delete reservation
	delete(h.reservations, request.ReservationId)

	return nil
}

func (h *ReservationServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	h.mux.RLock()
	defer h.mux.RUnlock()

	ids := make([]int64, 0, len(h.reservations))
	dates := make([]*proto.ReservationData, 0, len(h.reservations))
	for id, reservation := range h.reservations {
		_, toBeAccepted := h.toAccept[id]

		if !toBeAccepted {
			ids = append(ids, id)
			dates = append(dates, reservation)
		}
	}

	response.Ids = ids
	response.Dates = dates

	return nil
}
