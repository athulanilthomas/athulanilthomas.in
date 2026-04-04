package handler

import (
	"time"

	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type HealthHandler struct {
	service *service.GithubService
}

func NewHealthHandler(service *service.GithubService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, time.Now())
}
