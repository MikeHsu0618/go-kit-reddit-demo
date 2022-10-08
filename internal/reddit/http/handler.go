package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint2 "go-kit-reddit-demo/internal/reddit/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

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
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

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
