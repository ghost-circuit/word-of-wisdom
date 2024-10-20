package pow_algorithm

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const (
	// ByteSize defines the number of bits in a byte. It's a constant used to convert between bits and bytes.
	ByteSize = 8

	// NonceSize specifies the size of the nonce in bytes. A nonce is a random or incremental number added to the input of the hash function.
	NonceSize = 16

	// MaxByteValue defines the maximum value of a byte, which is 255 (0xFF in hexadecimal).
	// It is used when applying bitwise operations for hash verification.
	MaxByteValue = 0xFF
)

// ProofOfWorkManager is a struct that holds the difficulty level and the length of the challenge.
// - difficulty: the number of leading bits that must be zero in a valid hash.
// - length: the number of random bytes in the challenge.
type ProofOfWorkManager struct {
	difficulty uint8
}

// NewProofOfWorkManager creates a new ProofOfWorkManager instance.
// - difficulty: the number of leading zero bits required for the solution.
func NewProofOfWorkManager(difficulty uint8) *ProofOfWorkManager {
	return &ProofOfWorkManager{
		difficulty: difficulty,
	}
}

// GenerateChallenge generates a random challenge of specified length using a secure random number generator.
// It returns the generated challenge (as a byte slice), the difficulty, and an error if the random generator fails.
// This challenge is used as input to the Proof of Work algorithm.
func (pow *ProofOfWorkManager) GenerateChallenge() ([]byte, uint8, error) {
	// Create a byte slice of the required length
	b := make([]byte, NonceSize)

	// Fill the byte slice with cryptographically secure random bytes
	_, err := rand.Read(b)
	if err != nil {
		// If the random read fails, return an error
		return nil, 0, err
	}

	// Return the generated challenge along with the difficulty
	return b, pow.difficulty, nil
}

// ValidateSolution checks whether the given solution is valid for the provided challenge.
// A solution is valid if the hash of the concatenated challenge and solution has the required number of leading zero bits (determined by the difficulty).
func (pow *ProofOfWorkManager) ValidateSolution(challenge, solution []byte) bool {
	// Combine the challenge and solution (typically the nonce) into one byte slice
	data := append(challenge, solution...)

	// Compute the SHA-256 hash of the combined data
	hash := sha256.Sum256(data)

	// fullBytes represents how many full bytes (8 bits each) we need to check for zero.
	fullBytes := int(pow.difficulty / ByteSize)

	// remainingBits represents the number of remaining bits to check after checking full bytes.
	remainingBits := pow.difficulty % ByteSize

	// Check all full bytes first. All full bytes must be zero for a valid solution.
	for i := 0; i < fullBytes; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	// Check remaining bits (if difficulty is not a multiple of 8).
	// We only need to check part of a byte if there are remaining bits to validate.
	if remainingBits > 0 {
		byteIdx := fullBytes                                        // Index of the byte containing the remaining bits
		bitMask := byte(MaxByteValue << (ByteSize - remainingBits)) // Mask to isolate the required bits

		// If the masked part of the byte isn't zero, the solution is invalid
		if (hash[byteIdx] & bitMask) != 0 {
			return false
		}
	}

	// If all checks pass, the solution is valid
	return true
}

// Solve tries to find a valid solution (nonce) for the given challenge by brute-forcing through possible nonces.
// It increments the nonce until a valid solution is found (i.e., the hash has the required number of leading zero bits).
// Returns the solution (nonce) or an error if no solution can be found (e.g., nonce space exhaustion).
func (pow *ProofOfWorkManager) Solve(challenge []byte) ([]byte, error) {
	var nonce uint64 // Start the nonce at 0

	// Infinite loop to brute-force nonces until a valid solution is found
	for {
		// Convert the current nonce value to a byte slice
		nonceBytes := uint64ToBytes(nonce)

		// Check if the current nonce is a valid solution
		if pow.ValidateSolution(challenge, nonceBytes) {
			return nonceBytes, nil // Return the nonce if it's valid
		}

		// Increment the nonce and check for overflow
		if nonce == ^uint64(0) { // If the nonce reaches its maximum value (overflow condition)
			return nil, fmt.Errorf("failed to find a valid solution, nonce space exhausted")
		}

		nonce++ // Increment the nonce for the next iteration
	}
}

// uint64ToBytes is a helper function that converts a uint64 number into a byte slice.
// This is necessary to represent the nonce as a byte array.
func uint64ToBytes(n uint64) []byte {
	// Create a byte slice of the size of the nonce (8 bytes)
	b := make([]byte, NonceSize)

	// Use big-endian encoding to convert the uint64 value to a byte slice
	binary.BigEndian.PutUint64(b, n)

	return b
}
