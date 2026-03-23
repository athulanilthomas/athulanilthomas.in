package service

import (
	"context"
	"sync"

	"github.com/google/go-github/github"
)

type GithubService struct {
	client   *github.Client
	username string
	mux      *sync.RWMutex
}

func NewGithubService(client *github.Client) (*GithubService, error) {
	svc := &GithubService{
		client:   client,
		username: "athulanilthomas",
		mux:      &sync.RWMutex{},
	}

	return svc, nil
}

func (g *GithubService) GetLastUpdatedRepositories(ctx context.Context) ([]*github.Repository, error) {
	opts := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}

	g.mux.RLock()
	client := g.client
	g.mux.RUnlock()

	repos, _, err := client.Repositories.List(ctx, g.username, opts)

	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (g *GithubService) SetClient(client *github.Client) {
	g.mux.Lock()
	g.client = client
	g.mux.Unlock()
}
