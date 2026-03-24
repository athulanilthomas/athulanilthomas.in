package github

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewGithubClient(token *GithubToken) (*github.Client, error) {
	if token == nil {
		return nil, errors.New("github token is nil")
	}
	if token.source == nil {
		return nil, errors.New("github token source is nil")
	}

	httpClient := oauth2.NewClient(context.Background(), *token.source)

	return github.NewClient(httpClient), nil
}
