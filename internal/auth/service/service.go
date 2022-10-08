package service

import (
	"context"
	"go-kit-reddit-demo/internal/pkg/jwt"
)

type AuthService interface {
	GenerateToken(ctx context.Context, id uint64) (token string, err error)
	ValidateToken(ctx context.Context, token string) (claims *jwt.UserClaims, err error)
}

type basicAuthService struct {
	jwtManager jwt.JwtManager
}

func (b *basicAuthService) GenerateToken(ctx context.Context, id uint64) (token string, err error) {
	return b.jwtManager.Generate(id)
}

func (b *basicAuthService) ValidateToken(ctx context.Context, token string) (claims *jwt.UserClaims, err error) {

	return b.jwtManager.Validate(token)
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService(jwtManager jwt.JwtManager) AuthService {
	return &basicAuthService{
		jwtManager: jwtManager,
	}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware, jwtManager jwt.JwtManager) AuthService {
	var svc AuthService = NewBasicAuthService(jwtManager)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
