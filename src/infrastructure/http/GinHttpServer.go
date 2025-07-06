package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type GinHttpServer struct {
	engine *gin.Engine
}

func NewGinHttpServer() *GinHttpServer {
	return &GinHttpServer{
		engine: gin.Default(),
	}
}

func (g *GinHttpServer) On(method string, path string, controllers interfaces.HttpControllers, middlewares ...interfaces.HttpControllers) {

	g.engine.Handle(method, path, func(ctx *gin.Context) {
		httpContext := NewGinHttpContext(ctx)

		for _, middleware := range middlewares {
			middleware.Execute(httpContext)
			if ctx.Writer.Written() {
				return
			}
		}
		controllers.Execute(httpContext)
	})
}

func (g *GinHttpServer) Run(port int) {
	g.engine.Run(fmt.Sprintf(":%d", port))
}
