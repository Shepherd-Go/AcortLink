package main

import (
	"fmt"
	"log"

	"acort.link/acort.link/cmd/api/router"
	"acort.link/acort.link/cmd/provider"
	"acort.link/acort.link/config"
	"github.com/labstack/echo/v4"
)

func main() {

	container := provider.BuildContainer()

	if err := container.Invoke(func(router *router.Router, server *echo.Echo, config config.Config) {

		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Server.Port)))

	}); err != nil {

		log.Fatal(err)
	}

}
