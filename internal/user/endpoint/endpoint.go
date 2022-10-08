package endpoint

import (
	"context"
	"github.com/gookit/validate"
	"go-kit-reddit-demo/internal/user/entity"
	"go-kit-reddit-demo/internal/user/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		res, err := s.Create(ctx, req.Username, req.Pwd)
		return CreateResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		v := validate.Struct(request)
		if !v.Validate() {
			return nil, v.Errors
		}
		res, err := s.Login(ctx, req.Username, req.Pwd)
		return LoginResponse{
			Err:  err,
			User: res,
		}, nil
	}
}

func (e Endpoints) Create(ctx context.Context, username string, pwd string) (res string, err error) {
	request := CreateRequest{
		Pwd:      pwd,
		Username: username,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Res, response.(CreateResponse).Err
}

func (e Endpoints) Login(ctx context.Context, username string, pwd string) (user *entity.User, err error) {
	request := LoginRequest{
		Pwd:      pwd,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).User, response.(LoginResponse).Err
}
