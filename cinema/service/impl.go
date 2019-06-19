package service

import (
	"context"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"log"
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
	log.Printf("Create | Creating cinema with name %s, %d rows and %d seats\n", in.Name, in.Row, in.Seats)

	handler.mux.Lock()
	defer handler.mux.Unlock()

	if len(in.Name) == 0 {
		err := fmt.Errorf("cannot create a cinema with an empty name")
		log.Printf("Create | ERROR -> %s\n", err.Error())
		return err
	}

	handler.lastID++

	// Create seats
	seatCount := in.Seats * in.Row
	seats := make([]*proto.SeatData, 0, seatCount)
	for i := int64(0); i < in.Row; i++ {
		for j := int64(0); j < in.Seats; j++ {
			seats = append(seats, &proto.SeatData{Row: i + 1, Seat: j + 1, Occupied: false})
		}
	}

	data := proto.CinemaData{
		Name:      in.Name,
		Id:        handler.lastID,
		Seats:     seats,
		RowCount:  in.Row,
		SeatCount: in.Seats,
	}

	handler.cinemas[data.Id] = &data
	out.Data = handler.cinemas[data.Id]

	log.Printf("Create | Successfully created cinema with id %d\n", data.Id)
	return nil
}

func (handler *CinemaServiceHandler) Delete(ctx context.Context, in *proto.DeleteRequest, out *proto.DeleteResponse) error {
	log.Printf("Delete | Deleting cinema with id %d\n", in.Id)

	handler.mux.Lock()

	if _, ok := handler.cinemas[in.Id]; !ok {
		handler.mux.Unlock()

		err := fmt.Errorf("sorry, cannot find cinema with %d", in.Id)
		log.Printf("Delete | ERROR -> %s\n", err.Error())
		return err
	}

	delete(handler.cinemas, in.Id)

	handler.mux.Unlock()

	presentationService := handler.dependencies.PresentationService()
	_, err := presentationService.DeleteForCinemas(ctx, &presentation.DeleteForCinemasRequest{
		CinemaIds: []int64{in.Id},
	})
	if err != nil {
		err2 := fmt.Errorf("failed to delete presentations for cinema id %d. Error:\n%s", in.Id, err.Error())
		log.Printf("Delete | ERROR -> %s\n", err2.Error())
		return err2
	}

	log.Printf("Delete | Successfully deleted cinema with id %d\n", in.Id)
	return nil
}

func (handler *CinemaServiceHandler) Read(ctx context.Context, in *proto.ReadRequest, out *proto.ReadResponse) error {
	log.Printf("Read | Reading cinema with id %d\n", in.Id)

	handler.mux.RLock()
	defer handler.mux.RUnlock()

	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		out.Success = false

		err := fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
		log.Printf("Read | ERROR -> %s\n", err.Error())
		return err
	}
	out.Data = cinema
	out.Success = true

	log.Printf("Read | Successfully read cinema with id %d\n", in.Id)
	return nil
}

func (handler *CinemaServiceHandler) Occupy(ctx context.Context, in *proto.OccupiedRequest, out *proto.OccupiedResponse) error {
	log.Printf("Occupy | Occupying seats %v for cinema id %d\n", in.Seats, in.Id)

	handler.mux.Lock()
	defer handler.mux.Unlock()

	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		err := fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
		log.Printf("Occupy | ERROR -> %s\n", err.Error())
		return err
	}

	for _, seat := range in.Seats {
		cinema.Seats[((seat.Row-1)*cinema.SeatCount)+seat.Seat-1].Occupied = true
	}
	handler.cinemas[in.Id] = cinema

	out.Seats = cinema.Seats

	log.Printf("Occupy | Successfully occupied seats %v for cinema id %d\n", in.Seats, in.Id)
	return nil
}

func (handler *CinemaServiceHandler) Free(ctx context.Context, in *proto.OccupiedRequest, out *proto.OccupiedResponse) error {
	log.Printf("Free | Freeing seats %v for cinema id %d\n", in.Seats, in.Id)

	handler.mux.Lock()
	defer handler.mux.Unlock()

	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		err := fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
		log.Printf("Free | ERROR -> %s\n", err.Error())
		return err
	}

	for _, seat := range in.Seats {
		cinema.Seats[((seat.Row-1)*cinema.SeatCount)+seat.Seat-1].Occupied = false
	}
	handler.cinemas[in.Id] = cinema

	out.Seats = cinema.Seats

	log.Printf("Free | Successfully freed seats %v for cinema id %d\n", in.Seats, in.Id)
	return nil
}

func (handler *CinemaServiceHandler) List(ctx context.Context, in *proto.ListRequest, out *proto.ListResponse) error {
	log.Printf("List | Listing all cinemas...\n")

	handler.mux.RLock()
	defer handler.mux.RUnlock()

	size := len(handler.cinemas)
	data := make([]*proto.CinemaData, 0, size)

	for _, cinema := range handler.cinemas {
		data = append(data, cinema)
	}
	out.Data = data

	log.Printf("List | Successfully listed all %d cinemas\n", len(out.Data))
	return nil
}

func (handler *CinemaServiceHandler) AreAvailable(ctx context.Context, in *proto.AvailableRequest, out *proto.AvailableResponse) error {
	log.Printf("AreAvailable | Checking if seats %v are available for cinema %d\n", in.Seats, in.Id)

	handler.mux.RLock()
	defer handler.mux.RUnlock()

	cinema, ok := handler.cinemas[in.Id]

	if !ok {
		err := fmt.Errorf("sorry, couldn't find cinema with id %d", in.Id)
		log.Printf("AreAvailable | ERROR -> %s\n", err.Error())
		return err
	}

	available := true
	for _, seatPtr := range cinema.Seats {
		// Check if seat is to be checked
		check := false
		for _, sPtr := range in.Seats {
			if seatPtr.Row == sPtr.Row && seatPtr.Seat == sPtr.Seat {
				check = true
				break
			}
		}

		if check {
			if seatPtr.Occupied {
				available = false
				break
			}
		}
	}

	out.Available = available

	var logMessage string
	if available {
		logMessage = "Seats are available!"
	} else {
		logMessage = "Seats are unavailable!"
	}
	log.Printf("AreAvailable | Successfully checked whether seats %v in cinema %d are available -> %s\n", in.Seats, in.Id, logMessage)
	return nil
}
