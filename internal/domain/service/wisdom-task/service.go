package wisdom_task

import (
	"github.com/ghost-circuit/word-of-wisdom/internal/domain/interfaces"
)

// WisdomTaskService is a service for the WordOfWisdom domain.
type WisdomTaskService struct {
	quoteRepo  interfaces.QuoteRepository
	powManager interfaces.ProofOfWorkManager
}

// NewWisdomTaskService creates a new WisdomTaskService.
func NewWisdomTaskService(quoteRepo interfaces.QuoteRepository, powManager interfaces.ProofOfWorkManager) *WisdomTaskService {
	return &WisdomTaskService{
		quoteRepo:  quoteRepo,
		powManager: powManager,
	}
}
