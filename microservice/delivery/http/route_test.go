package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bibit/microservice/infrastructure/omdb"
	"github.com/bibit/microservice/service"
	"github.com/stretchr/testify/assert"
)

func TestListMovie(t *testing.T) {
	requestBody := `{
		"movie": "Batman",
		"page": "2"
	}`
	req, err := http.NewRequest("GET", "/v/1/movie/list", bytes.NewBuffer([]byte(requestBody)))
	assert.Equal(t, nil, err)

	movieAPI := omdb.NewOmdbService(
		"http://www.omdbapi.com",
		"faf7e5bb",
	)

	API := service.NewServiceAPI(movieAPI)

	server := initServer(API)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetMovieHandler)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestListMovieBadRequest(t *testing.T) {
	requestBody := `{
		"movie": "Batman",
		"page": "2"
	`
	req, err := http.NewRequest("GET", "/v/1/movie/list", bytes.NewBuffer([]byte(requestBody)))
	assert.Equal(t, nil, err)

	movieAPI := omdb.NewOmdbService(
		"http://www.omdbapi.com",
		"faf7e5bb",
	)

	API := service.NewServiceAPI(movieAPI)

	server := initServer(API)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetMovieHandler)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestListMovieInternalError(t *testing.T) {
	requestBody := `{
		"movies": "Batman",
		"page": "2"
	}`
	req, err := http.NewRequest("GET", "/v/1/movie/list", bytes.NewBuffer([]byte(requestBody)))
	assert.Equal(t, nil, err)

	movieAPI := omdb.NewOmdbService(
		"http://www.omdbapi.com",
		"faf7e5bb",
	)

	API := service.NewServiceAPI(movieAPI)

	server := initServer(API)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetMovieHandler)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
