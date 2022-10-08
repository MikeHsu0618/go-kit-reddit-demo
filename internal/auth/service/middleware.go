package service

import (
	"context"
	jwt "go-kit-reddit-demo/internal/pkg/jwt"

	log "github.com/go-kit/kit/log"
)

type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GenerateToken(ctx context.Context, id uint64) (token string, err error) {
	defer func() {
		l.logger.Log("method", "GenerateToken", "id", id, "token", token, "err", err)
	}()
	return l.next.GenerateToken(ctx, id)
}

func (l loggingMiddleware) ValidateToken(ctx context.Context, token string) (claims *jwt.UserClaims, err error) {
	defer func() {
		l.logger.Log("method", "ValidateToken", "token", token, "claims", claims, "err", err)
	}()
	return l.next.ValidateToken(ctx, token)
}
