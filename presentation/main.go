package presentation

import (
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/service"
	reservation "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation/proto"
	"log"
)

const serviceName = "presentation-service"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	err := proto.RegisterPresentationHandler(service.Server(), s.NewPresentationServiceHandler(
		&s.PresentationServiceDependencies{
			ReservationService: func() reservation.ReservationService {
				return reservation.NewReservationService("reservation-service", service.Client())
			},
		},
	))
	if err != nil {
		log.Fatalf("Failed to register presentation service handler. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}
