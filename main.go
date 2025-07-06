package main

import (
	"github.com/marcosfrancomarinho/src/infrastructure/http"
	"github.com/marcosfrancomarinho/src/presentation/routers"
	"github.com/marcosfrancomarinho/src/shared/container"
)

func main() {
	server := http.NewGinHttpServer()
	container := &container.Container{}
	router := routers.NewRouters(server)
	router.Register(container)
	server.Run(8080)

}
