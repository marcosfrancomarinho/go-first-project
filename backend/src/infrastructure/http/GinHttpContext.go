package http

import (
	"errors"
	"fmt"
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


func (g *GinHttpContext) GetParams(key string) (*string, error) {
	id := g.ctx.Param(key)
	if len(id) == 0 {
		return nil, fmt.Errorf("parametro %s não foi fornecido", key)
	}
	return &id, nil
}

func (g *GinHttpContext) GetQuery(keys ...string) (map[string]string, error) {
	queries := make(map[string]string)

	for _, key := range keys {
		query := g.ctx.Query(key)

		if len(query) == 0 {
			return nil, fmt.Errorf("query param '%s' inválido ou não definido", key)
		}

		queries[key] = query
	}
	return queries, nil
}
