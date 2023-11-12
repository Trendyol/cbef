package logger

import (
	"log/slog"
	"os"
)

// Logger records structured log informations
type Logger struct {
	*slog.Logger
}

// New creates a logger instance with custom logger functions.
func New() *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
	}
}

// Fatal logs fatal reasons with exit.
func (l *Logger) Fatal(msg string, args ...any) {
	l.Error(msg, args)
	os.Exit(1)
}
