package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/client"
	powalgorithm "github.com/alisher-baizhumanov/word-of-wisdom/pkg/pow-algorithm"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/system/logger"
)

const (
	HandlerQuote byte = 0x01 // Constant for the "quote" handler
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("Failed to init the application", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := client.NewConfig()
	if err != nil {
		return fmt.Errorf("main.run: %w", err)
	}

	logger.InitLogger(cfg.IsSugarLogger)

	slog.Debug("Client configuration loaded", slog.Any("config", cfg))

	// create a ticker to control the request rate
	ticker := time.NewTicker(time.Second / time.Duration(cfg.RPS))
	defer ticker.Stop()

	// restrict the number of concurrent requests
	for i := 0; i < cfg.TotalRequests; i++ {
		<-ticker.C

		go func(requestNum int) {
			if errReq := sendRequest(cfg.ServerAddr, HandlerQuote); errReq != nil {
				slog.Error("Request failed", slog.Int("request_num", requestNum), slog.Any("error", errReq))

				return
			}

			slog.Info("Request succeeded", slog.Int("request_num", requestNum))
		}(i)
	}

	// wait for all requests to finish
	time.Sleep(time.Duration(cfg.TotalRequests/cfg.RPS) * time.Second)

	return nil
}

func sendRequest(serverAddr string, handlerID byte) error {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	buffer := make([]byte, 9) // 8 bytes for the challenge and 1 byte for the difficulty
	_, err = conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read challenge and difficulty: %w", err)
	}

	challenge := buffer[:8] // first 8 bytes are the challenge
	difficulty := buffer[8] // last byte is the difficulty

	powAlgo := powalgorithm.NewProofOfWorkManager(difficulty)
	solution, err := powAlgo.Solve(challenge)
	if err != nil {
		return err
	}

	var requestBuffer bytes.Buffer
	requestBuffer.Write(solution)

	// Write the handlerID (1 byte) to the buffer
	err = requestBuffer.WriteByte(handlerID)
	if err != nil {
		return fmt.Errorf("failed to write handler ID: %w", err)
	}

	// Send the binary request to the server
	_, err = conn.Write(requestBuffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	// Step 5: Read the response from the server
	response, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	slog.Info("Server response", slog.String("response", string(response)))

	return nil
}
