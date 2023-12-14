package utils

import (
	"fmt"
	"strconv"

	"github.com/catness812/PAD-labs/user_svc/internal/pb"
	"github.com/gookit/slog"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func FindJournalService(serviceName string) (pb.JournalServiceClient, error) {
	client, err := createConsulClient()
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	serviceEntries, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		slog.Fatalf("Error querying Consul: %v", err)
		return nil, err
	}

	if len(serviceEntries) == 0 {
		slog.Fatalf("No healthy instances of %s found in Consul", serviceName)
		return nil, err
	}

	selectedService := serviceEntries[0]

	conn, err := grpc.Dial(selectedService.Service.Address+":"+strconv.Itoa(selectedService.Service.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Fatalf("Error connecting: %v", err)
		return nil, err
	}

	svcClient := pb.NewJournalServiceClient(conn)
	return svcClient, nil
}
