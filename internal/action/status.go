package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

// Status fetches the current status of all eventing functions.
func Status(ctx context.Context, cluster *gocb.Cluster) (*gocb.EventingStatus, error) {
	statuses, err := cluster.EventingFunctions().FunctionsStatus(&gocb.EventingFunctionsStatusOptions{
		Timeout:       10 * time.Second,
		RetryStrategy: gocb.NewBestEffortRetryStrategy(nil),
		Context:       ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch functions statuses: %w", err)
	}

	return statuses, nil
}
