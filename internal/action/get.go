package action

import (
	"context"
	"errors"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

// Get gets an eventing function.
func (a *action) Get(ctx context.Context, name string) (*gocb.EventingFunction, error) {
	f, err := a.cluster.EventingFunctions().GetFunction(name, &gocb.GetEventingFunctionOptions{
		Timeout:       a.timeout,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	})
	if err != nil && !errors.Is(err, gocb.ErrEventingFunctionNotFound) {
		return nil, fmt.Errorf("failed to get function: %w", err)
	}

	return f, nil
}
