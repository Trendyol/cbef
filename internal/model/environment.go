package model

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Environment represents application configurations.
type Environment struct {
	ConfigFile   string
	CommitAuthor string
}

// NewEnvironment creates a Environment struct via environment vriables.
func NewEnvironment() (*Environment, error) {
	var e Environment
	e.ConfigFile = os.Getenv("CONFIG_FILE")
	e.CommitAuthor = os.Getenv("CI_COMMIT_AUTHOR")

	if err := e.validate(); err != nil {
		return nil, fmt.Errorf("validate environment variables: %s", err.Error())
	}

	return &e, nil
}

// validate validates provided environment variable requirements.
func (e *Environment) validate() error {
	if len(e.ConfigFile) == 0 {
		return errors.New("environment variable 'CONFIG_FILE' does not provided")
	}

	if len(e.CommitAuthor) == 0 {
		return errors.New("environment variable 'CI_COMMIT_AUTHOR' does not provided")
	}

	return nil
}

// FillStringFromEnvironment converts string to environment variable.
func FillStringFromEnvironment(s string) string {
	if strings.HasPrefix(s, "{{") && strings.HasSuffix(s, "}}") {
		s = strings.TrimPrefix(s, "{{")
		s = strings.TrimSuffix(s, "}}")
		s = os.Getenv(s)

		return s
	}

	return s
}
