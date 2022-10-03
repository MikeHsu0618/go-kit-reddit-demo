package service

import "context"

// AuthService describes the service.
type AuthService interface {
	GenerateToken(ctx context.Context, id int) (res string, err error)
	ValidateToken(ctx context.Context, id int) (res string, err error)
	RefreshToken(ctx context.Context, id int) (res string, err error)
}

type basicAuthService struct{}

func (b *basicAuthService) GenerateToken(ctx context.Context, id int) (res string, err error) {
	res = "1233"
	return res, err
}
func (b *basicAuthService) ValidateToken(ctx context.Context, id int) (res string, err error) {
	// TODO implement the business logic of ValidateToken
	return res, err
}
func (b *basicAuthService) RefreshToken(ctx context.Context, id int) (res string, err error) {
	// TODO implement the business logic of RefreshToken
	return res, err
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
