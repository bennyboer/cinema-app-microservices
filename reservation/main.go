package main

import (
	"github.com/micro/go-micro"
	cinema "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/service"
	user "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	"log"
)

const serviceName = "reservation-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	err := proto.RegisterReservationHandler(service.Server(), s.NewReservationServiceHandler(
		&s.ReservationServiceDependencies{
			UserService: func() user.UserService {
				return user.NewUserService("user-service", service.Client())
			},
			PresentationService: func() presentation.PresentationService {
				return presentation.NewPresentationService("presentation-service", service.Client())
			},
			CinemaService: func() cinema.CinemaService {
				return cinema.NewCinemaService("cinema-service", service.Client())
			},
		},
	))
	if err != nil {
		log.Fatalf("Failed to register reservation service handler. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}
