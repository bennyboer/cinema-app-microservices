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

type Client struct {
	reservationService  reservation.ReservationService
	presentationService presentation.PresentationService
	movieService        movie.MovieService
	userService         user.UserService
	cinemaService       cinema.CinemaService
}

func NewClient(cl client.Client) *Client {
	return &Client{
		reservationService:  reservation.NewReservationService("reservation-service", cl),
		presentationService: presentation.NewPresentationService("presentation-service", cl),
		movieService:        movie.NewMovieService("movie-service", cl),
		userService:         user.NewUserService("user-service", cl),
		cinemaService:       cinema.NewCinemaService("cinema-service", cl),
	}
}

func (c *Client) setUpData(client client.Client) {
	fmt.Println("Creating 2 cinemas...")
	for i := 0; i < 2; i++ {
		cinemaC, err := c.cinemaService.Create(context.TODO(), &cinema.CreateRequest{
			Name:  fmt.Sprintf("Kino %d", i),
			Seats: 4,
			Row:   4,
		})

		if err != nil {
			fmt.Printf("couldn't create cinema %d\n", i)
		} else {
			fmt.Printf("create %v\n", cinemaC)
		}
	}

	fmt.Println("Creating 4 users...")
	for i := 0; i < 4; i++ {
		userC, err := c.userService.Create(context.TODO(), &user.CreateRequest{
			Data: &user.UserData{
				Name: fmt.Sprintf("User %d", i+1),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create user %d %s\n", i, err.Error())
		} else {
			fmt.Printf("Created user %v\n", userC)
		}
	}

	fmt.Println("Creating 4 movies...")
	for i := 0; i < 4; i++ {
		movieC, err := c.movieService.Create(context.TODO(), &movie.CreateRequest{
			Data: &movie.MovieData{
				Title: fmt.Sprintf("Film %d", i),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create movie %d %s\n", i, err.Error())
		} else {
			fmt.Printf("Created Movie %v\n", movieC)
		}
	}

	fmt.Println("Creating 4 presentations...")
	for i := 0; i < 4; i++ {
		presC, err := c.presentationService.Create(context.TODO(), &presentation.CreateRequest{
			Data: &presentation.PresentationData{
				MovieId:  int64(i),
				CinemaId: int64(i%2 + 1),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create presentation %d\n", i)
		} else {
			fmt.Printf("Created presentation %v\n", presC)
		}
	}

	fmt.Println("Creating 4 reservations...")
	for i := 0; i < 4; i++ {
		seats := make([]*reservation.Seat, 1)

		seats[0] = &reservation.Seat{
			Row:    int64(i + 1),
			Number: int64(i + 1),
		}

		reservationRsp, err := c.reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: int64(i + 1),
				UserId:         int64(i + 1),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create reservation %d %s\n", i, err.Error())
		} else {
			fmt.Printf("Created reservation %v\n", reservationRsp)

			resC, err := c.reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
				Id: reservationRsp.CreatedId,
			})

			if err != nil {
				fmt.Printf("couldn't create reservation accept %d %s\n", i, err.Error())
			} else {
				fmt.Printf("Created accept %v\n", resC)
			}
		}
	}
}

func (c *Client) deleteData(client client.Client) {
	cinemaList, _ := c.cinemaService.List(context.TODO(), &cinema.ListRequest{})

	for _, data := range cinemaList.Data {
		_, err := c.cinemaService.Delete(context.TODO(), &cinema.DeleteRequest{
			Id: data.Id,
		})

		if err != nil {
			fmt.Printf("error deleting cinema %d %s\n", data.Id, err.Error())
		}
	}

	cinemaList, _ = c.cinemaService.List(context.TODO(), &cinema.ListRequest{})
	if len(cinemaList.Data) != 0 {
		fmt.Printf("not all cinemas deleted")
	}

	userList, _ := c.userService.ReadAll(context.TODO(), &user.ReadAllRequest{})
	for _, data := range userList.Ids {
		_, err := c.userService.Delete(context.TODO(), &user.DeleteRequest{
			Id: data,
		})

		if err != nil {
			fmt.Printf("error deleting user %d %s\n", data, err.Error())
		}
	}
	userList, _ = c.userService.ReadAll(context.TODO(), &user.ReadAllRequest{})

	if userList.Ids != nil {
		fmt.Println("not all users deleted")
	}

	movies, _ := c.movieService.ReadAll(context.TODO(), &movie.ReadAllRequest{})
	if movies.Ids != nil {
		fmt.Println("Not all movies deleted")
	}

	presentations, _ := c.presentationService.ReadAll(context.TODO(), &presentation.ReadAllRequest{})
	if presentations.Ids != nil {
		fmt.Println("Not all presentations deleted")
	}

	reservations, _ := c.reservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})
	if len(reservations.Ids) != 0 {
		fmt.Println("Not all reservations deleted")
	}
}

func (c *Client) deleteTest(client client.Client) {
	cinemaCreate, err := c.cinemaService.Create(context.TODO(), &cinema.CreateRequest{
		Name:  "Test-Kino",
		Seats: 10,
		Row:   2,
	})
	if err != nil {
		fmt.Errorf("could not create cinema. Error: \n%s", err.Error())
	} else {
		fmt.Printf("Created cinema : %v\n", cinemaCreate)
	}

	userCreate, _ := c.userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Herbert",
		},
	})

	movieCreate, _ := c.movieService.Create(context.TODO(), &movie.CreateRequest{
		Data: &movie.MovieData{
			Title: "Der große Günther",
		},
	})

	presentationCreate, _ := c.presentationService.Create(context.TODO(), &presentation.CreateRequest{
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

	reservationRsp, _ := c.reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
		Data: &reservation.ReservationData{
			Seats:          seats,
			PresentationId: presentationCreate.CreatedId,
			UserId:         userCreate.CreatedId,
		},
	})

	_, _ = c.reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
		Id: reservationRsp.CreatedId,
	})

	_, err = c.cinemaService.Delete(context.TODO(), &cinema.DeleteRequest{
		Id: cinemaCreate.Data.Id,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	presentations, _ := c.presentationService.ReadAll(context.TODO(), &presentation.ReadAllRequest{})
	reservations, _ := c.reservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})

	if presentations.Ids != nil {
		println("Not all presentations were deleted!")
	}

	if reservations.Ids != nil {
		println("Not all reservations were deleted!")
	}
}

func (c *Client) reservationTest(client client.Client) {
	cinemaCreate, _ := c.cinemaService.Create(context.TODO(), &cinema.CreateRequest{
		Name:  "Test-Kino",
		Seats: 10,
		Row:   2,
	})

	user1, _ := c.userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Herbert",
		},
	})

	user2, _ := c.userService.Create(context.TODO(), &user.CreateRequest{
		Data: &user.UserData{
			Name: "Günther",
		},
	})

	movieCreate, _ := c.movieService.Create(context.TODO(), &movie.CreateRequest{
		Data: &movie.MovieData{
			Title: "Der große Günther",
		},
	})

	presentationCreate, _ := c.presentationService.Create(context.TODO(), &presentation.CreateRequest{
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

		rsp, _ := c.reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationCreate.CreatedId,
				UserId:         user1.CreatedId,
			},
		})
		fmt.Printf("1 | Reserve response: %v\n", rsp)

		acceptRsp, err := c.reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
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

		rsp, _ := c.reservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationCreate.CreatedId,
				UserId:         user2.CreatedId,
			},
		})
		fmt.Printf("2 | Reserve response: %v\n", rsp)

		acceptRsp, err := c.reservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
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

	reservations, _ := c.reservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})

	println("User ", reservations.Dates[0].UserId, " got the reservation")
}

func main() {
	service := micro.NewService()
	service.Init()

	testClient := NewClient(service.Client())

	testClient.setUpData(service.Client())

	// TESTS BEGIN
	testClient.deleteTest(service.Client())
	testClient.reservationTest(service.Client())
	// TESTS END

	testClient.deleteData(service.Client())
}
