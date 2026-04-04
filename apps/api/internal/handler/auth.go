package handler

import (
	"net/http"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/athulanilthomas/www/api/internal/spotify"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth    *spotify.Auth
	service *service.Service
	config  *config.Config
}

func NewAuthHandler(auth *spotify.Auth, service *service.Service, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		auth:    auth,
		service: service,
		config:  cfg,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	url := h.auth.GetAuthURL(h.config.SpotifyStateVariable)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) Callback(c *gin.Context) {
	state := c.Query("state")

	if state != h.config.SpotifyStateVariable {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "State mismatch"})
		return
	}

	token, err := h.auth.Token(state, c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := h.auth.NewClient(c.Request.Context(), token)

	h.service.SetClient(client)

	c.JSON(http.StatusOK, gin.H{"message": "successfully authenticated"})
}
