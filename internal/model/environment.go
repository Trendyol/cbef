package model

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Environment struct {
	ConfigFile       string
	FunctionFile     string
	CommitSHA        string
	ExecutionTimeout time.Duration
}

func (e *Environment) Fill() error {
	e.ConfigFile = os.Getenv("CONFIG_FILE")
	e.FunctionFile = os.Getenv("FUNCTION_FILE")
	e.CommitSHA = os.Getenv("CI_COMMIT_SHORT_SHA")
	timeout := os.Getenv("EXECUTION_TIMEOUT")
	if len(timeout) == 0 {
		timeout = "3m"
	}

	t, err := time.ParseDuration(timeout)
	if err != nil {
		return fmt.Errorf("failed to parse execution timeout duration: %w", err)
	}

	e.ExecutionTimeout = t

	return e.validate()
}

func (e *Environment) validate() error {
	if len(e.ConfigFile) == 0 {
		return errors.New("environment variable 'CONFIG_FILE' does not provided")
	}

	if len(e.FunctionFile) == 0 {
		return errors.New("environment variable 'FUNCTION_FILE' does not provided")
	}

	if len(e.CommitSHA) == 0 {
		return errors.New("environment variable 'CI_COMMIT_SHORT_SHA' does not provided")
	}

	return nil
}
