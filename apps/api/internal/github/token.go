package github

import (
	"time"

	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"

	"github.com/athulanilthomas/www/api/internal/config"
)

type GithubToken struct {
	source *oauth2.TokenSource
}

func NewGithubToken(cfg *config.Config) (*GithubToken, error) {
	clientID := cfg.GithubAppClientId
	privateKey := cfg.GithubAppPrivateKey

	tokenSource, err := githubauth.NewApplicationTokenSource(
		clientID,
		privateKey,
		githubauth.WithApplicationTokenExpiration(5*time.Minute),
	)

	if err != nil {
		return nil, err
	}

	return &GithubToken{source: &tokenSource}, nil
}
