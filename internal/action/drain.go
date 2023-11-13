package action

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

// DrainFunctions drains eventing functions.
func (a *action) DrainFunctions(ctx context.Context, functions map[string]struct{}) error {
	t := time.Tick(500 * time.Millisecond)
	for range t {
		if len(functions) == 0 {
			break
		}

		statuses, err := a.Status(ctx)
		if err != nil {
			return fmt.Errorf("failed to fetch functions statuses: %w", err)
		}

		for name := range functions {
			var found bool
			for _, f := range statuses.Functions {
				if f.Name != name {
					continue
				}

				if f.Status == gocb.EventingFunctionStateUndeployed {
					if err = a.Drop(ctx, f.Name); err != nil {
						return err
					}
				}

				if f.Status == gocb.EventingFunctionStatePaused {
					if err = a.Undeploy(ctx, f.Name); err != nil {
						return err
					}
				}

				found = true
			}

			if !found {
				delete(functions, name)
			}
		}
	}

	return nil
}
