package wisdom_task_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	"github.com/ghost-circuit/word-of-wisdom/internal/domain/entity"
	"github.com/ghost-circuit/word-of-wisdom/internal/domain/interfaces/mocks"
	wisdomtask "github.com/ghost-circuit/word-of-wisdom/internal/domain/service/wisdom-task"
)

func TestWordOfWisdomService_GetChallenge(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	// Create a mock for the PoWManager interface
	powManagerMock := mocks.NewProofOfWorkManagerMock(mc)
	service := wisdomtask.NewWisdomTaskService(nil, powManagerMock)

	expectedChallenge := []byte("challenge")
	expectedDifficulty := uint8(5)
	expectedError := error(nil)

	// Set up the expectation for the GenerateChallenge method
	powManagerMock.GenerateChallengeMock.Expect().Return(expectedChallenge, expectedDifficulty, expectedError)

	// Call the method
	challenge, difficulty, err := service.GetChallenge(context.Background())

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedChallenge, challenge)
	assert.Equal(t, expectedDifficulty, difficulty)
}

func TestGetRandomQuote_Error(t *testing.T) {
	t.Parallel()

	// Create a new minimock controller
	mc := minimock.NewController(t)

	// Create a new QuoteRepositoryMock
	mockQuoteRepo := mocks.NewQuoteRepositoryMock(mc)

	// Set up the mock to return an error
	expectedError := errors.New("some error")
	mockQuoteRepo.GetRandomQuoteMock.Return(entity.Quote{}, expectedError)

	// Call the method
	_, err := mockQuoteRepo.GetRandomQuote(context.Background())

	// Verify the error
	assert.ErrorIs(t, err, expectedError)
}
