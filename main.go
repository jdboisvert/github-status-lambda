package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct{}

func HandleRequest(ctx context.Context, event Event) error {
	err := UpdateGitHubProfileStatus(ctx)

	if err != nil {
		fmt.Println("Error updating GitHub profile status: ", err)
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
