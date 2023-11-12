package action

import (
	"context"
	"fmt"
	"strings"

	"github.com/couchbase/gocb/v2"
)

// StopFunctions stops eventing functions by given prefix.
func (a *action) StopFunctions(ctx context.Context, prefix, excludedFunction string) (map[string]struct{}, error) {
	statuses, err := a.Status(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch functions statuses: %w", err)
	}

	processes := make(map[string]struct{})
	for _, fn := range statuses.Functions {
		if fn.Name == excludedFunction {
			continue
		}

		if !strings.Contains(fn.Name, fmt.Sprintf("%s-", prefix)) {
			continue
		}

		switch fn.Status {
		case gocb.EventingFunctionStatePausing:
			processes[fn.Name] = struct{}{}
		case gocb.EventingFunctionStateDeployed:
			if err = a.Pause(ctx, fn.Name); err != nil {
				return nil, fmt.Errorf("failed to pause function: %w", err)
			}

			processes[fn.Name] = struct{}{}
		case gocb.EventingFunctionStateDeploying:
			if err = a.Undeploy(ctx, fn.Name); err != nil {
				return nil, fmt.Errorf("failed to undeploy function: %w", err)
			}

			processes[fn.Name] = struct{}{}
		}
	}

	return processes, nil
}
