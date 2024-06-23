package server

import (
	"net/http"
	"tinytrail/internal/endpoints"
	"tinytrail/internal/middlewares"

	"github.com/jmoiron/sqlx"
)

// AppContext holds the context of the application.
type AppContext struct {
	// DB is the database connection pool.
	DB *sqlx.DB
}

// Register an endpoint with a pattern, handler and optional middlewares.
// The middlewares are applied in the order they are passed.
func registerEndpoint(pattern string, handler http.Handler, middlewares ...middlewares.Middleware) {
	h := handler

	for _, m := range middlewares {
		h = m(h)
	}

	http.Handle(pattern, h)
}

// Register all endpoints that is used in the application.
func RegisterEndpoints(appContext *AppContext) {
	withDatabaseMiddleware := middlewares.WithDatabase(appContext.DB)

	registerEndpoint("POST /new", http.HandlerFunc(endpoints.NewTrailEndpoint), middlewares.Logger, withDatabaseMiddleware)
	registerEndpoint("GET /t/{shortenedURLID}", http.HandlerFunc(endpoints.RedirectEndpoint), middlewares.Logger, withDatabaseMiddleware)
	registerEndpoint("GET /list", http.HandlerFunc(endpoints.ListEndpoint), middlewares.Logger, withDatabaseMiddleware)
}
