package quote

import (
	"context"
	"fmt"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/consts"
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/entity"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/adapter/postgres"
)

// GetRandomQuote returns a random quote.
func (r *Repository) GetRandomQuote(ctx context.Context) (entity.Quote, error) {
	var quote Quote

	query := postgres.Query{
		Name: "GetRandomQuote",
		QueryRaw: `
			SELECT id, text, author
			FROM quotes
			WHERE id = (
			  SELECT id
			  FROM quotes
			  OFFSET floor(random() * (SELECT count(1) FROM quotes))
			  LIMIT 1
			)`,
	}

	err := r.MasterConnection.ScanOne(ctx, &quote, query)
	if err != nil {
		return entity.Quote{}, fmt.Errorf("quote.Repository.GetRandomQuote: %w, %w", consts.ErrDatabase, err)
	}

	return convertQuoteModelToEntity(quote), nil
}
