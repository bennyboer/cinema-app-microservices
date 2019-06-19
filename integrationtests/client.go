package main

import (
	"github.com/micro/go-micro/client"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	movie "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
)

type Client struct {
	ReservationService  reservation.ReservationService
	PresentationService presentation.PresentationService
	MovieService        movie.MovieService
	UserService         user.UserService
	CinemaService       cinema.CinemaService
	Users               []int64
	Cinemas             []int64
	Movies              []int64
	Presentations       []int64
	Reservations        []int64
}

func NewClient(cl client.Client) *Client {
	return &Client{
		ReservationService:  reservation.NewReservationService("reservation-service", cl),
		PresentationService: presentation.NewPresentationService("presentation-service", cl),
		MovieService:        movie.NewMovieService("movie-service", cl),
		UserService:         user.NewUserService("user-service", cl),
		CinemaService:       cinema.NewCinemaService("cinema-service", cl),
	}
}
