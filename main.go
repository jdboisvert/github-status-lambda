package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct{}

func HandleRequest(ctx context.Context, event Event) error {
	log.Println("Hello World!")
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
