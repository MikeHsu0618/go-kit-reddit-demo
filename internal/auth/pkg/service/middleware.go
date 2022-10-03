package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GenerateToken(ctx context.Context, id int) (res string, err error) {
	defer func() {
		l.logger.Log("method", "GenerateToken", "id", id, "res", res, "err", err)
	}()
	return l.next.GenerateToken(ctx, id)
}
func (l loggingMiddleware) ValidateToken(ctx context.Context, id int) (res string, err error) {
	defer func() {
		l.logger.Log("method", "ValidateToken", "id", id, "res", res, "err", err)
	}()
	return l.next.ValidateToken(ctx, id)
}
func (l loggingMiddleware) RefreshToken(ctx context.Context, id int) (res string, err error) {
	defer func() {
		l.logger.Log("method", "RefreshToken", "id", id, "res", res, "err", err)
	}()
	return l.next.RefreshToken(ctx, id)
}
