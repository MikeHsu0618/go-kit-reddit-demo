package service

import (
	"context"
	"flag"
	"fmt"
	authHttp "go-kit-reddit-demo/internal/auth/client/http"
	logger2 "go-kit-reddit-demo/internal/pkg/logger"
	postHttp "go-kit-reddit-demo/internal/post/client/http"
	"go-kit-reddit-demo/internal/reddit/endpoint"
	service2 "go-kit-reddit-demo/internal/reddit/service"
	http1 "go-kit-reddit-demo/internal/reddit/transport/http"
	userHttp "go-kit-reddit-demo/internal/user/client/http"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger log.Logger
var fs = flag.NewFlagSet("reddit", flag.ExitOnError)
var httpAddr = fs.String("http-addr", ":8381", "HTTP listen address")

func Run() {
	fs.Parse(os.Args[1:])

	logger = logger2.NewLogger()

	authClient, err := authHttp.New("localhost:8081", nil)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	userClient, err := userHttp.New("localhost:8181", nil)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	postClient, err := postHttp.New("localhost:8281", nil)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	svc := service2.New(getServiceMiddleware(logger), authClient, userClient, postClient)
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger)
	// Add your http options here
	addOptions := []httptransport.ServerOption{
		// add Jwt Token to context
		httptransport.ServerBefore(func(ctx context.Context, r *http2.Request) context.Context {
			ctx = context.WithValue(ctx, "token", r.Header.Get("Authorization"))
			return ctx
		}),
	}
	for i, option := range options {
		options[i] = append(option, addOptions...)
	}
	httpHandler := http1.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service2.Middleware) {
	mw = []service2.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	// Append your middleware here
	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "reddit",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	// Add you endpoint middleware here

	return
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
