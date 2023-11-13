package action

import (
	"context"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

// Drop drops an eventing function.
func (a *action) Drop(ctx context.Context, name string) error {
	if err := a.cluster.EventingFunctions().DropFunction(name, &gocb.DropEventingFunctionOptions{
		Timeout:       a.timeout,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	}); err != nil {
		return fmt.Errorf("failed to drop function: %w", err)
	}

	return nil
}
