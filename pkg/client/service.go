package client

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	powalgorithm "github.com/alisher-baizhumanov/word-of-wisdom/pkg/pow-algorithm"
)

// Service represents a client service.
type Service struct {
	serviceID  int
	powManager *powalgorithm.ProofOfWorkManager
	client     *GrpcClient

	requestCount int32
}

func NewService(manager *powalgorithm.ProofOfWorkManager, client *GrpcClient, serviceID int, requestCount int32) *Service {
	return &Service{
		powManager:   manager,
		client:       client,
		serviceID:    serviceID,
		requestCount: requestCount,
	}
}

func (s *Service) ExecuteSequential() {
	start := time.Now()
	counter := 0

	defer func() {
		end := time.Now()

		slog.Info("Finished sequential work",
			slog.Int("serviceID", s.serviceID),
			slog.Duration("duration", end.Sub(start)),
			slog.Int("count", counter),
		)
	}()

	flag := true

	counter -= 1
	for flag {
		flag = s.executeOne()
		counter++
	}
}

func (s *Service) ExecuteParallel() {
	start := time.Now()
	counter := 0

	defer func() {
		end := time.Now()
		slog.Info("Finished parallel work",
			slog.Int("serviceID", s.serviceID),
			slog.Duration("duration", end.Sub(start)),
			slog.Int("count", counter),
		)
	}()

	var wg sync.WaitGroup
	channel := make(chan bool, s.requestCount)

	for i := int32(0); i < s.requestCount; i++ {
		counter++
		wg.Add(1)

		go func() {
			defer wg.Done()

			if s.executeOne() {
				channel <- true
			} else {
				channel <- false
			}
		}()
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(channel)
	}()

	// Process the results
	for flag := range channel {
		if !flag {
			break
		}
	}

}

func (s *Service) executeOne() bool {
	remainingCount := atomic.AddInt32(&s.requestCount, -1)

	if remainingCount < 0 {
		atomic.AddInt32(&s.requestCount, 1)

		return false
	}

	taskID, err := uuid.NewV7()
	if err != nil {
		slog.Error("failed to generate uuid",
			slog.Int("serviceID", s.serviceID),
			slog.Any("error", err),
		)
	}

	if err = s.execute(taskID); err != nil {
		slog.Error("failed to execute",
			slog.Any("error", err),
			slog.Int("serviceID", s.serviceID),
			slog.String("taskID", taskID.String()),
		)
	}

	return true
}

func (s *Service) execute(taskID uuid.UUID) error {
	challenge, difficulty, err := s.client.GetChallenge()
	if err != nil {
		return fmt.Errorf("client.Service.execute, get challenge: %w", err)
	}

	slog.Info("Received challenge",
		slog.Int("serviceID", s.serviceID),
		slog.String("challenge", hex.EncodeToString(challenge)),
		slog.Int("difficulty", int(difficulty)),
		slog.String("taskID", taskID.String()),
	)

	solution, err := s.powManager.SolveCustomDifficulty(challenge, difficulty)
	if err != nil {
		return fmt.Errorf("client.Service.execute, find challenge: %w", err)
	}

	slog.Info("Found solution",
		slog.Int("serviceID", s.serviceID),
		slog.String("challenge", hex.EncodeToString(challenge)),
		slog.String("solution", hex.EncodeToString(solution)),
	)

	quote, err := s.client.SubmitSolution(challenge, solution)
	if err != nil {
		return fmt.Errorf("client.Service.execute, submit solution: %w", err)

	}

	slog.Info("Saved quote",
		slog.Int("serviceID", s.serviceID),
		slog.Int64("id", quote.ID),
		slog.String("text", quote.Text),
		slog.String("author", quote.Author),
	)

	return nil
}
