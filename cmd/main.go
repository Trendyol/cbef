package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/trendyol/cbef/internal/action"
	"github.com/trendyol/cbef/internal/config"
	"github.com/trendyol/cbef/internal/couchbase"
	"github.com/trendyol/cbef/internal/logger"
	"github.com/trendyol/cbef/internal/model"
)

const (
	executionTimeout = 3 * time.Minute
	processTimeout   = 10 * time.Second
	codeAuditComment = "// Created by %s on %s via github.com/trendyol/cbef.\n\n%s"
)

func main() {
	log := logger.New()

	env, err := model.NewEnvironment()
	if err != nil {
		log.Fatal("failed to load environment", "error", err.Error())
	}

	f, err := config.Parse(env.ConfigFile)
	if err != nil {
		log.Fatal("failed to parse config file", "error", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	cluster, err := couchbase.Connect(ctx, processTimeout, f.Cluster)
	if err != nil {
		log.Fatal("failed to connect couchbase", "error", err.Error()) //nolint:gocritic
	}

	code, err := os.ReadFile(filepath.Join(filepath.Dir(env.ConfigFile), f.Name+".js"))
	if err != nil {
		log.Fatal("failed to read function file", "error", err.Error())
	}

	f.Code = fmt.Sprintf(codeAuditComment, env.CommitAuthor, time.Now().Format(time.DateTime), code)

	act := action.NewAction(cluster, processTimeout)

	fn, err := act.Get(ctx, f.Name)
	if err != nil {
		log.Fatal("failed to get function", "error", err.Error())
	}

	if fn != nil && fn.Settings.ProcessingStatus {
		if err = act.Pause(ctx, f.Name); err != nil {
			log.Fatal("failed to pause function", "error", err.Error())
		}
	}

	if err = act.Upsert(ctx, f); err != nil {
		log.Fatal("failed to upsert function", "error", err.Error())
	}
}
