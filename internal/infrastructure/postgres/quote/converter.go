package quote

import "github.com/ghost-circuit/word-of-wisdom/internal/domain/entity"

func convertQuoteModelToEntity(m Quote) entity.Quote {
	return entity.Quote{
		ID:     m.ID,
		Text:   m.Text,
		Author: m.Author,
	}
}
