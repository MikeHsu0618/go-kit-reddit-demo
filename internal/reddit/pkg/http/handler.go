package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "go-kit-reddit-demo/internal/reddit/pkg/endpoint"
	http1 "net/http"
	"strconv"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/login").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)))
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	//bearerToken := r.Header.Get("Authorization")
	//token := strings.Split(bearerToken, " ")[1]

	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreatePostHandler creates the handler logic
func makeCreatePostHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create-post").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreatePostEndpoint, decodeCreatePostRequest, encodeCreatePostResponse, options...)))
}

// decodeCreatePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreatePostRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreatePostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreatePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreatePostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListPostHandler creates the handler logic
func makeListPostHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/list-post").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListPostEndpoint, decodeListPostRequest, encodeListPostResponse, options...)))
}

// decodeListPostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListPostRequest(ctx context.Context, r *http1.Request) (interface{}, error) {
	userId := r.URL.Query().Get("user_id")
	fmt.Println(userId)
	if len(userId) > 0 {
		fmt.Println(userId)
		if uId, err := strconv.ParseInt(userId, 10, 64); err == nil {
			fmt.Println(uint64(uId))
			fmt.Printf("%T, %v", uId, uId)
			return endpoint.ListPostRequest{
				UserId: uint64(uId),
			}, nil
		}
	}

	return endpoint.ListPostRequest{
		UserId: 0,
	}, nil
}

// encodeListPostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListPostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
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

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
