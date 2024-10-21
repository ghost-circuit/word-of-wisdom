package client

import (
	"log/slog"
	"time"
)

type Printer struct {
	printerID int
}

func NewPrinter(id int) *Printer {
	return &Printer{
		printerID: id,
	}
}

func (p *Printer) PrintQuote(quote Quote) {
	slog.Info("Saved quote",
		slog.Int("printerID", p.printerID),
		slog.Int64("id", quote.ID),
		slog.String("text", quote.Text),
		slog.String("author", quote.Author),
	)
}

func (p *Printer) PrintError(msg string, err error) {
	slog.Error(msg, slog.Any("error", err))
}

func (p *Printer) PrintCurrentChallenge(challenge []byte, difficulty uint8) {
	slog.Info("Received challenge",
		slog.Int("printerID", p.printerID),
		slog.String("challenge", string(challenge)),
		slog.Int("difficulty", int(difficulty)),
	)
}

func (p *Printer) PrintSolution(challenge, solution []byte) {
	slog.Info("Received challenge",
		slog.Int("printerID", p.printerID),
		slog.String("challenge", string(challenge)),
		slog.String("solution", string(solution)),
	)
}

func (p *Printer) PrintFinishWork(duration time.Duration) {
	slog.Info("Finished work",
		slog.Int("printerID", p.printerID),
		slog.Duration("duration", duration),
	)
}
