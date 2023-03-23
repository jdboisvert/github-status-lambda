package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// UpdateGitHubProfileStatus updates the GitHub profile status.
func UpdateGitHubProfileStatus() {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		log.Panicln("GITHUB_TOKEN environment variable is not set.")
		return
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	// TODO This is what worked for me in the playground.
	// mutation {
	// 	changeUserStatus(
	// 	  input: {clientMutationId: "jdboisvert", emoji: ":rocket:", message: "I wonder what a test is."}
	// 	) {
	// 	  clientMutationId
	// 	  status {
	// 		message
	// 		emoji
	// 		updatedAt
	// 	  }
	// 	}
	//   }

	// TODO This will be randomized later on to be a quote and emoji.
	emoji := githubv4.String(":rocket:")
	message := githubv4.String("I wonder what a test is.")

	input := githubv4.ChangeUserStatusInput{
		Emoji:   &emoji,
		Message: &message,
	}

	var mutation struct {
		UpdateUserStatus struct {
			ClientMutationID string
		} `graphql:"updateUserStatus(input: $input)"`
	}

	err := client.Mutate(context.Background(), &mutation, map[string]interface{}{
		"input": input, // I suspect that this will fail since it will not align with the GraphQL schema.
	}, nil)
	if err != nil {
		fmt.Printf("Error updating user status: %v\n", err)
		return
	}

	fmt.Println("User status updated successfully.")
}
