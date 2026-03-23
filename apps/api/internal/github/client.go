package github

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewGithubClient(token *GithubToken) (*github.Client, error) {
	httpClient := oauth2.NewClient(context.Background(), *token.source)

	return github.NewClient(httpClient), nil
}
