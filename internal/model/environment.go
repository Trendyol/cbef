package model

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Environment represents application configurations.
type Environment struct {
	ConfigFile       string
	FunctionFile     string
	ExecutionTimeout time.Duration
	CommitAuthor     string
}

// NewEnvironment creates a Environment struct via environment vriables.
func NewEnvironment() (*Environment, error) {
	e := &Environment{}
	if err := e.fill(); err != nil {
		return nil, fmt.Errorf("fill environment variables: %s", err.Error())
	}

	if err := e.validate(); err != nil {
		return nil, fmt.Errorf("validate environment variables: %s", err.Error())
	}

	return e, nil
}

// fill fills and validates Environment struct via environment variables.
func (e *Environment) fill() error {
	e.ConfigFile = os.Getenv("CONFIG_FILE")
	e.FunctionFile = os.Getenv("FUNCTION_FILE")
	e.CommitAuthor = os.Getenv("CI_COMMIT_AUTHOR")
	timeout := os.Getenv("EXECUTION_TIMEOUT")

	if len(timeout) == 0 {
		timeout = "3m"
	}

	t, err := time.ParseDuration(timeout)
	if err != nil {
		return fmt.Errorf("failed to parse execution timeout duration: %w", err)
	}

	e.ExecutionTimeout = t

	return nil
}

// validate validates provided environment variable requirements.
func (e *Environment) validate() error {
	if len(e.ConfigFile) == 0 {
		return errors.New("environment variable 'CONFIG_FILE' does not provided")
	}

	if len(e.FunctionFile) == 0 {
		return errors.New("environment variable 'FUNCTION_FILE' does not provided")
	}

	if len(e.CommitAuthor) == 0 {
		return errors.New("environment variable 'CI_COMMIT_AUTHOR' does not provided")
	}

	return nil
}
