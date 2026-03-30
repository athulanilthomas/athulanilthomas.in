package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MiddlewareParams struct {
	fx.In
	Handlers []gin.HandlerFunc `group:"middlewares"`
}

type Middlewares struct {
	Handlers []gin.HandlerFunc
}

func NewMiddlewares(mp MiddlewareParams) *Middlewares {
	return &Middlewares{
		Handlers: mp.Handlers,
	}
}
