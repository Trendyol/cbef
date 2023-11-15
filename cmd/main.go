package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/trendyol/cbef/internal/action"
	"github.com/trendyol/cbef/internal/config"
	"github.com/trendyol/cbef/internal/couchbase"
	"github.com/trendyol/cbef/internal/logger"
	"github.com/trendyol/cbef/internal/model"
)

const (
	connectTimeout   = 5 * time.Second
	processTimeout   = 10 * time.Second
	codeAuditComment = "// Created by %s on %s via github.com/trendyol/cbef.\n\n%s"
)

func main() {
	log := logger.New()

	env, err := model.NewEnvironment()
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := config.Parse(env.ConfigFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), env.ExecutionTimeout)
	defer cancel()

	cluster, err := couchbase.Connect(ctx, connectTimeout, f.Cluster)
	if err != nil {
		log.Fatal(err.Error()) //nolint:gocritic
	}

	code, err := os.ReadFile(env.FunctionFile)
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
		log.Fatal("failed to upsert function", "error", err.Error(), "function", f.Name)
	}
}
