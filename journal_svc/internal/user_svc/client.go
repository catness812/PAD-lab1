package usersvc

import (
	"github.com/catness812/PAD-labs/journal_svc/internal/pb"
	"github.com/catness812/PAD-labs/journal_svc/internal/utils"
	"github.com/gookit/slog"
)

func UserServiceClient() pb.UserServiceClient {
	client, err := utils.FindUserService("user-grpc-svc")
	if err != nil {
		slog.Fatalf("Error finding service: %v", err)
	}

	return client
}
