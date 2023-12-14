package journalsvc

import (
	"github.com/catness812/PAD-labs/user_svc/internal/pb"
	"github.com/catness812/PAD-labs/user_svc/internal/utils"
	"github.com/gookit/slog"
)

func JournalServiceClient() pb.JournalServiceClient {
	client, err := utils.FindJournalService("journal-grpc-svc")
	if err != nil {
		slog.Fatalf("Error finding service: %v", err)
	}

	return client
}
