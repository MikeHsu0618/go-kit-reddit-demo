package http

import (
	"context"
	"encoding/json"
	endpoint2 "go-kit-reddit-demo/internal/post/endpoint"
	http1 "net/http"
	"strconv"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeCreateHandler
func makeCreateHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)))
}

func decodeCreateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeCreateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListHandler
func makeListHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("GET").
		Path("/list").
		Handler(
			handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
				handlers.AllowedOrigins([]string{"*"}),
			)(http.NewServer(endpoints.ListEndpoint, decodeListRequest, encodeListResponse, options...)))
}

func decodeListRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.ListRequest{}
	return req, nil
}

func encodeListResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListByIdHandler
func makeListByIdHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/list-by-id").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListByIdEndpoint, decodeListByIdRequest, encodeListByIdResponse, options...)))
}

func decodeListByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.ListByIdRequest{}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return nil, err
	}
	req.Id = uint64(id)
	return req, err
}

func encodeListByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
