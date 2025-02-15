package utils

import (
	"log/slog"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/domain"
)

func LogEvent(err domain.Error) {
	switch err.Severity {
	case domain.SEVERITY_ERROR:
		slog.Error(err.Message)
	case domain.SEVERITY_WARNING:
		slog.Warn(err.Message)
	case domain.SEVERITY_SUCCESS:
		slog.Info(err.Message)
	default:
		slog.Info(err.Message)
	}
}
