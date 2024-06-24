package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"tinytrail/internal/endpoints"
	"tinytrail/internal/middlewares"
	"tinytrail/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestListEndpointEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	testAppContext, err := utils.NewTestAppContext()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := middlewares.Apply(http.HandlerFunc(endpoints.ListEndpoint), middlewares.WithDatabase(testAppContext.DB))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, "[]\n", rr.Body.String(), "handler returned unexpected body")
}

func TestListEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	testAppContext, err := utils.NewTestAppContext()
	if err != nil {
		t.Fatal(err)
	}

	originalURL := "http://example.com"
	var shortID string
	var expireAt time.Time
	err = testAppContext.DB.QueryRowx("INSERT INTO trails (original_url) VALUES ($1) RETURNING id, expiration", originalURL).Scan(&shortID, &expireAt)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := middlewares.Apply(http.HandlerFunc(endpoints.ListEndpoint), middlewares.WithDatabase(testAppContext.DB))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	var shortenedURLs []endpoints.ShortenedURL
	err = json.Unmarshal(rr.Body.Bytes(), &shortenedURLs)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(shortenedURLs), "shortened urls have unexpected length")

	shortenedURL := shortenedURLs[0]

	assert.Equal(t, originalURL, shortenedURL.OriginalURL, "shortened urls have unexpected original url")
	assert.Equal(t, shortID, shortenedURL.ShortID, "shortened urls have unexpected short id")
	assert.True(t, shortenedURL.Expiration.Equal(expireAt), "shortened urls have unexpected expiration")
}
