// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	"go-kit-reddit-demo/internal/auth/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GenerateTokenEndpoint endpoint.Endpoint
	ValidateTokenEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.AuthService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		GenerateTokenEndpoint: MakeGenerateTokenEndpoint(s),
		ValidateTokenEndpoint: MakeValidateTokenEndpoint(s),
	}
	for _, m := range mdw["GenerateToken"] {
		eps.GenerateTokenEndpoint = m(eps.GenerateTokenEndpoint)
	}
	for _, m := range mdw["ValidateToken"] {
		eps.ValidateTokenEndpoint = m(eps.ValidateTokenEndpoint)
	}
	return eps
}
