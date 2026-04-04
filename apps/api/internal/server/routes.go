package server

import (
	"github.com/athulanilthomas/www/api/internal/handler"
	"github.com/athulanilthomas/www/api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	spotifyHandler *handler.SpotifyHandler,
	authHandler *handler.AuthHandler,
	githubHandler *handler.GithubHandler,
	healthHandler *handler.HealthHandler,
	middlewares *middleware.Middlewares,

) {
	v1 := router.Group("/api/v1")

	v1.GET("/health", healthHandler.Health)

	spotify := v1.Group("/spotify")
	github := v1.Group("/github", middlewares.Auth)

	spotify.GET("/login", authHandler.Login)
	spotify.GET("/callback", authHandler.Callback)
	spotify.GET("/now-playing", middlewares.Auth, spotifyHandler.GetCurrentlyPlaying)

	github.GET("/repos", githubHandler.GetLastUpdatedRepositories)
}
