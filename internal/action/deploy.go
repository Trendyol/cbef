package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

func Deploy(ctx context.Context, name string, cluster *gocb.Cluster) error {
	if err := cluster.EventingFunctions().DeployFunction(name, &gocb.DeployEventingFunctionOptions{
		Timeout:       10 * time.Second,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to deploy function: %w", err)
	}

	return nil
}
