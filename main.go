package main

import (
	"fmt"
	"github.com/ElvertMora/quasar-fire-app/apm"
	"github.com/ElvertMora/quasar-fire-app/config"
	_ "github.com/ElvertMora/quasar-fire-app/docs"
	"github.com/ElvertMora/quasar-fire-app/providers"
	"github.com/ElvertMora/quasar-fire-app/router"
	"github.com/labstack/echo/v4"
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

