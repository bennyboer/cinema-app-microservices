package main

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie/service"
	presentation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	"log"
)

const serviceName = "movie-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	err := proto.RegisterMovieHandler(service.Server(), s.NewMovieServiceHandler(
		&s.MovieServiceDependencies{
			PresentationService: func() presentation.PresentationService {
				return presentation.NewPresentationService("presentation-service", service.Client())
			},
		},
	))
	if err != nil {
		log.Fatalf("Failed to register movie service handler. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}
