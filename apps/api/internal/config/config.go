package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURL  string
	SpotifyRefreshToken string
	GithubAppClientId   string
	GithubAppPrivateKey []byte
	Port                int
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	githubAppPrivateKey, privateKeyError := base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_PRIVATE_KEY"))

	if err != nil {
		port = 8080
	}

	if privateKeyError != nil {
		githubAppPrivateKey = nil
	}

	config := &Config{
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URL"),
		SpotifyRefreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		GithubAppClientId:   os.Getenv("GITHUB_APP_CLIENT_ID"),
		GithubAppPrivateKey: githubAppPrivateKey,
		Port:                port,
	}

	if config.GithubAppClientId == "" || string(config.GithubAppPrivateKey) == "" {
		return nil, fmt.Errorf("Missing required github credentials in environment")
	}

	if config.SpotifyClientID == "" || config.SpotifyClientSecret == "" {
		return nil, fmt.Errorf("Missing required Spotify credentials in environment")
	}

	return config, nil
}
