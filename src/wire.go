//go:build wireinject
// +build wireinject

package app

import (
	"c2c-services/src/libs/injector"
	"github.com/google/wire"
)

func CreateNewApp() (*AppContainer, error) {
	panic(wire.Build(
		injector.ConfigProvider,
		injector.MongodbProvider,
		wire.Struct(new(AppContainer), "*"),
	))
}
