package quote

import "github.com/ghost-circuit/word-of-wisdom/pkg/adapter/postgres"

// Repository represents the quote repository.
type Repository struct {
	MasterConnection *postgres.ConnectionPool
}

// NewRepository initializes and returns a new quote repository with the given ConnectionPool.
func NewRepository(quoteRepo *postgres.ConnectionPool) *Repository {
	return &Repository{
		MasterConnection: quoteRepo,
	}
}
