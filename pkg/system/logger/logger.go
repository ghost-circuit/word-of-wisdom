package logger

import (
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
)

// InitLogger initializes the logger.
func InitLogger(isSugarLogger bool) {
	logger := getLogger(isSugarLogger)

	slog.SetDefault(logger)
}

func getLogger(isSugarLogger bool) *slog.Logger {
	if isSugarLogger {
		return slog.New(devslog.NewHandler(os.Stdout, &devslog.Options{
			HandlerOptions: &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
			NewLineAfterLog: true,
		}))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}
