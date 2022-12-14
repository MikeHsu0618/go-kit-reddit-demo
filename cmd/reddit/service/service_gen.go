// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	endpoint2 "go-kit-reddit-demo/internal/reddit/endpoint"
	"go-kit-reddit-demo/internal/reddit/service"
	http1 "go-kit-reddit-demo/internal/reddit/transport/http"
)

func createService(endpoints endpoint2.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"CreatePost": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger)},
		"ListPost":   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger)},
		"Login":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger)},
		"Register":   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger)},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Login"] = []endpoint1.Middleware{endpoint2.LoggingMiddleware(log.With(logger, "method", "Login")), endpoint2.InstrumentingMiddleware(duration.With("method", "Login"))}
	mw["Register"] = []endpoint1.Middleware{endpoint2.LoggingMiddleware(log.With(logger, "method", "Register")), endpoint2.InstrumentingMiddleware(duration.With("method", "Register"))}
	mw["CreatePost"] = []endpoint1.Middleware{endpoint2.LoggingMiddleware(log.With(logger, "method", "CreatePost")), endpoint2.InstrumentingMiddleware(duration.With("method", "CreatePost"))}
	mw["ListPost"] = []endpoint1.Middleware{endpoint2.LoggingMiddleware(log.With(logger, "method", "ListPost")), endpoint2.InstrumentingMiddleware(duration.With("method", "ListPost"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Login", "Register", "CreatePost", "ListPost"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
