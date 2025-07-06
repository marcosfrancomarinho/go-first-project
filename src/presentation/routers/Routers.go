package routers

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/shared/container"
)

type Routers struct {
	server interfaces.HttpServer
}

func NewRouters(server interfaces.HttpServer) *Routers {
	return &Routers{server: server}
}

func (r *Routers) Register(container *container.Container) {
	controllers := container.Dependencies()

	r.server.On("POST", "/register", controllers.CreatorUser)

}
