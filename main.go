package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct{}

func HandleRequest(ctx context.Context, event Event) error {
	// Your code here
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
