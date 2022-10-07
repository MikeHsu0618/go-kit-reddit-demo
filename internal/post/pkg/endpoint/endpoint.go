package endpoint

import (
	"context"
	"go-kit-reddit-demo/internal/post/pkg/entity"
	service "go-kit-reddit-demo/internal/post/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/gookit/validate"
)

type CreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	UserId  uint64 `json:"user_id" validate:"required"`
}

type CreateResponse struct {
	Res *entity.Post `json:"res"`
	Err error        `json:"err"`
}

func MakeCreateEndpoint(s service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		v := validate.Struct(request)
		if !v.Validate() {
			return nil, v.Errors
		}
		res, err := s.Create(ctx, req.Title, req.Content, req.UserId)
		return CreateResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r CreateResponse) Failed() error {
	return r.Err
}

type ListRequest struct{}

type ListResponse struct {
	Res []*entity.Post `json:"res"`
	Err error          `json:"err"`
}

func MakeListEndpoint(s service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.List(ctx)
		return ListResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r ListResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) Create(ctx context.Context, title string, content string, userId uint64) (res *entity.Post, err error) {
	request := CreateRequest{
		Content: content,
		Title:   title,
		UserId:  userId,
	}

	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Res, response.(CreateResponse).Err
}

func (e Endpoints) List(ctx context.Context) (res []*entity.Post, err error) {
	request := ListRequest{}
	response, err := e.ListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListResponse).Res, response.(ListResponse).Err
}

type ListByIdRequest struct {
	Id uint64 `json:"id"`
}

type ListByIdResponse struct {
	Res []*entity.Post `json:"res"`
	Err error          `json:"err"`
}

func MakeListByIdEndpoint(s service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListByIdRequest)
		res, err := s.ListById(ctx, req.Id)
		return ListByIdResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

func (r ListByIdResponse) Failed() error {
	return r.Err
}

func (e Endpoints) ListById(ctx context.Context, id uint64) (res []*entity.Post, err error) {
	request := ListByIdRequest{Id: id}
	v := validate.Struct(request)
	if !v.Validate() {
		return nil, err
	}
	response, err := e.ListByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListByIdResponse).Res, response.(ListByIdResponse).Err
}
