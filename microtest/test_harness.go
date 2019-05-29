package microtest

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"os"
	"sync"
	"testing"
)

const testClientServiceName = "_test_client_service"

// Cached client service
var clientService micro.Service
// Cached services
var serviceList []micro.Service

// Specification for a micro service
type ServiceSpec struct {
	// Name of the micro service
	ServiceName string
	// Function registering the service handler for the passed server
	HandlerRegistrationFunc func(server server.Server) error
}

// Will setup all passed services for you.
func TestServices(
	t *testing.T,
	services []ServiceSpec,
	testFunc func(t *testing.T, clientService micro.Service),
) {
	// Reset command line arguments as service.Init() may throw an error otherwise
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = os.Args[:1]

	if serviceList == nil {
		// Start services to test
		serviceList := make([]micro.Service, 0, len(services))
		serviceChannel := make(chan micro.Service)
		var wg sync.WaitGroup
		for i := 0; i < len(services); i++ {
			spec := services[i]

			go startService(spec.ServiceName, t, spec.HandlerRegistrationFunc, wg, serviceChannel)
			serviceList = append(serviceList, <-serviceChannel)
		}
	}

	if clientService == nil {
		// Start test client service to test with
		clientService = micro.NewService(
			micro.Name(testClientServiceName),
		)
		clientService.Init()
	}

	testFunc(t, clientService)
}

func startService(
	name string,
	t *testing.T,
	registerHandlerFunc func(server server.Server) error,
	wg sync.WaitGroup,
	c chan micro.Service,
) {
	wg.Add(1)

	var service micro.Service
	service = micro.NewService(
		micro.Name(name),
		micro.AfterStart(func() error {
			c <- service
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Done()
			return nil
		}),
	)

	service.Init()

	err := registerHandlerFunc(service.Server())
	if err != nil {
		t.Fatalf("Registering service handler failed. Error:\n%s", err.Error())
	}

	if err := service.Run(); err != nil {
		t.Fatalf("Failed to start service. Error:\n%s", err.Error())
	}
}

func terminateService(t *testing.T, service micro.Service) {
	err := service.Server().Stop()
	if err != nil {
		t.Fatalf("Service could not be terminated properly. Error:\n%s", err.Error())
	}
}
