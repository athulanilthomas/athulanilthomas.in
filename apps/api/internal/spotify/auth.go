package spotify

import (
	"context"
	"net/http"

	"github.com/athulanilthomas/www/api/internal/config"
	"golang.org/x/oauth2"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Auth struct {
	authenticator *spotifyauth.Authenticator
}

func NewAuth(cfg *config.Config) (*Auth, error) {
	a := spotifyauth.New(
		spotifyauth.WithClientID(cfg.SpotifyClientID),
		spotifyauth.WithClientSecret(cfg.SpotifyClientSecret),
		spotifyauth.WithRedirectURL(cfg.SpotifyRedirectURL),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState),
	)

	return &Auth{authenticator: a}, nil
}

func (a *Auth) GetAuthURL(state string) string {
	return a.authenticator.AuthURL(state)
}

func (a *Auth) Token(state string, r *http.Request) (*oauth2.Token, error) {
	return a.authenticator.Token(r.Context(), state, r)
}

func (a *Auth) NewClient(ctx context.Context, token *oauth2.Token) *spotify.Client {
	httpClient := a.authenticator.Client(ctx, token)

	return spotify.New(httpClient)
}
