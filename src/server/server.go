//Package server provides all server configuration for expose endpoints of the micro service
package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"superheroe-gokit-api/src/endpoint"
	"superheroe-gokit-api/src/entity"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.GetByID,
		decodeIDRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/superheroes").Handler(httptransport.NewServer(
		endpoints.GetAll,
		DecodeEmptyRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/superheroes").Handler(httptransport.NewServer(
		endpoints.Add,
		decodeRequest,
		encodeResponse,
	))
	r.Methods("PUT").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.Edit,
		decodeIDBodyRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.Delete,
		decodeIDRequest,
		encodeResponse,
	))

	return r
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req entity.Superheroe
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req entity.GetIDRequest
	vars := mux.Vars(r)

	req = entity.GetIDRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeIDBodyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req entity.Superheroe
	vars := mux.Vars(r)

	// field := reflect.ValueOf(vars["id"]).Field(0)
	id := vars["id"]
	if id == "" {
		return nil, errors.New("id is required")
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	req.ID = id

	return req, nil
}

func DecodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var appRequest entity.EmptyStruct
	return &appRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
