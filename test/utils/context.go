package utils

import (
	"errors"
	"fmt"
	"os"
	"tinytrail/internal/config"
	"tinytrail/internal/environment"
	"tinytrail/internal/server"

	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

func initTestDB() (*sqlx.DB, error) {
	configDir := os.Getenv(environment.WithPrefix("TEST_CONFIG_DIR"))
	if configDir == "" {
		return nil, errors.New(fmt.Sprintf("%s environment variable is not set", environment.WithPrefix("TEST_CONFIG_DIR")))
	}

	config, err := config.LoadConfig(configDir)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbConn = db

	return dbConn, nil
}

func GetTestDB() (*sqlx.DB, error) {
	if dbConn == nil {
		_, err := initTestDB()
		if err != nil {
			return nil, err
		}
	}

	return dbConn, nil
}

func NewTestAppContext() (*server.AppContext, error) {
	db, err := GetTestDB()
	if err != nil {
		return nil, err
	}

	return &server.AppContext{
		DB: db,
	}, nil
}
