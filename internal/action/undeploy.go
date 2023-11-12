package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

// Undeploy undeploys an eventing function.
func (a *action) Undeploy(ctx context.Context, name string) error {
	if err := a.cluster.EventingFunctions().UndeployFunction(name, &gocb.UndeployEventingFunctionOptions{
		Timeout:       10 * time.Second,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to undeploy function: %w", err)
	}

	return nil
}
