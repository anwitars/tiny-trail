package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Input JSON for the shorten endpoint.
type ShortenRequest struct {
	URL string `json:"url"`
}

// Output JSON for the shorten endpoint.
type ShortenedURL struct {
	OriginalURL string `json:"original_url"`
	ShortID     string `json:"short_id"`
}

// Shortens a URL and returns a short URL.
func ShortenEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)
	var req ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error in request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	shortenedURL := ShortenedURL{
		OriginalURL: req.URL,
	}

	err = db.QueryRowx("INSERT INTO trails (original_url) VALUES ($1) RETURNING id", shortenedURL.OriginalURL).Scan(&shortenedURL.ShortID)
	if err != nil {
		http.Error(w, "Error inserting into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(shortenedURL.ShortID))
}
