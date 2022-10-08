package endpoint

import (
	"context"
	"github.com/gookit/validate"
	"go-kit-reddit-demo/internal/user/entity"
	"go-kit-reddit-demo/internal/user/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Res string `json:"res"`
	Err error  `json:"err"`
}

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

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Pwd      string `json:"pwd" validate:"required"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	User *entity.User `json:"user"`
	Err  error        `json:"err"`
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

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
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

// Login implements Service. Primarily useful in a client.
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
