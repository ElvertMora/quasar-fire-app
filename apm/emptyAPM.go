package apm

import (
	"github.com/ElvertMora/quasar-fire-app/interfaces"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type emptyAPM struct{}

func createEmptyAPM() interfaces.APM {
	return &emptyAPM{}
}

func (e *emptyAPM) Start() {
	log.Info("Starting empty APM")
}

func (e *emptyAPM) Stop() {
	log.Info("Shutting down empty APM")
}

func (e emptyAPM) Middleware() echo.MiddlewareFunc {
	return nil
}
