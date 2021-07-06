package providers

import (
	"github.com/ElvertMora/quasar-fire-app/controllers"
	"github.com/ElvertMora/quasar-fire-app/router"
	"github.com/ElvertMora/quasar-fire-app/server"
	"github.com/ElvertMora/quasar-fire-app/services"
	"go.uber.org/dig"
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
