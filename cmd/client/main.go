package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/client"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/logger"
	pow_algorithm "github.com/alisher-baizhumanov/word-of-wisdom/pkg/pow-algorithm"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to open config file", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	cfg, err := client.NewConfig()
	if err != nil {
		return fmt.Errorf("main.run: %w", err)
	}

	logger.InitLogger(cfg.IsSugarLogger)
	slog.Debug("config file", slog.Any("cfg", cfg))

	var wg sync.WaitGroup

	for i := 0; i < cfg.CountWorker; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			grpcClient, errConn := client.NewClient(cfg.ServerAddr)
			if errConn != nil {
				slog.Error("error connection", slog.Any("error", errConn))

				return
			}

			service := client.NewService(
				pow_algorithm.NewProofOfWorkManager(0),
				grpcClient,
				i,
				int32(cfg.RequestsPerWorker),
			)

			if i%2 == 0 {
				service.ExecuteSequential()
			} else {
				service.ExecuteParallel()
			}
		}(i)
	}

	wg.Wait()

	return nil
}
