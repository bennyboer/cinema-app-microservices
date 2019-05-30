package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
)

/// The user service implementation.
type UserServiceHandler struct {
	lastID int64
	users  map[int64]*proto.UserData
}

func NewUserServiceHandler() *UserServiceHandler {
	users := make(map[int64]*proto.UserData)

	return &UserServiceHandler{
		lastID: 0,
		users:  users,
	}
}

func (h *UserServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	if len(request.Data.Name) == 0 {
		return fmt.Errorf("cannot create user with empty name")
	}

	h.lastID++
	h.users[h.lastID] = request.Data

	response.CreatedId = h.lastID
	return nil
}

func (h *UserServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	data, ok := h.users[request.Id]
	if !ok {
		return fmt.Errorf("could not find user with id %d", request.Id)
	}

	response.Id = request.Id
	response.Data = data
	return nil
}

func (h *UserServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
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
	_, ok := h.users[request.Id]
	if !ok {
		return fmt.Errorf("user to delete with id %d could not be found", request.Id)
	}

	delete(h.users, request.Id)

	// TODO Notify reservation service

	return nil
}
