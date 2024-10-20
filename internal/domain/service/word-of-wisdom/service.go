package word_of_wisdom

import (
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/interfaces"
)

// WordOfWisdomService is a service for the WordOfWisdom domain.
type WordOfWisdomService struct {
	quoteRepo  interfaces.QuoteRepository
	powManager interfaces.ProofOfWorkManager
}

// NewWordOfWisdomService creates a new WordOfWisdomService.
func NewWordOfWisdomService(quoteRepo interfaces.QuoteRepository, powManager interfaces.ProofOfWorkManager) *WordOfWisdomService {
	return &WordOfWisdomService{
		quoteRepo:  quoteRepo,
		powManager: powManager,
	}
}
