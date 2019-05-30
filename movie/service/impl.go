package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
)

// The movie service implementation.
type MovieServiceHandler struct {
	lastID int64
	movies map[int64]*proto.MovieData
}

func NewMovieServiceHandler() *MovieServiceHandler {
	movies := make(map[int64]*proto.MovieData)

	return &MovieServiceHandler{
		lastID: 0,
		movies: movies,
	}
}

func (h *MovieServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	if len(request.Data.Title) == 0 {
		return fmt.Errorf("cannot create movie with empty title")
	}

	h.lastID++
	h.movies[h.lastID] = request.Data

	response.CreatedId = h.lastID
	return nil
}

func (h *MovieServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	data, ok := h.movies[request.Id]
	if !ok {
		return fmt.Errorf("could not find movie with id %d", request.Id)
	}

	response.Id = request.Id
	response.Data = data
	return nil
}

func (h *MovieServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	size := len(h.movies)

	ids := make([]int64, 0, size)
	dates := make([]*proto.MovieData, 0, size)

	for id, data := range h.movies {
		ids = append(ids, id)
		dates = append(dates, data)
	}

	response.Ids = ids
	response.Dates = dates

	return nil
}

func (h *MovieServiceHandler) Update(context context.Context, request *proto.UpdateRequest, response *proto.UpdateResponse) error {
	if len(request.Data.Title) == 0 {
		return fmt.Errorf("title of a movie cannot be empty")
	}

	data, ok := h.movies[request.Id]
	if !ok {
		return fmt.Errorf("movie to update with id %d could not be found", request.Id)
	}

	data.Title = request.Data.Title

	return nil
}

func (h *MovieServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	_, ok := h.movies[request.Id]
	if !ok {
		return fmt.Errorf("movie to delete with id %d could not be found", request.Id)
	}

	delete(h.movies, request.Id)

	return nil
}
