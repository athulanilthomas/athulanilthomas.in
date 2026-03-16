package handler

import (
	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type SpotifyHandler struct {
	service *service.Service
}

func NewSpotifyHandler(service *service.Service) *SpotifyHandler {
	return &SpotifyHandler{service: service}
}

func (h *SpotifyHandler) GetCurrentlyPlaying(c *gin.Context) {
	currentlyPlaying, err := h.service.GetCurrentlyPlaying(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if currentlyPlaying == nil || !currentlyPlaying.Playing {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, currentlyPlaying)
}
