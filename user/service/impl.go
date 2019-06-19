package service

import (
	"context"
	"fmt"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"log"
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
	log.Printf("Create | Creating user with name: %s\n", request.Data.Name)

	if len(request.Data.Name) == 0 {
		err := fmt.Errorf("cannot create user with empty name")
		log.Printf("Create | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	h.lastID++
	h.users[h.lastID] = request.Data

	response.CreatedId = h.lastID

	log.Printf("Create | User with id %d successfully created\n", response.CreatedId)
	return nil
}

func (h *UserServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	log.Printf("Read | Reading user with id %d\n", request.Id)

	h.mux.RLock()
	defer h.mux.RUnlock()

	data, ok := h.users[request.Id]
	if !ok {
		err := fmt.Errorf("could not find user with id %d", request.Id)
		log.Printf("Read | ERROR -> %s\n", err.Error())
		return err
	}

	response.Id = request.Id
	response.Data = data

	log.Printf("User with id %d has been successfully read\n", response.Id)
	return nil
}

func (h *UserServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	log.Printf("ReadAll | Reading all users...\n")

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

	log.Printf("ReadAll | %d users successfully read\n", len(response.Ids))
	return nil
}

func (h *UserServiceHandler) Update(context context.Context, request *proto.UpdateRequest, response *proto.UpdateResponse) error {
	log.Printf("Update | Updating user with id %d to have the name %s\n", request.Id, request.Data.Name)

	h.mux.Lock()
	defer h.mux.Unlock()

	if len(request.Data.Name) == 0 {
		err := fmt.Errorf("name of a user cannot be empty")
		log.Printf("Update | ERROR -> %s\n", err.Error())
		return err
	}

	data, ok := h.users[request.Id]
	if !ok {
		err := fmt.Errorf("user to update with id %d could not be found", request.Id)
		log.Printf("Update | ERROR -> %s\n", err.Error())
		return err
	}

	data.Name = request.Data.Name

	log.Printf("Update | Successfully updated user with id %d to have the name %s\n", request.Id, data.Name)
	return nil
}

func (h *UserServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	log.Printf("Delete | Deleting user with id %d\n", request.Id)

	h.mux.Lock()

	_, ok := h.users[request.Id]
	if !ok {
		err := fmt.Errorf("user to delete with id %d could not be found", request.Id)
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	delete(h.users, request.Id)

	h.mux.Unlock()

	// Notify reservation service that the user has been deleted -> Remove all reservations related to the user
	err := h.deleteRelatedReservations(context, []int64{request.Id})
	if err != nil {
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	log.Printf("Delete | Successfully deleted user with id %d\n", request.Id)
	return nil
}

// Delete all reservations related to the passed user ids.
func (h *UserServiceHandler) deleteRelatedReservations(context context.Context, userIDs []int64) error {
	reservationService := h.getReservationService()

	_, err := reservationService.CancelForUsers(context, &reservation.CancelForUsersRequest{
		UserIds: userIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

// Get an instance of the reservation service.
func (h *UserServiceHandler) getReservationService() reservation.ReservationService {
	return h.dependencies.ReservationService()
}
