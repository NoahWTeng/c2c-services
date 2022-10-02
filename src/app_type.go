package app

import (
	"c2c-services/config"
	"c2c-services/src/core/database/mongodb"
	pqdb "c2c-services/src/core/database/pq"
)

type AppContainer struct {
	config  *config.Global
	mongodb *mongodb.Handler
	pqdb    *pqdb.Handler
}

type Application interface {
	Init() error
}
