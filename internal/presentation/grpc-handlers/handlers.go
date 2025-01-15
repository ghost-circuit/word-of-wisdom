package grpc_handlers

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	wisdomtask "github.com/ghost-circuit/word-of-wisdom/internal/domain/service/wisdom-task"
	desc "github.com/ghost-circuit/word-of-wisdom/pkg/generated/wisdom"
)

// WordOfWisdomHandlers is a collection of handlers for the
type WordOfWisdomHandlers struct {
	desc.UnimplementedWordOfWisdomServiceServer

	service *wisdomtask.WisdomTaskService
}

// NewWordOfWisdomHandlers creates a new WordOfWisdomHandlers.
func NewWordOfWisdomHandlers(service *wisdomtask.WisdomTaskService) *WordOfWisdomHandlers {
	return &WordOfWisdomHandlers{
		service: service,
	}
}

// GetChallenge returns a new challenge.
func (h *WordOfWisdomHandlers) GetChallenge(ctx context.Context, _ *emptypb.Empty) (*desc.ChallengeResponse, error) {
	challenge, difficulty, err := h.service.GetChallenge(ctx)
	if err != nil {
		return nil, err
	}

	return convertChallengeResponse(challenge, difficulty), nil
}

// SubmitSolution submits a solution to the challenge.
func (h *WordOfWisdomHandlers) SubmitSolution(ctx context.Context, solution *desc.SolutionRequest) (*desc.SolutionResponse, error) {
	form := convertSolutionRequest(solution)
	if err := form.Validate(); err != nil {
		return nil, err
	}

	quote, err := h.service.SubmitSolution(ctx, form)
	if err != nil {
		return nil, err
	}

	return convertSolutionResponse(quote), nil
}
