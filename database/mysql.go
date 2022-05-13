package database

import (
	"anaconda/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// New ...
func New() *sqlx.DB {
	config := config.Get()

	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("connection error!")
	}

	db.SetConnMaxLifetime(config.DBConnMaxLifetime)
	db.SetConnMaxIdleTime(config.DBConnMaxIdleTime)
	db.SetMaxIdleConns(config.DBConnMaxIdle)
	db.SetMaxOpenConns(config.DBConnMaxOpen)

	return db
}
