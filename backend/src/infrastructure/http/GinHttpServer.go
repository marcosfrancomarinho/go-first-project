package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
)

type GinHttpServer struct {
	engine *gin.Engine
}

func NewGinHttpServer() *GinHttpServer {
	gin.SetMode(gin.ReleaseMode)
	return &GinHttpServer{
		engine: gin.Default(),
	}
}

func (g *GinHttpServer) On(
	method string,
	path string,
	controllers interfaces.HttpControllers,
	middlewares ...interfaces.HttpControllers,
) {
	g.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Token"},
		AllowCredentials: true,
	}))
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

func (g *GinHttpServer) Run(port int) error {
	return g.engine.Run(fmt.Sprintf(":%d", port))
}
