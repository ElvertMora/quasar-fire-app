package apm

import (
	i "github.com/ElvertMora/quasar-fire-app/interfaces"
	"sync"
)

var (
	instance i.APM
	once     sync.Once
)

func Get() i.APM {
	once.Do(func() {
		instance = getInstance()
	})

	return instance
}

func getInstance() i.APM {
	return createEmptyAPM()
}
