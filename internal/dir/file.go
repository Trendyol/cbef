package dir

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadFileFromPath reads file and parse into any model.
func ReadFileFromPath[T any](path string) (*T, error) {
	cfg, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var f T
	if err = json.Unmarshal(cfg, &f); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &f, nil
}
