package interfaces

import (
	"quasar-fire-app/controllers/request"
)

type ITrilaterationService interface {
	GetLocation(distances ...float64) (x, y float64)
}

type IMessagesService interface {
	GetMessage(messages ...[]string) (msg string)
}

type ISatelliteService interface {
	All() ([]*request.Satellite, error)
	One(id string) (*request.Satellite, error)
	Delete(name string) error
	Save(s *request.Satellite) error
}

