package providers

import (
	"go.uber.org/dig"
	"quasar-fire-app/controllers"
	"quasar-fire-app/router"
	"quasar-fire-app/server"
	"quasar-fire-app/services"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()
	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(services.NewTrilaterationService)
	_ = Container.Provide(services.NewMessagesService)
	_ = Container.Provide(services.NewSatelliteService)

	_ = Container.Provide(controllers.NewTopSecretSplitHandler)
	_ = Container.Provide(controllers.NewTopSecretHandler)
	_ = Container.Provide(controllers.NewSatelliteHandler)
	return Container
}
