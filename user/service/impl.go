package service

import (
	"context"
	"fmt"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"sync"
)

/// The user service implementation.
type UserServiceHandler struct {
	lastID       int64
	users        map[int64]*proto.UserData
	dependencies UserServiceDependencies
	mux          sync.RWMutex
}

// Services the user service depends on.
type UserServiceDependencies struct {
	ReservationService func() reservation.ReservationService
}

func NewUserServiceHandler(dependencies *UserServiceDependencies) *UserServiceHandler {
	users := make(map[int64]*proto.UserData)

	return &UserServiceHandler{
		lastID:       0,
		users:        users,
		dependencies: *dependencies,
	}
}

func (h *UserServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	if len(request.Data.Name) == 0 {
		return fmt.Errorf("cannot create user with empty name")
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	h.lastID++
	h.users[h.lastID] = request.Data

	response.CreatedId = h.lastID
	return nil
}

func (h *UserServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	h.mux.RLock()
	defer h.mux.RUnlock()

	data, ok := h.users[request.Id]
	if !ok {
		return fmt.Errorf("could not find user with id %d", request.Id)
	}

	response.Id = request.Id
	response.Data = data
	return nil
}

func (h *UserServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	h.mux.RLock()
	defer h.mux.RUnlock()

	size := len(h.users)

	ids := make([]int64, 0, size)
	dates := make([]*proto.UserData, 0, size)

	for id, data := range h.users {
		ids = append(ids, id)
		dates = append(dates, data)
	}

	response.Ids = ids
	response.Dates = dates

	return nil
}

func (h *UserServiceHandler) Update(context context.Context, request *proto.UpdateRequest, response *proto.UpdateResponse) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	if len(request.Data.Name) == 0 {
		return fmt.Errorf("name of a user cannot be empty")
	}

	data, ok := h.users[request.Id]
	if !ok {
		return fmt.Errorf("user to update with id %d could not be found", request.Id)
	}

	data.Name = request.Data.Name

	return nil
}

func (h *UserServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	_, ok := h.users[request.Id]
	if !ok {
		return fmt.Errorf("user to delete with id %d could not be found", request.Id)
	}

	delete(h.users, request.Id)

	// Notify reservation service that the user has been deleted -> Remove all reservations related to the user
	err := h.deleteRelatedReservations(context, request.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete all reservations related to the passed user id.
func (h *UserServiceHandler) deleteRelatedReservations(context context.Context, userID int64) error {
	reservationService := h.getReservationService()

	rsp, err := reservationService.ReadAll(context, &reservation.ReadAllRequest{})
	if err != nil {
		return err
	}

	for i := 0; i < len(rsp.Ids); i++ {
		reservationID := rsp.Ids[i]
		data := rsp.Dates[i]

		if data.UserId == userID {
			_, err := reservationService.Cancel(context, &reservation.CancelReservationRequest{
				ReservationId: reservationID,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Get an instance of the reservation service.
func (h *UserServiceHandler) getReservationService() reservation.ReservationService {
	return h.dependencies.ReservationService()
}
