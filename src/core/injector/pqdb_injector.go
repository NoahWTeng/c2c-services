package injector

import (
	"c2c-services/config"
	pqdb "c2c-services/src/core/database/pq"
)

func PqdbProvider(config *config.Global) *pqdb.Handler {

	return pqdb.NewPqdbConnection(config.PqdbVariables)
}
