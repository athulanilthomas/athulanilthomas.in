package spotify

import (
	"context"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/zmb3/spotify/v2"

	"golang.org/x/oauth2"
)

func NewSpotifyClient(auth *Auth, cfg *config.Config) *spotify.Client {
	token := &oauth2.Token{
		RefreshToken: cfg.SpotifyRefreshToken,
	}

	httpClient := auth.authenticator.Client(context.Background(), token)
	return spotify.New(httpClient)
}
