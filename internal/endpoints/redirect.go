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

	defer db.Close()

	originalURL := ""

	err := db.Get(&originalURL, "SELECT original_url FROM trails WHERE id = $1", shortenedURLID)
	if err != nil {
		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if originalURL == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
