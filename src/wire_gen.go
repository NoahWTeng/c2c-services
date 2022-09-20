// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"c2c-services/src/libs/injector"
)

// Injectors from wire.go:

func CreateNewApp() (*AppContainer, error) {
	global, err := injector.ConfigProvider()
	if err != nil {
		return nil, err
	}
	handler := injector.MongodbProvider(global)
	appContainer := &AppContainer{
		config:  global,
		mongodb: handler,
	}
	return appContainer, nil
}
