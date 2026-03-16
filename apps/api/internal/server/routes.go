package server

import (
	"github.com/athulanilthomas/www/api/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, spotifyHandler *handler.SpotifyHandler, authHandler *handler.AuthHandler) {
	v1 := router.Group("/api/v1")

	v1.GET("/login", authHandler.Login)
	v1.GET("/callback", authHandler.Callback)
	v1.GET("/now-playing", spotifyHandler.GetCurrentlyPlaying)
}
