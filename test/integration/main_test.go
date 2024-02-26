package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"

	"github.com/kwilteam/kwil-db/test/integration"
	"github.com/kwilteam/kwil-db/test/utils"
)

func setupNetwork(ctx context.Context) *testcontainers.DockerNetwork {
	network, err := utils.CreateNetwork(ctx)
	if err != nil {
		panic(err)
	}

	integration.SetTestNetwork(network.ID)
	return network
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	defer setupNetwork(ctx).Remove(ctx)

	os.Exit(m.Run())
}
