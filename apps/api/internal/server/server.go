package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: router,
		},
	}
}

func NewRouter() *gin.Engine {
	return gin.Default()
}

func Start(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.httpServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.httpServer.Shutdown(ctx)
		},
	})
}
