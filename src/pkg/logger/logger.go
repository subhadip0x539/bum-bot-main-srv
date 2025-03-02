package logger

import (
	"os"

	"log/slog"
)

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
