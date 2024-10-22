package client

//go:generate mkdir -p ./mocks
//go:generate minimock -o ./mocks -s _minimock.go

// PowManager is an interface for the Proof of Work manager.
type PowManager interface {
	SolveCustomDifficulty(challenge []byte, difficulty uint8) ([]byte, error)
}

type Client interface {
	SubmitSolution(challenge []byte, solution []byte) (Quote, error)
	GetChallenge() ([]byte, uint8, error)
}
