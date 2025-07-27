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
	return g.ctx.ShouldBindJSON(input)
}

func (g *GinHttpContext) Send(status int, datas any, token ...string) {
	if len(token) > 0 {
		g.ctx.Header("token", token[0])
	}
	g.ctx.JSON(status, datas)
}

func (g *GinHttpContext) SendError(err error) {
	code := "ERROR"
	if errors.Is(err, jwt.ErrInvalidType) || errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, exceptions.ErrTokenInvalid) {
		code = "TOKEN_ERROR"
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

func (g *GinHttpContext) SetIdentifiers(key string, datas any) {
	g.ctx.Set(key, datas)
}

func (g *GinHttpContext) GetIdentifiers(key string) (any, error) {
	value, exitst := g.ctx.Get(key)
	if !exitst {
		return nil, errors.New("dados da chave informada não encontrado")
	}
	return value, nil
}
