package config

import (
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
	Port                int
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 8080
	}

	config := &Config{
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URL"),
		SpotifyRefreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		Port:                port,
	}

	if config.SpotifyClientID == "" || config.SpotifyClientSecret == "" {
		return nil, fmt.Errorf("Missing required Spotify credentials in environment")
	}

	return config, nil
}
