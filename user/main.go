package main

import (
	"github.com/micro/go-micro"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/service"
	"log"
)

const serviceName = "user-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	err := proto.RegisterUserHandler(service.Server(), s.NewUserServiceHandler(
		&s.UserServiceDependencies{
			ReservationService: func() reservation.ReservationService {
				return reservation.NewReservationService("reservation-service", service.Client())
			},
		},
	))
	if err != nil {
		log.Fatalf("Failed to register user service handler. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}
