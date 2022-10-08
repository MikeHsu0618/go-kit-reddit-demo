// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	"go-kit-reddit-demo/internal/reddit/endpoint"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeLoginHandler(m, endpoints, options["Login"])
	makeRegisterHandler(m, endpoints, options["Register"])
	makeCreatePostHandler(m, endpoints, options["CreatePost"])
	makeListPostHandler(m, endpoints, options["ListPost"])
	return handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(m)
}
