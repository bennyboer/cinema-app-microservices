package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"sync"
)

// Implementation of the presentation service handler.
type PresentationServiceHandler struct {
	lastID        int64
	presentations map[int64]*proto.PresentationData
	dependencies  PresentationServiceDependencies
	mux           sync.RWMutex
}

// Service dependencies of the presentation service.
type PresentationServiceDependencies struct {
	ReservationService func() reservation.ReservationService
}

func NewPresentationServiceHandler(dependencies *PresentationServiceDependencies) *PresentationServiceHandler {
	presentations := make(map[int64]*proto.PresentationData)

	return &PresentationServiceHandler{
		lastID:        0,
		presentations: presentations,
		dependencies:  *dependencies,
	}
}

func (h *PresentationServiceHandler) findPresentation(cinemaID int64, movieID int64) (data *proto.PresentationData, found bool) {
	for _, data := range h.presentations {
		if data.CinemaId == cinemaID && data.MovieId == movieID {
			return data, true
		}
	}

	return nil, false
}

func (h *PresentationServiceHandler) Create(context context.Context, request *proto.CreateRequest, response *proto.CreateResponse) error {
	if request.Data.CinemaId < 0 || request.Data.MovieId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if already exists
	_, found := h.findPresentation(request.Data.CinemaId, request.Data.MovieId)
	if found {
		return fmt.Errorf("presentation is already available")
	}

	// Create
	h.lastID++
	h.presentations[h.lastID] = request.Data

	response.CreatedId = h.lastID

	return nil
}

func (h *PresentationServiceHandler) FindForCinema(context context.Context, request *proto.FindForCinemaRequest, response *proto.FindForCinemaResponse) error {
	if request.CinemaId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	h.mux.RLock()
	defer h.mux.RUnlock()

	ids := make([]int64, 0)
	result := make([]*proto.PresentationData, 0)
	for id, data := range h.presentations {
		if data.CinemaId == request.CinemaId {
			ids = append(ids, id)
			result = append(result, data)
		}
	}

	response.Ids = ids
	response.Dates = result

	return nil
}

func (h *PresentationServiceHandler) FindForMovie(context context.Context, request *proto.FindForMovieRequest, response *proto.FindForMovieResponse) error {
	if request.MovieId < 0 {
		return fmt.Errorf("ids within the presentation data need to be non-negative")
	}

	h.mux.RLock()
	defer h.mux.RUnlock()

	ids := make([]int64, 0)
	result := make([]*proto.PresentationData, 0)
	for id, data := range h.presentations {
		if data.MovieId == request.MovieId {
			ids = append(ids, id)
			result = append(result, data)
		}
	}

	response.Ids = ids
	response.Dates = result

	return nil
}

func (h *PresentationServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	if request.Id < 0 {
		return fmt.Errorf("presentation id needs to be non-negative")
	}

	h.mux.RLock()
	defer h.mux.RUnlock()

	data, ok := h.presentations[request.Id]
	if !ok {
		return fmt.Errorf("could not find presentation with id %d", request.Id)
	}

	response.Data = data

	return nil
}

func (h *PresentationServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	h.mux.RLock()
	defer h.mux.RUnlock()

	ids := make([]int64, 0, len(h.presentations))
	result := make([]*proto.PresentationData, 0, len(h.presentations))
	for id, data := range h.presentations {
		ids = append(ids, id)
		result = append(result, data)
	}

	response.Ids = ids
	response.Dates = result

	return nil
}

func (h *PresentationServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	if request.Id < 0 {
		return fmt.Errorf("presentation id needs to be non-negative")
	}

	h.mux.Lock()

	_, ok := h.presentations[request.Id]
	if !ok {
		h.mux.Unlock()
		return fmt.Errorf("presentation could not be found")
	}

	delete(h.presentations, request.Id)

	h.mux.Unlock()

	// Notify reservation service that the presentation has been deleted -> Delete all related reservations
	err := h.deleteRelatedReservations(context, request.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete all reservations related to the passed presentation id.
func (h *PresentationServiceHandler) deleteRelatedReservations(context context.Context, presentationID int64) error {
	reservationService := h.getReservationService()

	rsp, err := reservationService.ReadAll(context, &reservation.ReadAllRequest{})
	if err != nil {
		return err
	}

	for i := 0; i < len(rsp.Ids); i++ {
		reservationID := rsp.Ids[i]
		data := rsp.Dates[i]

		if data.PresentationId == presentationID {
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

// Get a new instance of the reservation service.
func (h *PresentationServiceHandler) getReservationService() reservation.ReservationService {
	return h.dependencies.ReservationService()
}
