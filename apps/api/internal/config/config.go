package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SpotifyClientID         string
	SpotifyClientSecret     string
	SpotifyRedirectURL      string
	SpotifyRefreshToken     string
	GithubAppClientId       string
	GithubAppPrivateKey     []byte
	GithubAppInstallationId int64
	Port                    int
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	githubAppPrivateKey, privateKeyError := base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_PRIVATE_KEY"))
	githubAppInstallationId, installationIdError := strconv.ParseInt(os.Getenv("GITHUB_APP_INSTALLATION_ID"), 10, 64)

	if err != nil {
		port = 8080
	}

	if privateKeyError != nil {
		githubAppPrivateKey = nil
	}

	if installationIdError != nil {
		githubAppInstallationId = 0
	}

	config := &Config{
		SpotifyClientID:         os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret:     os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:      os.Getenv("SPOTIFY_REDIRECT_URL"),
		SpotifyRefreshToken:     os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		GithubAppClientId:       os.Getenv("GITHUB_APP_CLIENT_ID"),
		GithubAppPrivateKey:     githubAppPrivateKey,
		GithubAppInstallationId: githubAppInstallationId,
		Port:                    port,
	}

	if config.GithubAppClientId == "" || string(config.GithubAppPrivateKey) == "" || config.GithubAppInstallationId == 0 {
		return nil, fmt.Errorf("Missing required Github credentials in environment")
	}

	if config.SpotifyClientID == "" || config.SpotifyClientSecret == "" {
		return nil, fmt.Errorf("Missing required Spotify credentials in environment")
	}

	return config, nil
}
