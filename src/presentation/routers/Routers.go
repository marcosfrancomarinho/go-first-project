package routers

import (
	I "github.com/marcosfrancomarinho/src/domain/interfaces"
	C "github.com/marcosfrancomarinho/src/shared/container"
)

type Routers struct {
	server I.HttpServer
}

func NewRouters(server I.HttpServer) *Routers {
	return &Routers{server: server}
}

func (r *Routers) Register(container *C.Container) {
	controllers := container.Dependecies()

	r.server.On("POST", "/register", controllers["creator-user"])

}
