package pow_algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateChallenge(t *testing.T) {
	pow := NewProofOfWorkManager(20)

	challenge1, _, err1 := pow.GenerateChallenge()
	challenge2, _, err2 := pow.GenerateChallenge()

	require.NoError(t, err1)
	require.NoError(t, err2)

	require.Len(t, challenge1, NonceSize)
	require.Len(t, challenge2, NonceSize)

	assert.NotEqual(t, challenge1, challenge2, "Expected challenges to be different")
}

func TestValidateSolution_Valid(t *testing.T) {
	pow := NewProofOfWorkManager(20)

	challenge, _, err := pow.GenerateChallenge()
	require.NoError(t, err)

	solution, err := pow.Solve(challenge)
	require.NoError(t, err)

	flag := pow.ValidateSolution(challenge, solution)
	assert.True(t, flag)
}

func TestValidateSolution_Invalid(t *testing.T) {
	pow := NewProofOfWorkManager(20)
	challenge, _, err := pow.GenerateChallenge()
	require.NoError(t, err)

	// Manually create an invalid solution by using a wrong nonce
	invalidSolution := make([]byte, NonceSize)
	for i := range invalidSolution {
		invalidSolution[i] = 255 // invalid random data
	}

	flag := pow.ValidateSolution(challenge, invalidSolution)
	assert.False(t, flag)
}

func TestSolve(t *testing.T) {
	pow := NewProofOfWorkManager(20)
	challenge, _, err := pow.GenerateChallenge()
	require.NoError(t, err)

	solution, err := pow.Solve(challenge)
	require.NoError(t, err)

	flag := pow.ValidateSolution(challenge, solution)
	assert.True(t, flag)
}
