package http

import (
	"github.com/gin-gonic/gin"
)

type GinHttpContext struct {
	ctx *gin.Context
}

func NewGinHttpContext(ctx *gin.Context) *GinHttpContext {
	return &GinHttpContext{ctx: ctx}
}

func (g *GinHttpContext) GetBody(input any) error {
	return g.ctx.ShouldBindJSON(input)
}

func (g *GinHttpContext) Send(status int, datas any) {
	g.ctx.JSON(status, datas)
}

func (g *GinHttpContext) SendError(err error) {
	g.ctx.JSON(400, gin.H{
		"error":  err.Error(),
		"status": false,
	})
}
