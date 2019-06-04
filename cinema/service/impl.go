package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"sync"
)

type CinemaServiceHandler struct {
	lastID       int64
	cinemas      map[int64]*proto.CinemaData
	dependencies CinemaServiceDependencies
	mux          sync.RWMutex
}

type CinemaServiceDependencies struct {
	PresentationService func() presentation.PresentationService
}

func NewCinemaServiceHandler(dependencies *CinemaServiceDependencies) *CinemaServiceHandler {
	cinemas := make(map[int64]*proto.CinemaData)

	return &CinemaServiceHandler{
		lastID:       0,
		cinemas:      cinemas,
		dependencies: *dependencies,
		mux:          sync.RWMutex{},
	}
}

func (handler *CinemaServiceHandler) Create(ctx context.Context, in *proto.CreateRequest, out *proto.CreateResponse) error {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	if len(in.Name) == 0 {
		return fmt.Errorf("cannot create a cinema with an empty name")
	}

	handler.lastID++
	seats := make([]*proto.SeatData, in.Seats*in.Row)

	for i := int64(0); i < in.Row; i++ {
		for j := int64(0); j < in.Seats; j++ {
			seats[i*j+j] = &proto.SeatData{Row: i + 1, Seat: j + 1, Occupied: false}
		}
	}
	data := proto.CinemaData{Name: in.Name, Id: handler.lastID, Seats: seats}
	handler.cinemas[data.Id] = &data
	out.Data = handler.cinemas[data.Id]

	return nil
}

func (handler *CinemaServiceHandler) Delete(ctx context.Context, in *proto.DeleteRequest, out *proto.DeleteResponse) error {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	if _, ok := handler.cinemas[in.Id]; !ok {
		return fmt.Errorf("sorry, cannot find cinema with %d", in.Id)
	}

	delete(handler.cinemas, in.Id)

	presentationService := handler.dependencies.PresentationService()

	toDelete, err := presentationService.FindForCinema(ctx, &presentation.FindForCinemaRequest{CinemaId: in.Id})
	if err != nil {
		return fmt.Errorf("couldn't look up presentations for cinema")
	}

	for data := range toDelete.Ids {
		_, err = presentationService.Delete(ctx, &presentation.DeleteRequest{Id: int64(data)})
		if err != nil {
			return fmt.Errorf("failed to delete presentation %d", data)
		}
	}

	return nil
}

func (handler *CinemaServiceHandler) Read(ctx context.Context, in *proto.ReadRequest, out *proto.ReadResponse) error {
	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		out.Success = false
		return fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
	}
	out.Data = cinema
	out.Success = true

	return nil
}

func (handler *CinemaServiceHandler) Occupy(ctx context.Context, in *proto.OccupiedRequest, out *proto.OccupiedResponse) error {
	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		return fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
	}

	for _, seat := range in.Seats {
		cinema.Seats[((seat.Row-1)*(seat.Seat-1))+seat.Seat-1].Occupied = true
	}
	handler.cinemas[in.Id] = cinema

	out.Seats = cinema.Seats
	return nil
}

func (handler *CinemaServiceHandler) Free(ctx context.Context, in *proto.OccupiedRequest, out *proto.OccupiedResponse) error {
	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		return fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
	}

	for _, seat := range in.Seats {
		cinema.Seats[((seat.Row-1)*(seat.Seat-1))+seat.Seat-1].Occupied = false
	}
	handler.cinemas[in.Id] = cinema

	out.Seats = cinema.Seats
	return nil
}

func (handler *CinemaServiceHandler) List(ctx context.Context, in *proto.ListRequest, out *proto.ListResponse) error {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	size := len(handler.cinemas)
	data := make([]*proto.CinemaData, 0, size)

	for _, cinema := range handler.cinemas {
		data = append(data, cinema)
	}
	out.Data = data

	return nil
}
