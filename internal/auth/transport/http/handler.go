package http

import (
	"context"
	"encoding/json"
	endpoint2 "go-kit-reddit-demo/internal/auth/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeGenerateTokenHandler
func makeGenerateTokenHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/generate-token").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GenerateTokenEndpoint, decodeGenerateTokenRequest, encodeGenerateTokenResponse, options...)))
}

func decodeGenerateTokenRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.GenerateTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeGenerateTokenResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeValidateTokenHandler
func makeValidateTokenHandler(m *mux.Router, endpoints endpoint2.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/validate-token").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ValidateTokenEndpoint, decodeValidateTokenRequest, encodeValidateTokenResponse, options...)))
}

func decodeValidateTokenRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint2.ValidateTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeValidateTokenResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint2.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
