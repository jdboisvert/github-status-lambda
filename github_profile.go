package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func updateGitHubStatus(ctx context.Context, accessToken, statusMessage, emoji string) error {
	httpClient := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.github.com/graphql", nil)
	if err != nil {
		log.Println("Error creating request: ", err)
		return err
	}
	req.Header.Set("Authorization", "bearer "+accessToken)

	query := fmt.Sprintf(`mutation {
		changeUserStatus(
		  input: {clientMutationId: "jdboisvert", emoji: %q, message: %q}
		) {
		  clientMutationId
		  status {
			message
			emoji
			updatedAt
		  }
		}
	  }`, emoji, statusMessage)

	payload := map[string]string{"query": query}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(reqBody))

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	return nil
}

func handler(ctx context.Context) error {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}

	statusEntries, err := ReadStatusEntriesFromFile()
	if err != nil {
		return fmt.Errorf("failed to read status entries: %v", err)
	}

	// Pick a random status entry.
	rand.Seed(time.Now().UnixNano())
	randomStatusEntry := statusEntries[rand.Intn(len(statusEntries))]

	if err := updateGitHubStatus(ctx, accessToken, randomStatusEntry.Message, randomStatusEntry.Emoji); err != nil {
		return fmt.Errorf("failed to update GitHub status: %v", err)
	}

	log.Println("GitHub status updated successfully with message: ", randomStatusEntry.Message)

	return nil
}

// UpdateGitHubProfileStatus updates the GitHub profile status.
func UpdateGitHubProfileStatus(ctx context.Context) error {
	return handler(ctx)
}
