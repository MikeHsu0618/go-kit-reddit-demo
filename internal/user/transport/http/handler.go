package http

import (
	"context"
	"encoding/json"
	endpoint2 "go-kit-reddit-demo/internal/user/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeCreateHandler creates the handler logic
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

// makeLoginHandler creates the handler logic
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
