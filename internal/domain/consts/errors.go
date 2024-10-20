package consts

import "errors"

var (
	// ErrDatabase represents a common database error.
	ErrDatabase = errors.New("common database error")

	// ErrEmptyChallengeOrSolution represents an error when challenge or solution is empty.
	ErrEmptyChallengeOrSolution = errors.New("empty challenge or solution")

	// ErrInvalidSolution represents an error when solution is invalid.
	ErrInvalidSolution = errors.New("invalid solution")
)
