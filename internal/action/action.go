package action

import (
	"context"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

// Action represents eventing function actions.
type Action interface {
	Upsert(ctx context.Context, f *model.Function) error
	Status(ctx context.Context) (*gocb.EventingStatus, error)
	Pause(ctx context.Context, name string) error
	Deploy(ctx context.Context, name string) error
	Undeploy(ctx context.Context, name string) error
	StopFunctions(ctx context.Context, prefix, excludedFunction string) (map[string]struct{}, error)
	WaitFunctionsProcesses(ctx context.Context, functions map[string]struct{}) error
}

// Action represents eventing function Action instance.
type action struct {
	cluster *gocb.Cluster
}

// NewAction creates new eventing function Action instance.
func NewAction(cluster *gocb.Cluster) Action {
	return &action{
		cluster: cluster,
	}
}
