package injector

import (
	"c2c-services/config"
	"c2c-services/src/core/database/mongodb"
)

func MongodbProvider(config *config.Global) *mongodb.Handler {
	return mongodb.NewMongodbConnection(config.MongodbVariables)
}
