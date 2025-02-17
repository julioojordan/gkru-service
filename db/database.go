package db

import (
	"database/sql"
	"fmt"
	"gkru-service/helper"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func NewDB(logger *logrus.Logger) *sql.DB {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		logger.Warn("Gagal memuat file .env, menggunakan environment variables default")
	}

	// Ambil environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	maxIdleConnStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConn, err := strconv.Atoi(maxIdleConnStr)
	if err != nil {
		logger.Warn("Gagal mengonversi DB_MAX_IDLE_CONNS, menggunakan default: 15")
		maxIdleConn = 15
	}

	maxConnStr := os.Getenv("DB_MAX_CONNS")
	maxConn, err := strconv.Atoi(maxConnStr)
	if err != nil {
		logger.Warn("Gagal mengonversi DB_MAX_CONNS, menggunakan default: 50")
		maxConn = 50 
	}

	// Buat connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Koneksi ke database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.WithError(err).Error("Gagal untuk membuat koneksi ke database")
		helper.PanicIfError(err)
	}

	// Cek koneksi database
	if err := db.Ping(); err != nil {
		logger.WithError(err).Error("Gagal untuk membuat koneksi ke database")
		panic(err)
	}

	// Konfigurasi connection pooling
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(120 * time.Minute)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db
}
