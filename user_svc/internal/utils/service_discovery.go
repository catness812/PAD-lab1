package utils

import (
	"github.com/catness812/PAD-lab1/user_svc/internal/config"
	"github.com/gookit/slog"
	"github.com/hashicorp/consul/api"
)

func createConsulClient() (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	return api.NewClient(config)
}

func RegisterService() {
	client, err := createConsulClient()
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "user-grpc-svc"
	registration.Name = "user-grpc-svc"
	registration.Address = config.Cfg.Host
	registration.Port = config.Cfg.GrpcPort

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		slog.Error(err)
		panic(err)
	}
}
