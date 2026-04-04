package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MiddlewareParams struct {
	fx.In
	Auth      gin.HandlerFunc `name:"auth"`
	RateLimit gin.HandlerFunc `name:"ratelimit"`
}

type Middlewares struct {
	Auth      gin.HandlerFunc
	RateLimit gin.HandlerFunc
}

func NewMiddlewares(mp MiddlewareParams) *Middlewares {
	return &Middlewares{
		Auth:      mp.Auth,
		RateLimit: mp.RateLimit,
	}
}
