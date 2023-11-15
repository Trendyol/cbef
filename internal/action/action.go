package action

import (
	"context"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

// Action represents eventing function actions.
type Action interface {
	Get(ctx context.Context, name string) (*gocb.EventingFunction, error)
	Pause(ctx context.Context, name string) error
	Upsert(ctx context.Context, f *model.Function) error
}

// Action represents eventing function Action instance.
type action struct {
	cluster *gocb.Cluster
	timeout time.Duration
}

// NewAction creates new eventing function Action instance.
func NewAction(cluster *gocb.Cluster, timeout time.Duration) Action {
	return &action{
		cluster: cluster,
		timeout: timeout,
	}
}
