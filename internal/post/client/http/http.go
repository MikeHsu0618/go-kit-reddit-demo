package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpoint1 "go-kit-reddit-demo/internal/post/pkg/endpoint"
	http2 "go-kit-reddit-demo/internal/post/pkg/http"
	service "go-kit-reddit-demo/internal/post/pkg/service"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.PostService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = http.NewClient("POST", copyURL(u, "/create"), encodeHTTPGenericRequest, decodeCreateResponse, options["Create"]...).Endpoint()
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = http.NewClient("GET", copyURL(u, "/list"), encodeHTTPGenericRequest, decodeListResponse, options["List"]...).Endpoint()
	}

	var listByIdEndpoint endpoint.Endpoint
	{
		listByIdEndpoint = http.NewClient("GET", copyURL(u, "/list-by-id"), encodeHTTPGenericRequest, decodeListByIdResponse, options["ListById"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateEndpoint:   createEndpoint,
		ListByIdEndpoint: listByIdEndpoint,
		ListEndpoint:     listEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeCreateResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
