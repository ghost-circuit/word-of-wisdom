package wisdom_task

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ghost-circuit/word-of-wisdom/internal/domain/consts"
	"github.com/ghost-circuit/word-of-wisdom/internal/domain/entity"
	"github.com/ghost-circuit/word-of-wisdom/internal/domain/form"
)

// GetChallenge returns a challenge for the user to solve.
func (w *WisdomTaskService) GetChallenge(_ context.Context) ([]byte, uint8, error) {
	challenge, difficulty, err := w.powManager.GenerateChallenge()
	if err != nil {
		slog.Error("failed to get challenge", slog.Any("error", err))

		return nil, 0, fmt.Errorf("service.WisdomTaskService.GetChallenge: %w", err)
	}

	return challenge, difficulty, nil
}

// SubmitSolution submits a solution to the challenge and returns a quote.
func (w *WisdomTaskService) SubmitSolution(ctx context.Context, solutionSubmit form.SubmitSolution) (entity.Quote, error) {
	ok := w.powManager.ValidateSolution(solutionSubmit.Challenge, solutionSubmit.Solution)
	if !ok {
		slog.Info("solution is invalid")

		return entity.Quote{}, consts.ErrInvalidSolution
	}

	quote, err := w.quoteRepo.GetRandomQuote(ctx)
	if err != nil {
		slog.Error("failed to get random quote", slog.Any("error", err))

		return entity.Quote{}, fmt.Errorf("service.WisdomTaskService.SubmitSolution: %w", err)
	}

	return quote, nil
}
