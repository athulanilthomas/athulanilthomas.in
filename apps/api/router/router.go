package router

import (
	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/athulanilthomas/www/api/internal/github"
	"github.com/athulanilthomas/www/api/internal/handler"
	"github.com/athulanilthomas/www/api/internal/middleware"
	"github.com/athulanilthomas/www/api/internal/server"
	"github.com/athulanilthomas/www/api/internal/service"
	"github.com/athulanilthomas/www/api/internal/spotify"
	"github.com/gin-gonic/gin"
)

func BuildRouter() (*gin.Engine, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	auth, err := spotify.NewAuth(cfg)
	if err != nil {
		return nil, err
	}

	githubToken, err := github.NewGithubToken(cfg)
	if err != nil {
		return nil, err
	}

	githubClient, err := github.NewGithubClient(githubToken)
	if err != nil {
		return nil, err
	}

	spotifyClient := spotify.NewSpotifyClient(auth, cfg)
	serviceInstance, err := service.NewService(spotifyClient)
	if err != nil {
		return nil, err
	}

	githubService, err := service.NewGithubService(githubClient)
	if err != nil {
		return nil, err
	}

	spotifyHandler := handler.NewSpotifyHandler(serviceInstance)
	authHandler := handler.NewAuthHandler(auth, serviceInstance, cfg)
	githubHandler := handler.NewGithubHandler(githubService)
	healthHandler := handler.NewHealthHandler(githubService)

	rateLimitMiddleware := middleware.NewRateLimitterMiddleware(cfg)
	authMiddleware := middleware.NewAuthenticationMiddleware(cfg)

	middlewares := middleware.NewMiddlewares(middleware.MiddlewareParams{
		Auth:      authMiddleware.Handler,
		RateLimit: rateLimitMiddleware.Handler,
	})

	router := server.NewRouter(middlewares)
	server.RegisterRoutes(router, spotifyHandler, authHandler, githubHandler, healthHandler, middlewares)

	return router, nil
}
