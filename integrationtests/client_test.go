package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"os"
	"sync"
	"testing"
)

func beforeTest() *Client {
	fmt.Println("RUNNING BEFORE TEST ROUTINE")
	service := micro.NewService()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = os.Args[:1]

	service.Init()

	c := NewClient(service.Client())
	setUpData(c)

	return c
}

func afterTest(c *Client) {
	fmt.Println("RUNNING AFTER TEST ROUTINE")
	cleanupData(c)
}

func TestServices_DeleteCascade(t *testing.T) {
	c := beforeTest()
	defer afterTest(c)

	// Start test
	cinemaID := c.Cinemas[0]

	// Check current state
	presentations, err := c.PresentationService.FindForCinema(context.TODO(), &presentation.FindForCinemaRequest{
		CinemaId: cinemaID,
	})
	if err != nil {
		t.Errorf("Could not fetch Presentations for cinema. Error:\n%s", err.Error())
		return
	}
	presentationCount := len(presentations.Ids)

	reservations, err := c.ReservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})
	if err != nil {
		t.Errorf("could not fetch Reservations. Error:\n%s", err.Error())
		return
	}
	reservationsForPresentations := make([]int64, 0)
	for i, reservationDataPtr := range reservations.Dates {
		for _, presentationID := range presentations.Ids {
			if reservationDataPtr.PresentationId == presentationID {
				reservationsForPresentations = append(reservationsForPresentations, reservations.Ids[i])
			}
		}
	}
	reservationCount := len(reservationsForPresentations)

	fmt.Printf(`DELETE TEST -> CURRENT STATE:
For cinema with id %d
	Presentations: %d
	Reservations: %d
`, cinemaID, presentationCount, reservationCount)

	fmt.Printf("Trying to delete cinema with id %d...\n", cinemaID)
	_, err = c.CinemaService.Delete(context.TODO(), &cinema.DeleteRequest{
		Id: c.Cinemas[0],
	})
	if err != nil {
		t.Errorf("Deleting cinema failed. Error:\n%s", err.Error())
		return
	}
	fmt.Printf("Cinema with id %d has been successfully deleted\n", cinemaID)

	// Check if Presentations and Reservations have been deleted
	presentations, err = c.PresentationService.FindForCinema(context.TODO(), &presentation.FindForCinemaRequest{
		CinemaId: cinemaID,
	})
	if err != nil {
		t.Errorf("could not fetch Presentations for the deleted cinema. Error:\n%s", err.Error())
		return
	}
	if len(presentations.Ids) != 0 {
		t.Errorf("Presentations have not been deleted properly")
		return
	} else {
		fmt.Printf("Presentations related to cinema with id %d have been successfully deleted by the service\n", cinemaID)
	}

	reservations, err = c.ReservationService.ReadAll(context.TODO(), &reservation.ReadAllRequest{})
	if err != nil {
		t.Errorf("Could not fetch Reservations. Error:\n%s", err.Error())
		return
	}
	for _, reservationID := range reservations.Ids {
		for _, oldReservationID := range reservationsForPresentations {
			if reservationID == oldReservationID {
				t.Errorf("reservation with id %d should have been deleted by now!\n", reservationID)
				return
			}
		}
	}
	fmt.Printf("Reservatiosn related to the cinema with id %d have been successfully deleted by the service\n", cinemaID)
}

func TestServices_ConcurrentReservation(t *testing.T) {
	c := beforeTest()
	defer afterTest(c)

	// Start test
	user1ID := c.Users[0]
	user2ID := c.Users[1]

	presentationID := c.Presentations[0]

	seats := []*reservation.Seat{
		{Row: 1, Number: 4},
	}

	// User1 and User2 want to reserve the same seats for the same presentation concurrently
	var wg sync.WaitGroup
	var barrier sync.WaitGroup
	barrier.Add(1)
	wg.Add(2)

	user1GotReservation := false
	user2GotReservation := false

	go func() {
		defer wg.Done()
		barrier.Wait()

		rsp, err := c.ReservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationID,
				UserId:         user1ID,
			},
		})
		if err != nil {
			t.Errorf("User 1 | could not reserve. Did not expect error: %s", err.Error())
			return
		} else {
			fmt.Printf("User 1 | Reserve response -> Created ID: %d and seats are available? %t\n", rsp.CreatedId, rsp.Available)
		}

		if !rsp.Available {
			t.Errorf("User 1 | seats are no more available")
			return
		}

		acceptRsp, err := c.ReservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
			Id: rsp.CreatedId,
		})
		fmt.Printf("User 1 | Accept response: %v\n", acceptRsp)
		if err != nil {
			fmt.Printf("User 1 | Could not accept the reservation. Error:\n%s\n", err.Error())
		} else {
			fmt.Printf("User 1 | Reservation has been accepted.\n")
			user1GotReservation = true
		}
	}()

	go func() {
		defer wg.Done()
		barrier.Wait()

		rsp, err := c.ReservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
			Data: &reservation.ReservationData{
				Seats:          seats,
				PresentationId: presentationID,
				UserId:         user2ID,
			},
		})
		if err != nil {
			t.Errorf("User 2 | could not reserve. Did not expect error: %s", err.Error())
			return
		} else {
			fmt.Printf("User 2 | Reserve response -> Created ID: %d and seats are available? %t\n", rsp.CreatedId, rsp.Available)
		}

		if !rsp.Available {
			t.Errorf("User 2 | seats are no more available")
			return
		}

		acceptRsp, err := c.ReservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
			Id: rsp.CreatedId,
		})
		fmt.Printf("User 2 | Accept response: %v\n", acceptRsp)
		if err != nil {
			fmt.Printf("User 2 | Could not accept the reservation. Error:\n%s\n", err.Error())
		} else {
			fmt.Printf("User 2 | Reservation has been accepted.\n")
			user2GotReservation = true
		}
	}()

	barrier.Done()
	wg.Wait()

	if (user1GotReservation && user2GotReservation) || (!user1GotReservation && !user2GotReservation) {
		t.Errorf("expected one user to get the reservation!")
	}
}

func setUpData(c *Client) {
	fmt.Println("Creating 2 Cinemas...")
	for i := 0; i < 2; i++ {
		cinemaC, err := c.CinemaService.Create(context.TODO(), &cinema.CreateRequest{
			Name:  fmt.Sprintf("Kino %d", i),
			Seats: 4,
			Row:   4,
		})

		if err != nil {
			fmt.Printf("couldn't create cinema %d\n", i)
		} else {
			fmt.Printf("create %v\n", cinemaC)
			c.Cinemas = append(c.Cinemas, cinemaC.Data.Id)
		}
	}

	fmt.Println("Creating 4 Users...")
	for i := 0; i < 4; i++ {
		userC, err := c.UserService.Create(context.TODO(), &user.CreateRequest{
			Data: &user.UserData{
				Name: fmt.Sprintf("User %d", i+1),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create user %d %s\n", i, err.Error())
		} else {
			fmt.Printf("Created user %v\n", userC)
			c.Users = append(c.Users, userC.CreatedId)
		}
	}

	fmt.Println("Creating 4 Movies...")
	for i := 0; i < 4; i++ {
		movieC, err := c.MovieService.Create(context.TODO(), &movie.CreateRequest{
			Data: &movie.MovieData{
				Title: fmt.Sprintf("Film %d", i),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create movie %d %s\n", i, err.Error())
		} else {
			fmt.Printf("Created Movie %v\n", movieC)
			c.Movies = append(c.Movies, movieC.CreatedId)
		}
	}

	fmt.Println("Creating 4 Presentations...")
	for i := 0; i < 4; i++ {
		presC, err := c.PresentationService.Create(context.TODO(), &presentation.CreateRequest{
			Data: &presentation.PresentationData{
				MovieId:  int64(i),
				CinemaId: int64(i%2 + 1),
			},
		})

		if err != nil {
			fmt.Printf("couldn't create presentation %d\n", i)
		} else {
			fmt.Printf("Created presentation %v\n", presC)
			c.Presentations = append(c.Presentations, presC.CreatedId)
		}
	}

	fmt.Println("Creating 4 Reservations...")
	for i := 0; i < 4; i++ {
		seats := make([]*reservation.Seat, 1)

		seats[0] = &reservation.Seat{
			Row:    int64(i + 1),
			Number: int64(i + 1),
		}

		reservationRsp, err := c.ReservationService.Reserve(context.TODO(), &reservation.ReservationRequest{
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

			resC, err := c.ReservationService.AcceptReservation(context.TODO(), &reservation.AcceptReservationRequest{
				Id: reservationRsp.CreatedId,
			})

			if err != nil {
				fmt.Printf("couldn't create reservation accept %d %s\n", i, err.Error())
			} else {
				fmt.Printf("Created accept %v\n", resC)
				c.Reservations = append(c.Reservations, reservationRsp.CreatedId)
			}
		}
	}
}

func cleanupData(c *Client) {
	if _, err := c.PresentationService.Clear(context.TODO(), &presentation.ClearRequest{}); err != nil {
		fmt.Printf("ERROR: Could not clear Presentations. Error:\n%s", err.Error())
	}

	if _, err := c.ReservationService.Clear(context.TODO(), &reservation.ClearRequest{}); err != nil {
		fmt.Printf("ERROR: Could not clear Reservations. Error:\n%s", err.Error())
	}

	if _, err := c.MovieService.Clear(context.TODO(), &movie.ClearRequest{}); err != nil {
		fmt.Printf("ERROR: Could not clear Movies. Error:\n%s", err.Error())
	}

	if _, err := c.CinemaService.Clear(context.TODO(), &cinema.ClearRequest{}); err != nil {
		fmt.Printf("ERROR: Could not clear Cinemas. Error:\n%s", err.Error())
	}

	if _, err := c.UserService.Clear(context.TODO(), &user.ClearRequest{}); err != nil {
		fmt.Printf("ERROR: Could not clear Users. Error:\n%s", err.Error())
	}
}
