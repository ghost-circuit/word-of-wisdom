package grpc_handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/interfaces"
	desc "github.com/alisher-baizhumanov/word-of-wisdom/pkg/generated/wisdom"
)

type WordOfWisdomHandlers struct {
	desc.UnimplementedWordOfWisdomServiceServer

	service interfaces.WordOfWisdomService
}

func (h *WordOfWisdomHandlers) GetChallenge(ctx context.Context, _ *emptypb.Empty) (*desc.ChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenge not implemented")
}

func (h *WordOfWisdomHandlers) SubmitSolution(ctx context.Context, solution *desc.SolutionRequest) (*desc.SolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSolution not implemented")
}
