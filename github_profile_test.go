package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)

func TestUpdateGitHubStatus(t *testing.T) {
	defer gock.Off() // Flush pending mocks after each test

	t.Run("Test successful update", func(t *testing.T) {
		gock.New("https://api.github.com").
			Post("/graphql").
			Reply(http.StatusOK)

		err := UpdateGitHubProfileStatus(context.Background())
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("Test error in request", func(t *testing.T) {
		gock.New("https://api.github.com").
			Post("/graphql").
			ReplyError(errors.New("request error"))

		err := UpdateGitHubProfileStatus(context.Background())
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})

	t.Run("Test error in response status code", func(t *testing.T) {
		gock.New("https://api.github.com").
			Post("/graphql").
			Reply(http.StatusInternalServerError)

		err := UpdateGitHubProfileStatus(context.Background())
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}
