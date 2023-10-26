package utils

import (
	"fmt"

	"github.com/gookit/slog"
	"github.com/hashicorp/consul/api"
)

func createConsulClient() (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = "consul:8500"
	return api.NewClient(config)
}

func RegisterService(port int) {
	client, err := createConsulClient()
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = fmt.Sprintf("user-grpc-svc-%d", port)
	registration.Name = "user-grpc-svc"
	registration.Address = "user-service"
	registration.Port = port

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	slog.Infof("User Service on port %d registered", port)
}
