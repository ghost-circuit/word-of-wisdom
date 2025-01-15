package interfaces

import (
	"context"

	"github.com/ghost-circuit/word-of-wisdom/internal/domain/entity"
)

// QuoteRepository is an interface for interacting with the quote repository.
type QuoteRepository interface {
	GetRandomQuote(ctx context.Context) (entity.Quote, error)
}
