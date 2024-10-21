package client

import (
	"sync"
	"sync/atomic"
	"time"

	powalgorithm "github.com/alisher-baizhumanov/word-of-wisdom/pkg/pow-algorithm"
)

// Service
type Service struct {
	powManager   *powalgorithm.ProofOfWorkManager
	client       *GrpcClient
	printer      *Printer
	requestCount int32
}

func NewService(manager *powalgorithm.ProofOfWorkManager, client *GrpcClient, printer *Printer, requestCount int32) *Service {
	return &Service{
		powManager:   manager,
		client:       client,
		printer:      printer,
		requestCount: requestCount,
	}
}

func (s *Service) ExecuteSequential() {
	start := time.Now()

	defer func() {
		end := time.Now()

		s.printer.PrintFinishWork(end.Sub(start))
	}()

	flag := true

	for flag {
		flag = s.executeOne()
	}
}

func (s *Service) ExecuteParallel() {
	start := time.Now()

	defer func() {
		end := time.Now()

		s.printer.PrintFinishWork(end.Sub(start))
	}()

	var wg sync.WaitGroup

	channel := make(chan bool, 1)
	channel <- true
	defer close(channel)

	for flag := <-channel; flag; {
		wg.Add(1)

		go func() {
			defer wg.Done()

			channel <- s.executeOne()
		}()
	}

	wg.Wait()
}

func (s *Service) executeOne() bool {
	remainingCount := atomic.AddInt32(&s.requestCount, -1)

	if remainingCount < 0 {
		atomic.AddInt32(&s.requestCount, 1)

		return false
	}

	s.execute()

	return true
}

func (s *Service) execute() {
	challenge, difficulty, err := s.client.GetChallenge()
	if err != nil {
		s.printer.PrintError("failed to get challenge", err)

		return
	}

	s.printer.PrintCurrentChallenge(challenge, difficulty)

	solution, err := s.powManager.SolveCustomDifficulty(challenge, difficulty)
	if err != nil {
		s.printer.PrintError("failed to find solution", err)

		return
	}

	s.printer.PrintSolution(challenge, solution)

	quote, err := s.client.SubmitSolution(challenge, solution)
	if err != nil {
		s.printer.PrintError("failed to submit solution", err)

		return
	}

	s.printer.PrintQuote(quote)
}
