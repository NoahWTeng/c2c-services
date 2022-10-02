package pqdb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

func NewPqdbConnection(cfg *Config) *Handler {

	pqdbHandler := &Connector{
		config: cfg,
	}

	err := pqdbHandler.connect()

	if err != nil {
		log.Fatal(err)
	}

	handler := &Handler{
		Config:     cfg,
		Repository: pqdbHandler,
	}

	return handler

}

func (c *Connector) connect() error {
	var (
		connectOnce sync.Once
		err         error
		db          *sql.DB
		url         string
	)

	connectOnce.Do(func() {
		url = _formatUrl(c.config)

		db, err = sql.Open("postgres", url)

		err = db.Ping()

		if err != nil {
			log.Fatal(err)
			return
		}

	})

	if err != nil {
		return err
	}
	c.db = db

	fmt.Printf("Postgres success connected!\n")
	return nil
}

func _formatUrl(config *Config) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.Username, config.Password, config.DatabaseName)

	return psqlInfo
}
