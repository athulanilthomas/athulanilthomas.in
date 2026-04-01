package middleware

import (
	"net/http"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthenticationMiddleware struct {
	fx.Out
	Handler gin.HandlerFunc `group:"middlewares"`
}

func NewAuthenticationMiddleware(cfg *config.Config) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		Handler: func(ctx *gin.Context) {
			if token := ctx.GetHeader("X-API-Secret"); token != cfg.XApiSecret {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			ctx.Next()
		},
	}
}
