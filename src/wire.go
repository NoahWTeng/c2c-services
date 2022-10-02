//go:build wireinject
// +build wireinject

package app

import (
	"c2c-services/src/core/injector"
	"github.com/google/wire"
)

func CreateNewApp() (*AppContainer, error) {
	panic(wire.Build(
		injector.ConfigProvider,
		injector.MongodbProvider,
		injector.PqdbProvider,
		wire.Struct(new(AppContainer), "*"),
	))
}
