package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

// WaitFunctionsProcesses waits processes for provided eventing functions.
func (a *action) WaitFunctionsProcesses(ctx context.Context, functions map[string]struct{}) error {
	t := time.Tick(500 * time.Millisecond)
	for range t {
		if len(functions) == 0 {
			break
		}

		statuses, err := a.Status(ctx)
		if err != nil {
			return fmt.Errorf("failed to fetch functions statuses: %w", err)
		}

		for _, fn := range statuses.Functions {
			if _, ok := functions[fn.Name]; ok && (fn.Status == gocb.EventingFunctionStatePaused || fn.Status == gocb.EventingFunctionStateUndeployed) {
				delete(functions, fn.Name)
			}
		}
	}

	return nil
}
