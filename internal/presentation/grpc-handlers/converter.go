package grpc_handlers

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/consts"
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/entity"
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/form"
	desc "github.com/alisher-baizhumanov/word-of-wisdom/pkg/generated/wisdom"
)

func convertError(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := status.FromError(err); ok {
		return err
	}

	switch {
	case errors.Is(err, consts.ErrDatabase):
		return status.Error(codes.Internal, "")

	case errors.Is(err, consts.ErrEmptyChallengeOrSolution):
		return status.Error(codes.InvalidArgument, consts.ErrEmptyChallengeOrSolution.Error())
	case errors.Is(err, consts.ErrInvalidSolution):
		return status.Error(codes.InvalidArgument, consts.ErrInvalidSolution.Error())

	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

func convertChallengeResponse(challenge []byte, difficulty uint8) *desc.ChallengeResponse {
	return &desc.ChallengeResponse{
		Challenge:  challenge,
		Difficulty: uint32(difficulty),
	}
}

func convertSolutionResponse(quote entity.Quote) *desc.SolutionResponse {
	return &desc.SolutionResponse{
		Quote: &desc.Quote{
			Text:   quote.Text,
			Author: quote.Author,
			Id:     quote.ID,
		},
	}
}

func convertSolutionRequest(solution *desc.SolutionRequest) form.SubmitSolution {
	return form.SubmitSolution{
		Challenge: solution.GetChallenge(),
		Solution:  solution.GetSolution(),
	}
}
