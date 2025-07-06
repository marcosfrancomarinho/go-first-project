package main

import (
	h "github.com/marcosfrancomarinho/src/infrastructure/http"
	r "github.com/marcosfrancomarinho/src/presentation/routers"
	c "github.com/marcosfrancomarinho/src/shared/container"
)

func main() {
	server := h.NewGinHttpServer()
	container := &c.Container{}
	router := r.NewRouters(server)
	router.Register(container)
	server.Run(8080)

}
