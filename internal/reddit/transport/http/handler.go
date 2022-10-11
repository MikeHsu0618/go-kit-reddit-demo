package http

import (
	"context"
	"encoding/json"
	endpoint2 "go-kit-reddit-demo/internal/reddit/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeLoginHandler
func makeLoginHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/login").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)))
}

func decodeLoginRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeLoginResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreatePostHandler
func makeCreatePostHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create-post").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(
			endpoints.CreatePostEndpoint,
			decodeCreatePostRequest,
			encodeCreatePostResponse, options...,
		)))
}

func decodeCreatePostRequest(ctx context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.CreatePostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeCreatePostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListPostHandler
func makeListPostHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/list-post").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListPostEndpoint, decodeListPostRequest, encodeListPostResponse, options...)))
}

func decodeListPostRequest(ctx context.Context, r *http1.Request) (interface{}, error) {
	return endpoint2.ListPostRequest{}, nil
}

func encodeListPostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRegisterHandler
func makeRegisterHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/register").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)))
}

func decodeRegisterRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeRegisterResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
