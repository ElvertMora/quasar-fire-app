package main

import (
	"fmt"
	"github.com/labstack/echo"
	"quasar-fire-app/apm"
	"quasar-fire-app/config"
	_ "quasar-fire-app/docs"
	"quasar-fire-app/providers"
	"quasar-fire-app/router"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

// @title Quasar Fire API
// @version 1.0
// @description Operación Fuego de Quasar.
// @contact.name Elvert Mora
// @contact.email emorae@unal.edu.co
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/quasar-fire
func main() {

	apm.Get().Start()
	defer apm.Get().Stop()
	container := providers.BuildContainer()

	err := container.Invoke(func(server *echo.Echo, route *router.Router) {
		server.Debug = config.Environments().Postfix == "dev"
		route.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		panic(err)
	}

}

