package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DatabaseName    string `mapstructure:"db_name" json:"db_name"`
	DatabaseHost    string `mapstructure:"host" json:"host"`
	TimeOut         int    `mapstructure:"timeout" json:"timeout"`
	DialTimeOut     int64  `mapstructure:"dial_timeout" json:"dial_timeout"`
	PoolSize        int    `mapstructure:"pool_size" json:"pool_size"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	ReplicaSet      string `mapstructure:"replica_set" json:"replica_set"`
	AuthSource      string `mapstructure:"auth_source" json:"auth_source"`
	URI             string `mapstructure:"uri" json:"uri"`
	IDPrefix        string `mapstructure:"id_prefix" json:"id_prefix"`
	EnabledNRTracer bool   `mapstructure:"newrelic_tracer" json:"newrelic_tracer"`
}

type Connector struct {
	config *Config
	client *mongo.Client
	db     *mongo.Database
}

type Repository interface {
	DB(ctx context.Context) *mongo.Database
	Client() *mongo.Client
	Config() Config
}

type Handler struct {
	Config     *Config
	Repository Repository
}
