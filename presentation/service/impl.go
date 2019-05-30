package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
)

const initCapacity = 64;

// Implementation of the presentation service handler.
type PresentationServiceHandler struct {
	presentations []*proto.PresentationData
}

func NewPresentationServiceHandler() *PresentationServiceHandler {
	presentations := make([]*proto.PresentationData, 0, initCapacity)

	return &PresentationServiceHandler{
		presentations: presentations,
	}
}

func (h *PresentationServiceHandler) findPresentation(cinemaId int64, movieId int64) (data *proto.PresentationData, found bool) {
	for _, data := range h.presentations {
		if data.CinemaId == cinemaId && data.MovieId == movieId {
			return data, true
		}
	}

	return nil, false
}

func (h *PresentationServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	if request.Data.CinemaId < 0 || request.Data.MovieId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	// Check if already exists
	_, found := h.findPresentation(request.Data.CinemaId, request.Data.MovieId)
	if found {
		return fmt.Errorf("presentation is already available")
	}

	// Create
	h.presentations = append(h.presentations, request.Data)

	return nil
}

func (h *PresentationServiceHandler) FindForCinema(context context.Context, request *proto.FindForCinemaRequest, response *proto.FindForCinemaResponse) error {
	if request.CinemaId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	result := make([]*proto.PresentationData, 0)
	for _, data := range h.presentations {
		if data.CinemaId == request.CinemaId {
			result = append(result, data)
		}
	}

	response.Dates = result

	return nil
}

func (h *PresentationServiceHandler) FindForMovie(context context.Context, request *proto.FindForMovieRequest, response *proto.FindForMovieResponse) error {
	if request.MovieId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	result := make([]*proto.PresentationData, 0)
	for _, data := range h.presentations {
		if data.MovieId == request.MovieId {
			result = append(result, data)
		}
	}

	response.Dates = result

	return nil
}

func (h *PresentationServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	response.Dates = h.presentations

	return nil
}

func (h *PresentationServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	if request.Data.CinemaId < 0 || request.Data.MovieId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	index2Remove := -1
	for i, data := range h.presentations {
		if data.CinemaId == request.Data.CinemaId && data.MovieId == request.Data.MovieId {
			index2Remove = i
			break
		}
	}

	if index2Remove == -1 {
		return fmt.Errorf("presentation could not be found")
	}

	h.presentations = append(h.presentations[:index2Remove], h.presentations[index2Remove+1:]...)

	return nil
}
