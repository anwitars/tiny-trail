package endpoints

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Redirects the user to the original URL based on the shortened URL ID.
func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)
	shortenedURLID := r.PathValue("shortenedURLID")

	originalURL := ""
	var expired bool

	if err := db.QueryRowx("SELECT original_url, (expiration < CURRENT_TIMESTAMP) AS expired FROM trails WHERE id = $1", shortenedURLID).Scan(&originalURL, &expired); err != nil {
		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if originalURL == "" || expired {
		http.Error(w, "URL was either not found or has been expired", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
