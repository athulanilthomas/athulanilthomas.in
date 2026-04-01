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
	GithubAppClientID       string
	GithubAppPrivateKey     []byte
	GithubAppInstallationID int64
	RateLimitRPS            float64
	RateLimitBurst          int
	Port                    int
	XApiSecret              string
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	githubAppPrivateKey, privateKeyError := base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_PRIVATE_KEY"))
	githubAppInstallationID, installationIDError := strconv.ParseInt(os.Getenv("GITHUB_APP_INSTALLATION_ID"), 10, 64)

	rateLimitRps, rateLimitRpsError := strconv.ParseFloat(os.Getenv("RATE_LIMIT_RPS"), 64)
	rateLimitBurst, rateLimitBurstError := strconv.Atoi(os.Getenv("RATE_LIMIT_BURST"))

	if err != nil {
		port = 8080
	}

	if privateKeyError != nil {
		githubAppPrivateKey = nil
	}

	if installationIDError != nil {
		githubAppInstallationID = 0
	}

	if rateLimitRpsError != nil {
		rateLimitRps = 5
	}

	if rateLimitBurstError != nil {
		rateLimitBurst = 10
	}

	config := &Config{
		SpotifyClientID:         os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret:     os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:      os.Getenv("SPOTIFY_REDIRECT_URL"),
		SpotifyRefreshToken:     os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		GithubAppClientID:       os.Getenv("GITHUB_APP_CLIENT_ID"),
		GithubAppPrivateKey:     githubAppPrivateKey,
		GithubAppInstallationID: githubAppInstallationID,
		RateLimitRPS:            rateLimitRps,
		RateLimitBurst:          rateLimitBurst,
		Port:                    port,
		XApiSecret:              os.Getenv("X_API_SECRET"),
	}

	if config.GithubAppClientID == "" || string(config.GithubAppPrivateKey) == "" || config.GithubAppInstallationID == 0 {
		return nil, fmt.Errorf("Missing required Github credentials in environment")
	}

	if config.SpotifyClientID == "" || config.SpotifyClientSecret == "" {
		return nil, fmt.Errorf("Missing required Spotify credentials in environment")
	}

	return config, nil
}
