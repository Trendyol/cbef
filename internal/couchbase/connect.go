package couchbase

import (
	"context"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/model"
)

func Connect(ctx context.Context, cfg model.Cluster) (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(cfg.ConnectionString, gocb.ClusterOptions{
		Username: cfg.User,
		Password: cfg.Pass,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect couchbase: %s", err.Error())
	}

	if _, err = cluster.Ping(&gocb.PingOptions{Context: ctx}); err != nil {
		return nil, fmt.Errorf("failed to ping couchbase: %s", err.Error())
	}

	return cluster, nil
}
