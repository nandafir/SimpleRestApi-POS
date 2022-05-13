package database

import (
	"anaconda/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// New ...
func New() *sqlx.DB {
	config := config.Get()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal("connection error!")
	}

	db.SetConnMaxLifetime(config.DBConnMaxLifetime)
	db.SetConnMaxIdleTime(config.DBConnMaxIdleTime)
	db.SetMaxIdleConns(config.DBConnMaxIdle)
	db.SetMaxOpenConns(config.DBConnMaxOpen)

	return db
}
