// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "go-kit-reddit-demo/internal/reddit/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	LoginEndpoint      endpoint.Endpoint
	CreatePostEndpoint endpoint.Endpoint
	ListPostEndpoint   endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.RedditService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreatePostEndpoint: MakeCreatePostEndpoint(s),
		ListPostEndpoint:   MakeListPostEndpoint(s),
		LoginEndpoint:      MakeLoginEndpoint(s),
	}
	for _, m := range mdw["Login"] {
		eps.LoginEndpoint = m(eps.LoginEndpoint)
	}
	for _, m := range mdw["CreatePost"] {
		eps.CreatePostEndpoint = m(eps.CreatePostEndpoint)
	}
	for _, m := range mdw["ListPost"] {
		eps.ListPostEndpoint = m(eps.ListPostEndpoint)
	}
	return eps
}