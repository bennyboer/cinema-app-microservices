package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"sync"
)

func deleteTest(client client.Client) {
	reservationService := reservation.NewReservationService("reservation-service", client)
	presentationService := presentation.NewPresentationService("presentation-service", client)
	movieService := movie.NewMovieService("movie-service", client)
	userService := user.NewUserService("user-service", client)
	cinemaService := cinema.NewCinemaService("cinema-service", client)

	cinemaCreate, err := cinemaService.Create(context.TODO(), &cinema.CreateRequest{
		Name:  "Test-Kino",
		Seats: 10,
		Row:   2,
	})
	if err != nil {
		fmt.Errorf("could not create cinema. Error: \n%s", err.Error())
	} else {
		fmt.Printf("Created cinema : %v\n", cinemaCreate)
	}

	userCreate, _ := userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Herbert",
		},
	})

	movieCreate, _ := movieService.Create(context.TODO(), &movie.CreateRequest{
		Data: &movie.MovieData{
			Title: "Der große Günther",
		},
	})

	presentationCreate, _ := presentationService.Create(context.TODO(), &presentation.CreateRequest{
		Data: &presentation.PresentationData{
			CinemaId: cinemaCreate.Data.Id,
			MovieId:  movieCreate.CreatedId,
		},
	})

	seats := make([]*reservation.Seat, 1)
	seats[0] = &reservation.Seat{
		Row:    1,
		Number: 1,
	}

	reservationRsp, _ := reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
		Data: &reservation.ReservationData{
			Seats:          seats,
			PresentationId: presentationCreate.CreatedId,
			UserId:         userCreate.CreatedId,
		},
	})

	_, _ = reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	})

	_, err := cinemaService.Delete(context.TODO(), &cinema.DeleteRequest{
		Id: cinemaCreate.Data.Id,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	presentations, _ := presentationService.ReadAll(context.TODO(), &presentation.ReadAllRequest{})
	reservations, _ := reservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})

	if len(presentations.Ids) != 0 {
		println("Not all presentations were deleted!")
	}

	if len(reservations.Ids) != 0 {
		println("Not all reservations were deleted!")
	}
}

func reservationTest(client client.Client) {
	reservationService := reservation.NewReservationService("reservation-service", client)
	presentationService := presentation.NewPresentationService("presentation-service", client)
	movieService := movie.NewMovieService("movie-service", client)
	userService := user.NewUserService("user-service", client)
	cinemaService := cinema.NewCinemaService("cinema-service", client)

	cinemaCreate, _ := cinemaService.Create(context.TODO(), &cinema.CreateRequest{
		Name:  "Test-Kino",
		Seats: 10,
		Row:   2,
	})

	user1, _ := userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Herbert",
		},
	})

	user2, _ := userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Günther",
		},
	})

	movieCreate, _ := movieService.Create(context.TODO(), &movie.CreateRequest{
		Data: &movie.MovieData{
			Title: "Der große Günther",
		},
	})

	presentationCreate, _ := presentationService.Create(context.TODO(), &presentation.CreateRequest{
		Data: &presentation.PresentationData{
			CinemaId: cinemaCreate.Data.Id,
			MovieId:  movieCreate.CreatedId,
		},
	})

	seats := make([]*reservation.Seat, 1)
	seats[0] = &reservation.Seat{
		Row:    1,
		Number: 1,
	}

	var wg sync.WaitGroup
	var barrier sync.WaitGroup
	barrier.Add(1)
	wg.Add(2)

	go func() {
		barrier.Wait()

		rsp, _ := reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationCreate.CreatedId,
				UserId:         user1.CreatedId,
			},
		})
		fmt.Printf("1 | Reserve response: %v\n", rsp)

		acceptRsp, err := reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
			Id: rsp.CreatedId,
		})
		fmt.Printf("1 | Accept response: %v\n", acceptRsp)

		if err != nil {
			fmt.Println(err.Error())
		}

		wg.Done()
	}()

	go func() {
		barrier.Wait()

		rsp, _ := reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationCreate.CreatedId,
				UserId:         user2.CreatedId,
			},
		})
		fmt.Printf("2 | Reserve response: %v\n", rsp)

		acceptRsp, err := reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
			Id: rsp.CreatedId,
		})
		fmt.Printf("2 | Accept response: %v\n", acceptRsp)

		if err != nil {
			fmt.Println(err.Error())
		}

		wg.Done()
	}()

	barrier.Done()
	wg.Wait()

	reservations, _ := reservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})

	println("User ", reservations.Dates[0].UserId, " got the reservation")
}

func main() {
	service := micro.NewService()

	service.Init()
	deleteTest(service.Client())
	//reservationTest(service.Client())
}
