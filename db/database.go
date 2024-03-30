package db

import (
	"database/sql"
	"gkru-service/helper"
	"time"

	"github.com/sirupsen/logrus"
)

func NewDB(logger *logrus.Logger) *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gkru_app")
	// if db cloesed after service run
	if err != nil {
        logger.WithError(err).Error("Failed to connect to database")
        helper.PanicIfError(err)
    }
	
	// cek if db already connected
	if err := db.Ping(); err != nil {
        logger.WithError(err).Error("Failed to connect to database")
        panic(err)
    }

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}