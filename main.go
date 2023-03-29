package main

import (
	"context"
	// "github.com/aws/aws-lambda-go/lambda"
)

type Event struct{}

func HandleRequest(ctx context.Context, event Event) error {
	UpdateGitHubProfileStatus(ctx)
	return nil
}

func main() {
	HandleRequest(context.Background(), Event{})
	// lambda.Start(HandleRequest)
}
