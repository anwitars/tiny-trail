package middlewares

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// Middleware to inject the database connection pool into the request context.
func WithDatabase(db *sqlx.DB) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "db", db)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
