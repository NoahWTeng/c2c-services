package app

import (
	"c2c-services/config"
	"c2c-services/src/libs/database/mongodb"
)

type AppContainer struct {
	config  *config.Global
	mongodb *mongodb.Handler
}

type Application interface {
	Init() error
}
