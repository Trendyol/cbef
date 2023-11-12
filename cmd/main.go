package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/trendyol/cbef/internal/action"
	"github.com/trendyol/cbef/internal/logger"
	"github.com/trendyol/cbef/internal/model"
)

func main() {
	log := logger.New()

	var env model.Environment
	if err := env.Fill(); err != nil {
		log.Fatal("failed to fill environment variables", "error", err.Error())
	}

	var f model.Function
	cfg, err := os.ReadFile(env.ConfigFile)
	if err != nil {
		log.Fatal("failed to read config file", "error", err.Error())
	}

	if err = json.Unmarshal(cfg, &f); err != nil {
		log.Fatal("failed to parse config file", "error", err.Error())
	}

	cluster, err := gocb.Connect(f.Cluster.ConnectionString, gocb.ClusterOptions{
		Username: f.Cluster.User,
		Password: f.Cluster.Pass,
	})
	if err != nil {
		log.Fatal("failed to connect couchbase", "error", err.Error())
	}

	if _, err = cluster.Ping(&gocb.PingOptions{}); err != nil {
		log.Fatal("failed to ping couchbase", "error", err.Error())
	}

	code, err := os.ReadFile(env.FunctionFile)
	if err != nil {
		log.Fatal("failed to read function file", "error", err.Error())
	}

	f.Code = string(code)
	name := f.Name
	f.Name = fmt.Sprintf("%s-%s", f.Name, env.CommitSHA)

	ctx, cancel := context.WithTimeout(context.Background(), env.ExecutionTimeout)
	defer cancel()

	if err = action.Upsert(ctx, f, cluster); err != nil {
		log.Fatal("failed to create function", "error", err.Error(), "function", f.Name)
	}

	processes, err := stopFunctions(ctx, cluster, name, f.Name)
	if err != nil {
		log.Fatal("failed to stop functions", "error", err.Error())
	}

	if err = waitFunctionsProcesses(ctx, cluster, processes); err != nil {
		log.Fatal("failed to wait function processes", "error", err.Error())
	}

	if err = action.Deploy(ctx, f.Name, cluster); err != nil {
		log.Fatal("failed to deploy function", "error", err.Error(), "function", f.Name)
	}
}

func stopFunctions(ctx context.Context, cluster *gocb.Cluster, prefix, excludedFunction string) (map[string]struct{}, error) {
	statuses, err := action.Status(ctx, cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve functions statuses: %w", err)
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
			if err = action.Pause(ctx, fn.Name, cluster); err != nil {
				return nil, fmt.Errorf("failed to pause function: %w", err)
			}

			processes[fn.Name] = struct{}{}
		case gocb.EventingFunctionStateDeploying:
			if err = action.Undeploy(ctx, fn.Name, cluster); err != nil {
				return nil, fmt.Errorf("failed to undeploy function: %w", err)
			}

			processes[fn.Name] = struct{}{}
		}
	}

	return processes, nil
}

func waitFunctionsProcesses(ctx context.Context, cluster *gocb.Cluster, functions map[string]struct{}) error {
	t := time.Tick(500 * time.Millisecond)
	for range t {
		if len(functions) == 0 {
			break
		}

		statuses, err := action.Status(ctx, cluster)
		if err != nil {
			return fmt.Errorf("failed to retrieve functions statuses: %w", err)
		}

		for _, fn := range statuses.Functions {
			if _, ok := functions[fn.Name]; ok && (fn.Status == gocb.EventingFunctionStatePaused || fn.Status == gocb.EventingFunctionStateUndeployed) {
				delete(functions, fn.Name)
			}
		}
	}

	return nil
}
