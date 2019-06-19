package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"log"
	"sync"
)

// The movie service implementation.
type MovieServiceHandler struct {
	lastID       int64
	movies       map[int64]*proto.MovieData
	dependencies MovieServiceDependencies
	mux          sync.RWMutex
}

// Struct holding all services suppliers the movie service depends on.
type MovieServiceDependencies struct {
	PresentationService func() presentation.PresentationService
}

func NewMovieServiceHandler(dependencies *MovieServiceDependencies) *MovieServiceHandler {
	movies := make(map[int64]*proto.MovieData)

	return &MovieServiceHandler{
		lastID:       0,
		movies:       movies,
		dependencies: *dependencies,
	}
}

func (h *MovieServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	log.Printf("Create | Creating movie with title %s\n", request.Data.Title)

	if len(request.Data.Title) == 0 {
		err := fmt.Errorf("cannot create movie with empty title")
		log.Printf("Create | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	h.lastID++
	h.movies[h.lastID] = request.Data

	response.CreatedId = h.lastID

	log.Printf("Create | Successfully created movie with id %d and title %s\n", response.CreatedId, request.Data.Title)
	return nil
}

func (h *MovieServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	log.Printf("Read | Reading movie with id %d\n", request.Id)

	h.mux.RLock()
	defer h.mux.RUnlock()

	data, ok := h.movies[request.Id]
	if !ok {
		err := fmt.Errorf("could not find movie with id %d", request.Id)
		log.Printf("Read | ERROR -> %s\n", err.Error())
		return err
	}

	response.Id = request.Id
	response.Data = data

	log.Printf("Read | Successfully read movie with id %d and title %s\n", response.Id, response.Data.Title)
	return nil
}

func (h *MovieServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	log.Printf("ReadAll | Reading all movies...\n")

	h.mux.RLock()
	defer h.mux.RUnlock()

	size := len(h.movies)

	ids := make([]int64, 0, size)
	dates := make([]*proto.MovieData, 0, size)

	for id, data := range h.movies {
		ids = append(ids, id)
		dates = append(dates, data)
	}

	response.Ids = ids
	response.Dates = dates

	log.Printf("ReadAll | Successfully read %d movies\n", len(response.Ids))
	return nil
}

func (h *MovieServiceHandler) Update(context context.Context, request *proto.UpdateRequest, response *proto.UpdateResponse) error {
	log.Printf("Update | Updating movie with id %d to have title %s\n", request.Id, request.Data.Title)

	if len(request.Data.Title) == 0 {
		err := fmt.Errorf("title of a movie cannot be empty")
		log.Printf("Update | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	data, ok := h.movies[request.Id]
	if !ok {
		err := fmt.Errorf("movie to update with id %d could not be found", request.Id)
		log.Printf("Update | ERROR -> %s\n", err.Error())
		return err
	}

	data.Title = request.Data.Title

	log.Printf("Update | Successfully updated movie with id %d to have title %s\n", request.Id, data.Title)
	return nil
}

func (h *MovieServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	log.Printf("Delete | Deleting movie with id %d\n", request.Id)

	h.mux.Lock()

	_, ok := h.movies[request.Id]
	if !ok {
		err := fmt.Errorf("movie to delete with id %d could not be found", request.Id)
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	delete(h.movies, request.Id)

	h.mux.Unlock()

	// Notify presentation service that movie has been deleted -> Cancel all related presentations
	err := h.cancelRelatedPresentations(context, []int64{request.Id})
	if err != nil {
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	log.Printf("Delete | Successfully deleted movie with id %d\n", request.Id)
	return nil
}

// Cancel all presentations related to the passed movie ids.
func (h *MovieServiceHandler) cancelRelatedPresentations(context context.Context, movieIDs []int64) error {
	presentationService := h.getPresentationService()

	_, err := presentationService.DeleteForMovies(context, &presentation.DeleteForMoviesRequest{
		MovieIds: movieIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

// Get the presentation service
func (h *MovieServiceHandler) getPresentationService() presentation.PresentationService {
	return h.dependencies.PresentationService()
}
