package service

import (
	"context"

	"github.com/google/go-github/github"
)

type GithubService struct {
	client   *github.Client
	username string
}

func NewGithubService(client *github.Client) (*GithubService, error) {
	svc := &GithubService{
		client:   client,
		username: "athulanilthomas",
	}

	return svc, nil
}

func (g *GithubService) GetLastUpdatedRepositories(ctx context.Context) ([]*github.Repository, error) {
	// Fetch only the last 3 repos
	opts := &github.RepositoryListOptions{
		Type:      "owner",
		Sort:      "updated",
		Direction: "desc",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 3,
		},
	}

	repos, _, err := g.client.Repositories.List(ctx, g.username, opts)

	if err != nil {
		return nil, err
	}

	return repos, nil
}
