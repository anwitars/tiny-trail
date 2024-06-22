package main

import (
	"log/slog"
	"net/http"
	"tinytrail/internal/config"
	"tinytrail/internal/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config: %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", appConfig.DatabaseURL)
	if err != nil {
		slog.Error("Error connecting to database: %v", err)
		return
	}

	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	appContext := &server.AppContext{
		DB: db,
	}

	server.RegisterEndpoints(appContext)

	slog.Info("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
