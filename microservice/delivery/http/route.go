package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bibit/microservice/domain"
	"github.com/bibit/microservice/model"
	"github.com/gorilla/mux"
)

type httpHandler struct {
	svc domain.ServiceAPI
}

const (
	post = http.MethodPost
	get  = http.MethodGet
)

func httpResponseWrite(rw http.ResponseWriter, response interface{}, statusCode int) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)
	_ = json.NewEncoder(rw).Encode(response)
}

func decodeRequest(r *http.Request, item interface{}) error {
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		log.Println("[error getting request]", err)
		return err
	}

	return nil
}

func initServer(apps domain.ServiceAPI) httpHandler {
	return httpHandler{
		svc: apps,
	}
}

func NewHandler(m *mux.Router, apps domain.ServiceAPI) http.Handler {
	svr := initServer(apps)
	m.HandleFunc("/v/1/movie/list", svr.GetMovieHandler).Methods(get)

	return m
}

func (c *httpHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	var p model.GetMovieRequest
	err := decodeRequest(r, &p)
	if err != nil {
		httpResponseWrite(w, nil, http.StatusBadRequest)
		return
	}

	resp, err := c.svc.GetListMovie(r.Context(), &p)
	if err != nil {
		httpResponseWrite(w, nil, http.StatusInternalServerError)
		return
	}

	httpResponseWrite(w, resp, http.StatusOK)
}
