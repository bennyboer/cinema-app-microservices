package main

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/service"
	"log"
)

const serviceName = "presentation-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	err := proto.RegisterPresentationHandler(service.Server(), s.NewPresentationServiceHandler())
	if err != nil {
		log.Fatalf("Failed to register presentation service handler. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}
