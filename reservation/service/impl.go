package service

import (
	"context"
	"fmt"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"log"
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
	CinemaService       func() cinema.CinemaService
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

	// Convert seats to cinema seats
	cinemaSeats := make([]*cinema.SeatData, 0, len(seats))
	for _, seatPtr := range seats {
		cinemaSeats = append(cinemaSeats, &cinema.SeatData{
			Row:  seatPtr.Row,
			Seat: seatPtr.Number,
		})
	}

	cinemaService := h.getCinemaService()

	rsp, err := cinemaService.AreAvailable(context, &cinema.AvailableRequest{
		Id:    cinemaID,
		Seats: cinemaSeats,
	})
	if err != nil {
		return false, err
	}

	return rsp.Available, nil
}

// Try to mark the passed seats as available (or not if false is passed).
// Will fail if the seats are not available.
func (h *ReservationServiceHandler) markSeatsAsAvailable(context context.Context, available bool, seats []*proto.Seat, presentationID int64) error {
	cinemaID, err := h.getCinemaIDForPresentation(context, presentationID)
	if err != nil {
		return err
	}

	// Convert seats to cinema seats
	cinemaSeats := make([]*cinema.SeatData, 0, len(seats))
	for _, seatPtr := range seats {
		cinemaSeats = append(cinemaSeats, &cinema.SeatData{
			Row:  seatPtr.Row,
			Seat: seatPtr.Number,
		})
	}

	cinemaService := h.getCinemaService()

	if available {
		// Free the seats
		_, err = cinemaService.Free(context, &cinema.OccupiedRequest{
			Id:    cinemaID,
			Seats: cinemaSeats,
		})
	} else {
		// Occupy the seats
		_, err = cinemaService.Occupy(context, &cinema.OccupiedRequest{
			Id:    cinemaID,
			Seats: cinemaSeats,
		})
	}

	return err
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

// Get an instance of the cinema service.
func (h *ReservationServiceHandler) getCinemaService() cinema.CinemaService {
	return h.dependencies.CinemaService()
}

func (h *ReservationServiceHandler) Reserve(context context.Context, request *proto.ReservationRequest, response *proto.ReservationResponse) error {
	log.Printf("Reserve | Trying to reserve seats %v for presentation id %d for user id %d\n", request.Data.Seats, request.Data.PresentationId, request.Data.UserId)

	if !h.checkUserAvailable(context, request.Data.UserId) {
		err := fmt.Errorf("could not find user with id %d to make reservation for", request.Data.UserId)
		log.Printf("Reserve | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	available, err := h.checkSeatsAvailable(context, request.Data.Seats, request.Data.PresentationId)
	if err != nil {
		log.Printf("Reserve | ERROR -> %s\n", err.Error())
		return err
	}

	if !available {
		response.Available = false

		log.Printf("Reserve | Could not reserve seats %v for presentation %d for user %d since they are unavailable\n", request.Data.Seats, request.Data.PresentationId, request.Data.UserId)
		return nil
	}

	// Mark reservation as available
	response.Available = true

	h.lastID++
	id := h.lastID
	response.CreatedId = id

	h.reservations[id] = request.Data // Save for later
	h.toAccept[id] = true             // Mark to be accepted

	log.Printf("Reserve | Successfully created reservation offer for seats %v, presentation id %d and user id %d which is yet to be accepted\n", request.Data.Seats, request.Data.PresentationId, request.Data.UserId)
	return nil
}

func (h *ReservationServiceHandler) AcceptReservation(context context.Context, request *proto.AcceptReservationRequest, response *proto.AcceptReservationResponse) error {
	log.Printf("AcceptReservation | Accepting reservation offer with id %d\n", request.Id)

	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if reservation id is to be accepted
	_, ok := h.toAccept[request.Id]
	if !ok {
		err := fmt.Errorf("the passed reservation id is not marked to be accepted")
		log.Printf("AcceptReservation | ERROR -> %s\n", err.Error())
		return err
	}

	// Fetch stored reservation
	reservation, ok := h.reservations[request.Id]
	if !ok {
		err := fmt.Errorf("could not find reservation with the passed id")
		log.Printf("AcceptReservation | ERROR -> %s\n", err.Error())
		return err
	}

	// Try to mark seats as taken -> Will fail if seats have already been taken
	err := h.markSeatsAsAvailable(context, false, reservation.Seats, reservation.PresentationId)
	if err != nil {
		log.Printf("AcceptReservation | ERROR while trying to mark seats as taken -> %s\n", err.Error())
		return err
	}

	// Remove from marked for acceptation set
	delete(h.toAccept, request.Id)

	log.Printf("AcceptReservation | Successfully accepted reservation offer with id %d\n", request.Id)
	return nil
}

func (h *ReservationServiceHandler) Cancel(context context.Context, request *proto.CancelReservationRequest, response *proto.CancelReservationResponse) error {
	log.Printf("Cancel | Cancelling reservation with id %d\n", request.ReservationId)

	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if reservation exists
	reservation, ok := h.reservations[request.ReservationId]
	if !ok {
		err := fmt.Errorf("cannot cancel a non-existing reservation")
		log.Printf("Cancel | ERROR -> %s\n", err.Error())
		return err
	}

	// Mark seats as available
	err := h.markSeatsAsAvailable(context, true, reservation.Seats, reservation.PresentationId)
	if err != nil {
		log.Printf("Cancel | ERROR while trying to free seats -> %s\n", err.Error())
		return err
	}

	// Delete reservation
	delete(h.reservations, request.ReservationId)

	log.Printf("Cancel | Successfully cancelled reservation with id %d\n", request.ReservationId)
	return nil
}

func (h *ReservationServiceHandler) CancelForPresentations(context context.Context, request *proto.CancelForPresentationsRequest, response *proto.CancelForPresentationsResponse) error {
	log.Printf("CancelForPresentations | Cancelling all reservations for presentation ids %v\n", request.PresentationIds)

	// Create presentation id lookup
	lp := make(map[int64]bool, len(request.PresentationIds))
	for _, presentationId := range request.PresentationIds {
		lp[presentationId] = true
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	for id, reservation := range h.reservations {
		if _, del := lp[reservation.PresentationId]; del {
			// Delete reservation
			delete(h.reservations, id)
			log.Printf("CancelForPresentations | Deleted presentation with id %d\n", id)

			// Delete if in to be accepted
			if _, toBeAccepted := h.toAccept[id]; toBeAccepted {
				// Delete entry
				delete(h.toAccept, id)
			}
		}
	}

	log.Printf("CancelForPresentations | Successfully cancelled all reservations for presentation ids %v\n", request.PresentationIds)
	return nil
}

func (h *ReservationServiceHandler) CancelForUsers(context context.Context, request *proto.CancelForUsersRequest, response *proto.CancelForUsersResponse) error {
	log.Printf("CancelForUsers | Cancelling all reservations for user ids %v\n", request.UserIds)

	// Create user id lookup
	lp := make(map[int64]bool, len(request.UserIds))
	for _, userId := range request.UserIds {
		lp[userId] = true
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	for id, reservation := range h.reservations {
		if _, del := lp[reservation.UserId]; del {
			// Delete reservation
			delete(h.reservations, id)
			log.Printf("CancelForUsers | Deleted presentation with id %d\n", id)

			// Delete if in to be accepted
			if _, toBeAccepted := h.toAccept[id]; toBeAccepted {
				// Delete entry
				delete(h.toAccept, id)
			}
		}
	}

	log.Printf("CancelForUsers | Successfully cancelled all reservations for user ids %v\n", request.UserIds)
	return nil
}

func (h *ReservationServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	log.Printf("ReadAll | Reading all reservations...\n")

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

	log.Printf("ReadAll | Successfully read %d reservations\n", len(response.Ids))
	return nil
}
