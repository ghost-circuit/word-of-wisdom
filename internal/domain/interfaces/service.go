package interfaces

import (
	"context"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/entity"
)

type WordOfWisdomService interface {
	GetChallenge(ctx context.Context) ([]byte, uint8, error)
	SubmitSolution(ctx context.Context, solution []byte) (entity.Quote, error)
}
