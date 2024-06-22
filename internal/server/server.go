package server

import (
	"context"
	"net/http"
	"tinytrail/internal/endpoints"
	"tinytrail/internal/middlewares"

	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	DB *sqlx.DB
}

func (appContext *AppContext) HandleWithDB(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "db", appContext.DB)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func registerEndpoint(pattern string, handler http.Handler, middlewares ...middlewares.Middleware) {
	h := handler

	for _, m := range middlewares {
		h = m(h)
	}

	http.Handle(pattern, h)
}

func RegisterEndpoints(appContext *AppContext) {
	registerEndpoint("/", http.HandlerFunc(endpoints.HelloEndpoint), middlewares.Logger)

	registerEndpoint("POST /shorten", appContext.HandleWithDB(http.HandlerFunc(endpoints.ShortenEndpoint)), middlewares.Logger)
}
