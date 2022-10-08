package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	"go-kit-reddit-demo/internal/user/entity"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, username string, pwd string) (res string, err error) {
	defer func() {
		l.logger.Log("method", "Create", "username", username, "pwd", pwd, "res", res, "err", err)
	}()
	return l.next.Create(ctx, username, pwd)
}
func (l loggingMiddleware) Login(ctx context.Context, username string, pwd string) (user *entity.User, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "pwd", pwd, "user", user, "err", err)
	}()
	return l.next.Login(ctx, username, pwd)
}
