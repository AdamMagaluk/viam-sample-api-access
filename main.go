package main

import (
	"context"
	"os"

	"github.com/edaniels/golog"
	apppb "go.viam.com/api/app/v1"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := golog.NewDebugLogger("test")
	ctx := context.Background()

	accessToken := os.Getenv("VIAM_ACCESS_TOKEN")
	if accessToken == "" {
		logger.Fatal("Must pass in VIAM_ACCESS_TOKEN, use ")
	}

	if len(os.Args) < 3 {
		logger.Fatal("Usage: <org_id> <location name>")
	}

	organizationID := os.Args[1]
	locationName := os.Args[2]

	conn, err := rpc.DialDirectGRPC(context.Background(), "app.viam.com:443", logger,
		rpc.WithDialDebug(),
		rpc.WithStaticAuthenticationMaterial(accessToken),
	)
	if err != nil {
		logger.Fatal(err)
	}

	client := apppb.NewAppServiceClient(conn)

	res, err := client.CreateLocation(ctx, &apppb.CreateLocationRequest{OrganizationId: organizationID, Name: locationName})
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("Created location %s", res.Location.Id)
}
