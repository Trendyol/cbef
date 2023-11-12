package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

func Undeploy(ctx context.Context, name string, cluster *gocb.Cluster) error {
	if err := cluster.EventingFunctions().UndeployFunction(name, &gocb.UndeployEventingFunctionOptions{
		Timeout:       10 * time.Second,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to undeploy function: %w", err)
	}

	return nil
}
