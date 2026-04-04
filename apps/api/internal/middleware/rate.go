package middleware

import (
	"net/http"

	"github.com/athulanilthomas/www/api/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang.org/x/time/rate"
)

type RateLimitMiddleware struct {
	fx.Out
	Handler gin.HandlerFunc `name:"ratelimit"`
}

func NewRateLimitterMiddleware(cfg *config.Config) RateLimitMiddleware {
	limitter := rate.NewLimiter(rate.Limit(cfg.RateLimitRPS), cfg.RateLimitBurst)

	return RateLimitMiddleware{
		Handler: func(ctx *gin.Context) {
			if limitter.Allow() {
				ctx.Next()
				return
			}

			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			ctx.Abort()
		},
	}
}
