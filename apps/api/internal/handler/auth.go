package handler

import (
	"net/http"

	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/athulanilthomas/www/api/internal/spotify"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth    *spotify.Auth
	service *service.Service
}

func NewAuthHandler(auth *spotify.Auth, service *service.Service) *AuthHandler {
	return &AuthHandler{
		auth:    auth,
		service: service,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	url := h.auth.GetAuthURL("state123")

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) Callback(c *gin.Context) {
	state := c.Query("state")

	if state != "state123" {
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
