package client_test

import (
	"testing"

	"github.com/gojuno/minimock/v3"

	"github.com/ghost-circuit/word-of-wisdom/pkg/client"
	"github.com/ghost-circuit/word-of-wisdom/pkg/client/mocks"
)

func TestService_ExecuteSequential(t *testing.T) {
	t.Parallel()

	var (
		ctrl       = minimock.NewController(t)
		times      = uint64(10)
		challenge  = []byte("challenge")
		solution   = []byte("solution")
		difficulty = uint8(1)
		quote      = client.Quote{
			ID:     12,
			Text:   "text",
			Author: "author",
		}
	)

	clientMock := mocks.NewClientMock(ctrl)
	clientMock.GetChallengeMock.Times(times).Return(challenge, difficulty, nil)
	clientMock.SubmitSolutionMock.Times(times).Expect(challenge, solution).Return(quote, nil)

	powManagerMock := mocks.NewPowManagerMock(ctrl)
	powManagerMock.SolveCustomDifficultyMock.Times(times).Return(solution, nil)

	s := client.NewService(powManagerMock, clientMock, 0, int32(times))
	s.ExecuteSequential()
}

func TestService_ExecuteParallel(t *testing.T) {
	t.Parallel()

	var (
		ctrl       = minimock.NewController(t)
		times      = uint64(10)
		challenge  = []byte("challenge")
		solution   = []byte("solution")
		difficulty = uint8(1)
		quote      = client.Quote{
			ID:     12,
			Text:   "text",
			Author: "author",
		}
	)

	clientMock := mocks.NewClientMock(ctrl)
	clientMock.GetChallengeMock.Times(times).Return(challenge, difficulty, nil)
	clientMock.SubmitSolutionMock.Times(times).Expect(challenge, solution).Return(quote, nil)

	powManagerMock := mocks.NewPowManagerMock(ctrl)
	powManagerMock.SolveCustomDifficultyMock.Times(times).Return(solution, nil)

	s := client.NewService(powManagerMock, clientMock, 0, int32(times))
	s.ExecuteParallel()
}
