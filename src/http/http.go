//Package server provides all server configuration for expose endpoints of the micro service
package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"superheroe-gokit-api/src/endpoint"
	"superheroe-gokit-api/src/entity"
	"superheroe-gokit-api/src/util"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPServer creates a new HTTP server with the routes and handlers
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/health").Handler(httptransport.NewServer(
		endpoints.Health,
		DecodeEmptyRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.GetByID,
		decodeIDRequest,
		encodeSuperheroeResponse,
	))
	r.Methods("GET").Path("/superheroes").Handler(httptransport.NewServer(
		endpoints.GetAll,
		DecodeEmptyRequest,
		encodeSuperheroesResponse,
	))
	r.Methods("POST").Path("/superheroes").Handler(httptransport.NewServer(
		endpoints.Add,
		decodeRequest,
		encodeResponse,
	))
	r.Methods("PUT").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.Edit,
		decodeIDBodyRequest,
		encodeSuperheroeResponse,
	))
	r.Methods("DELETE").Path("/superheroes/{id}").Handler(httptransport.NewServer(
		endpoints.Delete,
		decodeIDRequest,
		encodeSuperheroeResponse,
	))

	return r
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req *entity.Superheroe
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req entity.IDRequest
	vars := mux.Vars(r)

	req = entity.IDRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeIDBodyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req *entity.Superheroe
	vars := mux.Vars(r)

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
	var appRequest interface{}
	return &appRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeSuperheroesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	f, ok := response.(*entity.Superheroes)
	if ok && f.Error != nil {
		status, msg := util.Error2Wrapper(f.Error)
		w.WriteHeader(status)
		return json.NewEncoder(w).Encode(msg)
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeSuperheroeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	f, ok := response.(*entity.SuperheroeResponse)
	if ok && f.Error != nil {
		status, msg := util.Error2Wrapper(f.Error)
		w.WriteHeader(status)
		return json.NewEncoder(w).Encode(msg)
	}
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		/* var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(util.InputError{
			Message: "Todo Mal",
		})
		w.WriteHeader(http.StatusNotFound)
		w.Write(buf.Bytes()) */

		next.ServeHTTP(w, r)
	})
}
