package pqdb

import (
	"database/sql"
)

type Config struct {
	DatabaseName string `mapstructure:"db_name" json:"db_name"`
	DatabaseHost string `mapstructure:"host" json:"host"`
	DatabasePort int    `mapstructure:"port" json:"port"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Sslmode      string `mapstructure:"sslmode" json:"sslmode"`
}

type Connector struct {
	config *Config
	db     *sql.DB
}

type Repository interface {
}

type Handler struct {
	Config     *Config
	Repository Repository
}
