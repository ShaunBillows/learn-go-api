package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	DbConnectionString = "DB_CONNECTION_STRING"

	ErrLoadingEnv     = ConnectionErr("error loading .env file")
	ErrConnectingToDb = ConnectionErr("failed to connect database")
)

type ConnectionErr string

func (e ConnectionErr) Error() string {
	return string(e)
}

func Connect() (*gorm.DB, error) {

	// Get data source name from env variable
	err := godotenv.Load()
	if err != nil {
		return nil, ErrLoadingEnv
	}
	dsn := os.Getenv(DbConnectionString)

	// Connect to MySQL db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, ErrConnectingToDb
	}

	return db, nil
}
