package main

import (
	"log/slog"
	"net/http"
	"tinytrail/internal/server"
)

func main() {
	server.RegisterEndpoints()

	slog.Info("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
