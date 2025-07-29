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

func NewGinHttpServer() interfaces.HttpServer {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Token"},
		AllowCredentials: true,
	}))
	return &GinHttpServer{
		engine: engine,
	}
}

func (g *GinHttpServer) On(
	method string,
	path string,
	controllers interfaces.HttpControllers,
	middlewares ...interfaces.HttpControllers,
) {

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
