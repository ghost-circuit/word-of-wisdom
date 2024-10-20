package interfaces

// ProofOfWorkManager is an interface for proof of work manager.
type ProofOfWorkManager interface {
	GenerateChallenge() ([]byte, uint8, error)
	ValidateSolution(challenge []byte, nonce []byte) bool
}
