package endpoint

import (
	service "auth/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GenerateTokenRequest collects the request parameters for the GenerateToken method.
type GenerateTokenRequest struct {
	Id int `json:"id"`
}

// GenerateTokenResponse collects the response parameters for the GenerateToken method.
type GenerateTokenResponse struct {
	Res string `json:"res"`
	Err error  `json:"err"`
}

// MakeGenerateTokenEndpoint returns an endpoint that invokes GenerateToken on the service.
func MakeGenerateTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateTokenRequest)
		res, err := s.GenerateToken(ctx, req.Id)
		return GenerateTokenResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r GenerateTokenResponse) Failed() error {
	return r.Err
}

// ValidateTokenRequest collects the request parameters for the ValidateToken method.
type ValidateTokenRequest struct {
	Id int `json:"id"`
}

// ValidateTokenResponse collects the response parameters for the ValidateToken method.
type ValidateTokenResponse struct {
	Res string `json:"res"`
	Err error  `json:"err"`
}

// MakeValidateTokenEndpoint returns an endpoint that invokes ValidateToken on the service.
func MakeValidateTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateTokenRequest)
		res, err := s.ValidateToken(ctx, req.Id)
		return ValidateTokenResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r ValidateTokenResponse) Failed() error {
	return r.Err
}

// RefreshTokenRequest collects the request parameters for the RefreshToken method.
type RefreshTokenRequest struct {
	Id int `json:"id"`
}

// RefreshTokenResponse collects the response parameters for the RefreshToken method.
type RefreshTokenResponse struct {
	Res string `json:"res"`
	Err error  `json:"err"`
}

// MakeRefreshTokenEndpoint returns an endpoint that invokes RefreshToken on the service.
func MakeRefreshTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshTokenRequest)
		res, err := s.RefreshToken(ctx, req.Id)
		return RefreshTokenResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r RefreshTokenResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GenerateToken implements Service. Primarily useful in a client.
func (e Endpoints) GenerateToken(ctx context.Context, id int) (res string, err error) {
	request := GenerateTokenRequest{Id: id}
	response, err := e.GenerateTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GenerateTokenResponse).Res, response.(GenerateTokenResponse).Err
}

// ValidateToken implements Service. Primarily useful in a client.
func (e Endpoints) ValidateToken(ctx context.Context, id int) (res string, err error) {
	request := ValidateTokenRequest{Id: id}
	response, err := e.ValidateTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ValidateTokenResponse).Res, response.(ValidateTokenResponse).Err
}

// RefreshToken implements Service. Primarily useful in a client.
func (e Endpoints) RefreshToken(ctx context.Context, id int) (res string, err error) {
	request := RefreshTokenRequest{Id: id}
	response, err := e.RefreshTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RefreshTokenResponse).Res, response.(RefreshTokenResponse).Err
}
