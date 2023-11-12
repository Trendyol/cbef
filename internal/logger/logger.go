package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
	}
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.Error(msg, args)
	os.Exit(1)
}
