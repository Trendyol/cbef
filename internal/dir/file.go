package dir

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFileFromPath[T any](path string) (*T, error) {
	cfg, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err.Error())
	}

	var f T
	if err = json.Unmarshal(cfg, &f); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %s", err.Error())
	}

	return &f, nil
}
