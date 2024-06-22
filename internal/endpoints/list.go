package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Lists all the shortened URLs in the database.
func ListEndpoint(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sqlx.DB)

	rows, err := db.Query("SELECT original_url, short_id FROM shortened_urls")
	if err != nil {
		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	shortenedURLs := []ShortenedURL{}
	for rows.Next() {
		var shortenedURL ShortenedURL
		err := rows.Scan(&shortenedURL.OriginalURL, &shortenedURL.ShortID)
		if err != nil {
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}

		shortenedURLs = append(shortenedURLs, shortenedURL)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(shortenedURLs)
}
