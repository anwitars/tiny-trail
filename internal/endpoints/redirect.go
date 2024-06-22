package endpoints

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)
	shortenedURLID := r.PathValue("shortenedURLID")

	defer db.Close()

	var originalURL string
	err := db.Get(&originalURL, "SELECT original_url FROM shortened_urls WHERE short_id = $1", shortenedURLID)

	if err == nil {
		http.Redirect(w, r, originalURL, http.StatusFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
