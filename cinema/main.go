package main

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/proto"
	impl "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/service"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"log"
)

const serviceName = "cinema-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	err := proto.RegisterCinemaHandler(service.Server(), impl.NewCinemaServiceHandler(
		&impl.CinemaServiceDependencies{
			PresentationService: func() presentation.PresentationService {
				return presentation.NewPresentationService("presentation-service", service.Client())
			},
		},
	))

	if err != nil {
		log.Fatalf("Error while trying to register dependencies. Error: \n\t%s", err.Error())
	}

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n\t%s", err.Error())
	}
}
