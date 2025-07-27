package main

import (
	"log"
	"github.com/marcosfrancomarinho/go-first-project/src/infrastructure/http"
	"github.com/marcosfrancomarinho/go-first-project/src/presentation/routers"
	"github.com/marcosfrancomarinho/go-first-project/src/shared/container"
)

func main() {
	server := http.NewGinHttpServer()
	container := container.GetInstance()
	router := routers.NewRouters(server)
	router.Register(container)
	if err := server.Run(8080); err != nil {
		log.Fatal(err)
	}
}

