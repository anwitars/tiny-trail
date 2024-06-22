package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"tinytrail/internal/config"
	"tinytrail/internal/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	//* Command line arguments and flags *
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		slog.Error(fmt.Sprintf("Invalid port: %d", *port))
		return
	}

	//* Load configuration and create database pool *
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

	//* Register endpoints and start server *
	server.RegisterEndpoints(appContext)

	slog.Info(fmt.Sprintf("Starting server on port %d", *port))

	address := fmt.Sprintf(":%d", *port)
	http.ListenAndServe(address, nil)
}
