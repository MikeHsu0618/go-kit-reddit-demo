// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "auth/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeGenerateTokenHandler(m, endpoints, options["GenerateToken"])
	makeValidateTokenHandler(m, endpoints, options["ValidateToken"])
	makeRefreshTokenHandler(m, endpoints, options["RefreshToken"])
	return m
}