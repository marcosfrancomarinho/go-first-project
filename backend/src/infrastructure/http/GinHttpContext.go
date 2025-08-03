package http

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/shared/exceptions"
)

type GinHttpContext struct {
	ctx *gin.Context
}

func NewGinHttpContext(ctx *gin.Context) interfaces.HttpContext {
	return &GinHttpContext{ctx}
}

func (g *GinHttpContext) GetBody(input any) error {
	return g.ctx.BindJSON(input)
}

func (g *GinHttpContext) Send(status int, datas any, token ...string) {
	if len(token) > 0 {
		g.ctx.Header("token", token[0])
	}
	g.ctx.JSON(status, datas)
}

func (g *GinHttpContext) SendError(err error) {
	code := "ERROR"
	tokenErros := []error{
		jwt.ErrInvalidType,
		jwt.ErrTokenMalformed,
		exceptions.ErrTokenInvalid,
		jwt.ErrTokenExpired,
		jwt.ErrSignatureInvalid,
	}
	for _, error := range tokenErros {
		if errors.Is(err, error) {
			code = "TOKEN_ERROR"
			break
		}
	}
	g.ctx.JSON(400, gin.H{
		"error":  err.Error(),
		"status": false,
		"code":   code,
	})
}

func (g *GinHttpContext) GetToken() string {
	token := g.ctx.GetHeader("token")
	return token
}

func (g *GinHttpContext) GetParams(input any) error {
	if err := g.ctx.BindUri(input); err != nil {
		return err
	}
	return nil
}

func (g *GinHttpContext) GetQuery(input any) error {
	if err := g.ctx.BindQuery(input); err != nil {
		return err
	}
	return nil
}
