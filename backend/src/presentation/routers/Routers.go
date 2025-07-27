package routers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/shared/container"
)

type Routers struct {
	server interfaces.HttpServer
}

func NewRouters(server interfaces.HttpServer) *Routers {
	return &Routers{server}
}

func (r *Routers) Register(container *container.Container) {
	handlers := container.Dependencies()

	r.server.On("POST", "/register", handlers.RegisterUserControllers)

	r.server.On("POST", "/login", handlers.LoginUserControllers)

	r.server.On("POST", "/product", handlers.CreatorProductControllers, handlers.UserAuthenticatorMiddlewares)

	r.server.On("GET", "/product", handlers.FindorProductControllers, handlers.UserAuthenticatorMiddlewares)

	r.server.On("DELETE", "/product/:id", handlers.DeleterProductController, handlers.UserAuthenticatorMiddlewares)
}
