package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/trendyol/cbef/internal/action"
	"github.com/trendyol/cbef/internal/couchbase"
	"github.com/trendyol/cbef/internal/dir"
	"github.com/trendyol/cbef/internal/logger"
	"github.com/trendyol/cbef/internal/model"
)

func main() {
	log := logger.New()

	env, err := model.NewEnvironment()
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := dir.ReadFileFromPath[model.Function](env.ConfigFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), env.ExecutionTimeout)
	defer cancel()

	cluster, err := couchbase.Connect(ctx, f.Cluster)
	if err != nil {
		log.Fatal(err.Error())
	}

	code, err := os.ReadFile(env.FunctionFile)
	if err != nil {
		log.Fatal("failed to read function file", "error", err.Error())
	}

	f.Code = fmt.Sprintf("// Created by %s on %s via github.com/trendyol/cbef.\n\n%s", env.CommitAuthor, time.Now().Format(time.DateTime), code)
	name := f.Name
	f.Name = fmt.Sprintf("%s-%s", f.Name, env.CommitSHA)

	act := action.NewAction(cluster)

	if err = act.Upsert(ctx, f); err != nil {
		log.Fatal("failed to create function", "error", err.Error(), "function", f.Name)
	}

	processes, drainableFunctions, err := act.StopFunctions(ctx, name, f.Name)
	if err != nil {
		log.Fatal("failed to stop functions", "error", err.Error())
	}

	if err = act.WaitFunctionsProcesses(ctx, processes); err != nil {
		log.Fatal("failed to wait functions processes", "error", err.Error())
	}

	if err = act.Deploy(ctx, f.Name); err != nil {
		log.Fatal("failed to deploy function", "error", err.Error(), "function", f.Name)
	}

	if err = act.DrainFunctions(ctx, drainableFunctions); err != nil {
		log.Fatal("failed to drain functions", "error", err.Error())
	}
}
