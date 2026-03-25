package handler

import (
	"github.com/athulanilthomas/www/api/internal/handler/dto"
	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type GithubHandler struct {
	service *service.GithubService
}

func NewGithubHandler(service *service.GithubService) *GithubHandler {
	return &GithubHandler{service: service}
}

func (h *GithubHandler) GetLastUpdatedRepositories(c *gin.Context) {
	repos, err := h.service.GetLastUpdatedRepositories(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToRepoDTOs(repos))
}
