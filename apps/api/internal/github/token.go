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
	clientID := cfg.GithubAppClientID
	privateKey := cfg.GithubAppPrivateKey
	installationID := cfg.GithubAppInstallationID

	appTokenSource, err := githubauth.NewApplicationTokenSource(
		clientID,
		privateKey,
		githubauth.WithApplicationTokenExpiration(5*time.Minute),
	)

	if err != nil {
		return nil, err
	}

	installationTokenSource := githubauth.NewInstallationTokenSource(installationID, appTokenSource)

	return &GithubToken{source: &installationTokenSource}, nil
}
