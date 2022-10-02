package mongodb

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongodbConnection(cfg *Config) *Handler {
	mongodbHandler := &Connector{
		config: cfg,
	}

	err := mongodbHandler.connect()

	if err != nil {
		log.Fatal(err)
	}

	handler := &Handler{
		Config:     cfg,
		Repository: mongodbHandler,
	}

	return handler

}

func (c *Connector) connect() error {
	var (
		connectOnce sync.Once
		err         error
		client      *mongo.Client
		url         string
	)
	connectOnce.Do(func() {
		url = _formatUrl(c.config)
		client, err = mongo.NewClient(options.Client().ApplyURI(url))
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.config.DialTimeOut))
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

	})
	if err != nil {
		return err
	}
	c.client = client
	c.db = c.client.Database(c.config.DatabaseName)

	fmt.Printf("Mongodb success connected!\n")
	return nil
}

func (c *Connector) DB(ctx context.Context) *mongo.Database {
	return c.db
}
func (c *Connector) Client() *mongo.Client {
	return c.client
}

func (c *Connector) Config() Config {
	return *c.config
}

func _formatUrl(config *Config) string {
	if config.URI != "" {
		return config.URI
	}

	var b bytes.Buffer
	b.WriteString("mongodb://")
	if config.Username != "" {
		b.WriteString(config.Username)
		b.WriteString(":")
	}
	if config.Password != "" {
		b.WriteString(config.Password)
		b.WriteString("@")
	}
	b.WriteString(config.DatabaseHost)
	b.WriteString("/")
	b.WriteString(config.DatabaseName)

	var urlQueryString []string

	if config.PoolSize != 0 {
		urlQueryString = append(urlQueryString, fmt.Sprintf("maxPoolSize=%d", config.PoolSize))
	}

	if config.ReplicaSet != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("replicaSet=%s", config.ReplicaSet))
	}

	if config.AuthSource != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("authSource=%s", config.AuthSource))
	}

	if len(urlQueryString) > 0 {
		s := strings.Join(urlQueryString, "&")
		s = "?" + s
		b.WriteString(s)
	}

	return b.String()
}
