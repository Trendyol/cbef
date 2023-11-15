package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/trendyol/cbef/internal/model"
)

// Parse reads config file and parse into function model.
func Parse(path string) (*model.Function, error) {
	cfg, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var f model.Function
	if err = json.Unmarshal(cfg, &f); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	f.Cluster.ConnectionString = model.FillStringFromEnvironment(f.Cluster.ConnectionString)
	f.Cluster.User = model.FillStringFromEnvironment(f.Cluster.User)
	f.Cluster.Pass = model.FillStringFromEnvironment(f.Cluster.Pass)

	return &f, nil
}
