package endpoint

import (
	"context"
	"go-kit-reddit-demo/internal/auth/service"
	jwt "go-kit-reddit-demo/internal/pkg/jwt"

	endpoint "github.com/go-kit/kit/endpoint"
)

type GenerateTokenRequest struct {
	Id uint64 `json:"id"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

func MakeGenerateTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateTokenRequest)
		token, err := s.GenerateToken(ctx, req.Id)
		return GenerateTokenResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

func (r GenerateTokenResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) GenerateToken(ctx context.Context, id uint64) (token string, err error) {
	request := GenerateTokenRequest{Id: id}
	response, err := e.GenerateTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GenerateTokenResponse).Token, response.(GenerateTokenResponse).Err
}

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	Claims *jwt.UserClaims `json:"claims"`
	Err    error           `json:"err"`
}

func MakeValidateTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateTokenRequest)
		claims, err := s.ValidateToken(ctx, req.Token)
		return ValidateTokenResponse{
			Claims: claims,
			Err:    err,
		}, nil
	}
}

func (r ValidateTokenResponse) Failed() error {
	return r.Err
}

func (e Endpoints) ValidateToken(ctx context.Context, token string) (claims *jwt.UserClaims, err error) {
	request := ValidateTokenRequest{Token: token}
	response, err := e.ValidateTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ValidateTokenResponse).Claims, response.(ValidateTokenResponse).Err
}
