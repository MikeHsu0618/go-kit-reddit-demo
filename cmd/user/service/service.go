package service

import (
	"flag"
	"fmt"
	"go-kit-reddit-demo/internal/pkg/config"
	logger2 "go-kit-reddit-demo/internal/pkg/logger"
	"go-kit-reddit-demo/internal/pkg/pg"
	"go-kit-reddit-demo/internal/user/endpoint"
	"go-kit-reddit-demo/internal/user/repository"
	service2 "go-kit-reddit-demo/internal/user/service"
	http1 "go-kit-reddit-demo/internal/user/transport/http"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	group "github.com/oklog/oklog/pkg/group"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
)

var logger log.Logger
var fs = flag.NewFlagSet("user", flag.ExitOnError)
var httpAddr = fs.String("http-addr", ":8181", "HTTP listen address")

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = logger2.NewLogger()

	cfg, err := config.Load("")
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	db, err := pg.NewUserDB(cfg)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	repo := repository.NewUserRepository(db)
	svc := service2.New(getServiceMiddleware(logger), repo)
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger)
	// Add your http options here

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
		Subsystem: "user",
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
