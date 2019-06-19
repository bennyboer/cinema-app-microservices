package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"log"
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
	log.Printf("Create | Creating presentation for movie id %d in cinema id %d\n", request.Data.MovieId, request.Data.CinemaId)

	if request.Data.CinemaId < 0 || request.Data.MovieId < 0 {
		err := fmt.Errorf("ids within the presentation data need to be non-negative")
		log.Printf("Create | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	// Check if already exists
	_, found := h.findPresentation(request.Data.CinemaId, request.Data.MovieId)
	if found {
		err := fmt.Errorf("presentation is already available")
		log.Printf("Create | ERROR -> %s\n", err.Error())
		return err
	}

	// Create
	h.lastID++
	h.presentations[h.lastID] = request.Data

	response.CreatedId = h.lastID

	log.Printf("Create | Successfully created presentation with id %d\n", response.CreatedId)
	return nil
}

func (h *PresentationServiceHandler) FindForCinema(context context.Context, request *proto.FindForCinemaRequest, response *proto.FindForCinemaResponse) error {
	log.Printf("FindForCinema | Finding all presentations in cinema with id %d\n", request.CinemaId)

	if request.CinemaId < 0 {
		err := fmt.Errorf("invalid cinema id passed")
		log.Printf("FindForCinema | ERROR -> %s\n", err.Error())
		return err
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

	log.Printf("FindForCinema | Successfully found %d presentations for cinema id %d\n", len(response.Ids), request.CinemaId)
	return nil
}

func (h *PresentationServiceHandler) FindForMovie(context context.Context, request *proto.FindForMovieRequest, response *proto.FindForMovieResponse) error {
	log.Printf("FindForMovie | Finding all presentations for movie with id %d\n", request.MovieId)

	if request.MovieId < 0 {
		err := fmt.Errorf("ids within the presentation data need to be non-negative")
		log.Printf("FindForMovie | ERROR -> %s\n", err.Error())
		return err
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

	log.Printf("FindForMovie | Successfully found %d presentations for movie with id %d\n", len(response.Ids), request.MovieId)
	return nil
}

func (h *PresentationServiceHandler) Read(context context.Context, request *proto.ReadRequest, response *proto.ReadResponse) error {
	log.Printf("Read | Reading presentation with id %d\n", request.Id)

	if request.Id < 0 {
		err := fmt.Errorf("presentation id needs to be non-negative")
		log.Printf("Read | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.RLock()
	defer h.mux.RUnlock()

	data, ok := h.presentations[request.Id]
	if !ok {
		err := fmt.Errorf("could not find presentation with id %d", request.Id)
		log.Printf("Read | ERROR -> %s\n", err.Error())
		return err
	}

	response.Data = data

	log.Printf("Read | Successfully read presentation with id %d\n", request.Id)
	return nil
}

func (h *PresentationServiceHandler) ReadAll(context context.Context, request *proto.ReadAllRequest, response *proto.ReadAllResponse) error {
	log.Printf("ReadAll | Reading all presentations...\n")

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

	log.Printf("ReadAll | Successfully read %d presentations\n", len(response.Ids))
	return nil
}

func (h *PresentationServiceHandler) Delete(context context.Context, request *proto.DeleteRequest, response *proto.DeleteResponse) error {
	log.Printf("Delete | Deleting presentation with id %d\n", request.Id)

	if request.Id < 0 {
		err := fmt.Errorf("presentation id needs to be non-negative")
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	h.mux.Lock()

	_, ok := h.presentations[request.Id]
	if !ok {
		h.mux.Unlock()

		err := fmt.Errorf("presentation could not be found")
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	delete(h.presentations, request.Id)

	h.mux.Unlock()

	// Notify reservation service that the presentation has been deleted -> Delete all related reservations
	err := h.deleteRelatedReservations(context, []int64{request.Id})
	if err != nil {
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	log.Printf("Delete | Successfully deleted presentation with id %d\n", request.Id)
	return nil
}

func (h *PresentationServiceHandler) DeleteForCinemas(context context.Context, request *proto.DeleteForCinemasRequest, response *proto.DeleteForCinemasResponse) error {
	log.Printf("DeleteForCinemas | Deleting presentations for cinema ids %v\n", request.CinemaIds)

	// Create lookup for cinema Ids
	lp := make(map[int64]bool, len(request.CinemaIds))
	for _, cinemaID := range request.CinemaIds {
		lp[cinemaID] = true
	}

	h.mux.Lock()

	deletedIds := make([]int64, 0)
	for id, data := range h.presentations {
		if _, del := lp[data.CinemaId]; del {
			delete(h.presentations, id)

			deletedIds = append(deletedIds, id)
		}
	}

	h.mux.Unlock()

	// Notify reservation service to delete all reservations for the deleted presentations
	err := h.deleteRelatedReservations(context, deletedIds)
	if err != nil {
		log.Printf("DeleteForCinemas | ERROR -> %s\n", err.Error())
		return err
	}

	log.Printf("DeleteForCinemas | Successfully deleted presentations for cinema ids %v\n", request.CinemaIds)
	return nil
}

func (h *PresentationServiceHandler) DeleteForMovies(context context.Context, request *proto.DeleteForMoviesRequest, response *proto.DeleteForMoviesResponse) error {
	log.Printf("DeleteForMovies | Deleting presentations for movie ids %v\n", request.MovieIds)

	// Create lookup for movie Ids
	lp := make(map[int64]bool, len(request.MovieIds))
	for _, movieID := range request.MovieIds {
		lp[movieID] = true
	}

	h.mux.Lock()

	deletedIds := make([]int64, 0)
	for id, data := range h.presentations {
		if _, del := lp[data.MovieId]; del {
			delete(h.presentations, id)

			deletedIds = append(deletedIds, id)
		}
	}

	h.mux.Unlock()

	// Notify reservation service to delete all reservations for the deleted presentations
	err := h.deleteRelatedReservations(context, deletedIds)
	if err != nil {
		log.Printf("DeleteForMovies | ERROR -> %s\n", err.Error())
		return err
	}

	log.Printf("DeleteForMovies | Successfully deleted presentations for movie ids %v\n", request.MovieIds)
	return nil
}

func (h *PresentationServiceHandler) Clear(context.Context, *proto.ClearRequest, *proto.ClearResponse) error {
	log.Printf("Clear | Clearing all service data...\n")

	h.mux.Lock()
	defer h.mux.Unlock()

	h.presentations = make(map[int64]*proto.PresentationData)
	h.lastID = 0

	log.Printf("Clear | Successfully cleared all service data\n")
	return nil
}

// Delete all reservations related to the passed presentation id.
func (h *PresentationServiceHandler) deleteRelatedReservations(context context.Context, presentationIDs []int64) error {
	reservationService := h.getReservationService()

	_, err := reservationService.CancelForPresentations(context, &reservation.CancelForPresentationsRequest{
		PresentationIds: presentationIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

// Get a new instance of the reservation service.
func (h *PresentationServiceHandler) getReservationService() reservation.ReservationService {
	return h.dependencies.ReservationService()
}
