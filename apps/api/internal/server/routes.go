package server

import (
	"github.com/athulanilthomas/www/api/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, spotifyHandler *handler.SpotifyHandler, authHandler *handler.AuthHandler, githubHandler *handler.GithubHandler) {
	v1 := router.Group("/api/v1")

	spotify := v1.Group("/spotify")
	github := v1.Group("/github")

	spotify.GET("/login", authHandler.Login)
	spotify.GET("/callback", authHandler.Callback)
	spotify.GET("/now-playing", spotifyHandler.GetCurrentlyPlaying)

	github.GET("/repos", githubHandler.GetLastUpdatedRepositories)
}
