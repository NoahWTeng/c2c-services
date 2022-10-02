package config

import (
	"c2c-services/src/core/database/mongodb"
	pqdb "c2c-services/src/core/database/pq"
)

type Base struct {
	Port        int    `mapstructure:"port" json:"port"`
	JWTKey      string `mapstructure:"jwt_secret" json:"jwt_secret"`
	Environment string `mapstructure:"environment" json:"environment"`
}

type Global struct {
	BaseVariables    *Base           `mapstructure:"base" json:"base"`
	MongodbVariables *mongodb.Config `mapstructure:"mongodb" json:"mongodb"`
	PqdbVariables    *pqdb.Config    `mapstructure:"postgresql" json:"postgresql"`
}
