package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	 "github.com/marcosfrancomarinho/src/domain/interfaces"
)

type GinHttpServer struct {
	engine *gin.Engine
}

func NewGinHttpServer() *GinHttpServer {
	return &GinHttpServer{
		engine: gin.Default(),
	}
}

func (g *GinHttpServer) On(method string, path string, controllers interfaces.HttpControllers) {

	g.engine.Handle(method, path, func(ctx *gin.Context) {
		httpContext := NewGinHttpContext(ctx)
		controllers.Execute(httpContext)
	})
}

func (g *GinHttpServer) Run(port int) {
	g.engine.Run(fmt.Sprintf(":%d", port))
}
