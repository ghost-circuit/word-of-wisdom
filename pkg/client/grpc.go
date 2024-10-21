package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/generated/wisdom"
	desc "github.com/alisher-baizhumanov/word-of-wisdom/pkg/generated/wisdom"
)

// GrpcClient represents a gRPC client.
type GrpcClient struct {
	grpcClient desc.WordOfWisdomServiceClient
}

// NewClient creates and returns a new GrpcClient instance.
func NewClient(address string) (*GrpcClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("client.NewClient: %w", err)
	}

	return &GrpcClient{
		grpcClient: wisdom.NewWordOfWisdomServiceClient(conn),
	}, nil
}

// GetChallenge requests a challenge from the server.
func (c *GrpcClient) GetChallenge() ([]byte, uint8, error) {
	challenge, err := c.grpcClient.GetChallenge(context.Background(), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("client.GetChallenge: %w", err)
	}

	return challenge.GetChallenge(), uint8(challenge.GetDifficulty()), nil
}

// SubmitSolution submits a solution to the server.
func (c *GrpcClient) SubmitSolution(challenge []byte, solution []byte) (Quote, error) {
	resp, err := c.grpcClient.SubmitSolution(context.Background(), &desc.SolutionRequest{
		Challenge: challenge,
		Solution:  solution,
	})
	if err != nil {
		return Quote{}, fmt.Errorf("client.SubmitSolution: %w", err)
	}

	return Quote{
		ID:     resp.GetQuote().GetId(),
		Text:   resp.GetQuote().GetText(),
		Author: resp.GetQuote().GetAuthor(),
	}, nil
}
