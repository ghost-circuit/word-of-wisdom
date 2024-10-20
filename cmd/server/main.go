package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/application/config"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/system/logger"
	waitsignal "github.com/alisher-baizhumanov/word-of-wisdom/pkg/system/wait-signal"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("Failed to init the application", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	// Load server configuration
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("main.run: %w", err)
	}

	logger.InitLogger(cfg.IsSugarLogger)

	slog.Debug("Server configuration loaded", slog.Any("config", cfg))

	// Create a parent context for the server

	slog.Info("Server is running...")

	// Wait for termination signal
	signal := waitsignal.WaitSignal()
	slog.Info("Shutdown signal received", slog.String("signal", signal.String()))

	// Gracefully shut down the server

	slog.Info("Server shutdown completed")

	return nil
}
