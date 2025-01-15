package form

import "github.com/ghost-circuit/word-of-wisdom/internal/domain/consts"

// SubmitSolution is a form for submitting a solution to a challenge.
type SubmitSolution struct {
	Challenge []byte
	Solution  []byte
}

// Validate validates the SubmitSolution form.
func (s SubmitSolution) Validate() error {
	if len(s.Challenge) == 0 || len(s.Solution) == 0 {
		return consts.ErrEmptyChallengeOrSolution
	}

	return nil
}
