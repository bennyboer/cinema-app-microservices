package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/microtest"
	"github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/proto"
	s "github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user/service"
	"testing"
)

func registerServiceHandler(server server.Server) error {
	return proto.RegisterUserHandler(server, s.NewUserServiceHandler())
}

func initClient(serviceName string, clientService micro.Service) proto.UserService {
	return proto.NewUserService(serviceName, clientService.Client())
}

func TestUserService_Create(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		rsp, err := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		if err != nil {
			t.Fatalf("create request returned unexpected error. Error:\n%s", err.Error())
		}

		if rsp.CreatedId < 0 {
			t.Errorf("expected created user id to be non-negative")
		}
	})
}

func TestUserService_Create_EmptyName(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, err := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "",
			},
		})

		if err == nil {
			t.Fatalf("expected create request to be unsuccessful")
		}
	})
}

func TestUserService_Create_AlreadyExistant(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, _ = client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		_, err := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		if err == nil {
			t.Fatalf("expected create request to be unsuccessful")
		}
	})
}

func TestUserService_Read(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		createRsp, _ := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		id := createRsp.CreatedId

		rsp, err := client.Read(context.TODO(), &proto.ReadRequest{
			Id: id,
		})
		if err != nil {
			t.Fatalf("read request returned unexpected error. Error:\n%s", err.Error())
		}

		if rsp.Id != id {
			t.Errorf("expected id %d, got %d\n", id, rsp.Id)
		}

		if rsp.Data.Name != "Max Mustermann" {
			t.Errorf("expected name of the read user to be '%s' and not '%s'\n", "Max Mustermann", rsp.Data.Name)
		}
	})
}

func TestUserService_Read_Unsuccessful(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, err := client.Read(context.TODO(), &proto.ReadRequest{
			Id: 1,
		})

		if err == nil {
			t.Fatalf("expected read request to be unsuccessful")
		}
	})
}

func TestUserService_ReadAll(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, _ = client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		_, _ = client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Another Name",
			},
		})

		_, _ = client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Hello World",
			},
		})

		rsp, err := client.ReadAll(context.TODO(), &proto.ReadAllRequest{})
		if err != nil {
			t.Fatalf("read request returned unexpected error. Error:\n%s", err.Error())
		}

		if len(rsp.Ids) != 3 {
			t.Errorf("expected to find all 3 user ids, got %d", len(rsp.Ids))
		}

		if len(rsp.Dates) != 3 {
			t.Errorf("expected to find all 3 user dates, got %d", len(rsp.Ids))
		}
	})
}

func TestUserService_Update(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		createRsp, _ := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		id := createRsp.CreatedId

		_, err := client.Update(context.TODO(), &proto.UpdateRequest{
			Id: id,
			Data: &proto.UserData{
				Name: "Another Name",
			},
		})
		if err != nil {
			t.Fatalf("update request returned unexpected error. Error:\n%s", err.Error())
		}

		readRsp, _ := client.Read(context.TODO(), &proto.ReadRequest{
			Id: id,
		})

		if readRsp.Data.Name != "Another Name" {
			t.Errorf("expected update to actually change the name")
		}
	})
}

func TestUserService_Update_Unsuccessful(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, err := client.Update(context.TODO(), &proto.UpdateRequest{
			Id: 5,
			Data: &proto.UserData{
				Name: "Another Name",
			},
		})
		if err == nil {
			t.Fatalf("expected update request to be unsuccessful")
		}
	})
}

func TestUserService_Delete(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		createRsp, _ := client.Create(context.TODO(), &proto.CreateRequest{
			Data: &proto.UserData{
				Name: "Max Mustermann",
			},
		})

		id := createRsp.CreatedId

		_, err := client.Delete(context.TODO(), &proto.DeleteRequest{
			Id: id,
		})
		if err != nil {
			t.Fatalf("delete request returned unexpected error. Error:\n%s", err.Error())
		}

		_, err = client.Read(context.TODO(), &proto.ReadRequest{
			Id: id,
		})
		if err == nil {
			t.Errorf("expected delete to actually delete the user")
		}
	})
}

func TestUserService_Delete_Unsuccessful(t *testing.T) {
	microtest.TestServices(t, []microtest.ServiceSpec{
		{ServiceName: serviceName, HandlerRegistrationFunc: registerServiceHandler},
	}, func(t *testing.T, clientService micro.Service) {
		client := initClient(serviceName, clientService)

		_, err := client.Delete(context.TODO(), &proto.DeleteRequest{
			Id: 6,
		})
		if err == nil {
			t.Fatalf("expected delete request to be unsuccessful")
		}
	})
}
