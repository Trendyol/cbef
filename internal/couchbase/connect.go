package couchbase

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

// Connect connects an couchbase cluster.
func Connect(ctx context.Context, timeout time.Duration, cfg model.Cluster) (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(cfg.ConnectionString, gocb.ClusterOptions{
		Username: cfg.User,
		Password: cfg.Pass,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect couchbase: %s", err.Error())
	}

	res, err := cluster.Ping(&gocb.PingOptions{Timeout: timeout, Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to ping couchbase: %s", err.Error())
	}

	for _, reports := range res.Services {
		for _, report := range reports {
			if len(report.Error) > 0 {
				return nil, fmt.Errorf("failed to ping couchbase: %s", report.Error)
			}
		}
	}

	return cluster, nil
}
