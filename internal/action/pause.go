package action

import (
	"context"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

// Pause pauses an eventing function.
func (a *action) Pause(ctx context.Context, name string) error {
	if err := a.cluster.EventingFunctions().PauseFunction(name, &gocb.PauseEventingFunctionOptions{
		Timeout:       a.timeout,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to pause function: %w", err)
	}

	return nil
}
