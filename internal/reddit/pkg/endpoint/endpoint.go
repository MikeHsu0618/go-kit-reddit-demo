package endpoint

import (
	"context"
	"github.com/gookit/validate"
	"go-kit-reddit-demo/internal/post/entity"
	service "go-kit-reddit-demo/internal/reddit/pkg/service"
	user "go-kit-reddit-demo/internal/user/pkg/entity"

	endpoint "github.com/go-kit/kit/endpoint"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"err"`
}

func MakeLoginEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		v := validate.Struct(request)
		if !v.Validate() {
			return nil, v.Errors
		}
		user, token, err := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			User:  user,
			Err:   err,
			Token: token,
		}, nil
	}
}

func (r LoginResponse) Failed() error {
	return r.Err
}

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	UserId  uint64 `json:"user_id" validate:"required"`
}

type CreatePostResponse struct {
	Post *entity.Post `json:"post"`
	Err  error        `json:"err"`
}

func MakeCreatePostEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePostRequest)
		v := validate.Struct(req)
		if !v.Validate() {
			return nil, v.Errors
		}
		post, err := s.CreatePost(ctx, req.Title, req.Content, req.UserId)
		return CreatePostResponse{
			Err:  err,
			Post: post,
		}, nil
	}
}

func (r CreatePostResponse) Failed() error {
	return r.Err
}

type ListPostRequest struct{}

type ListPostResponse struct {
	Posts []*entity.Post `json:"posts"`
	Err   error          `json:"err"`
}

func MakeListPostEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		posts, err := s.ListPost(ctx)
		return ListPostResponse{
			Err:   err,
			Posts: posts,
		}, nil
	}
}

func (r ListPostResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) Login(ctx context.Context, username string, password string) (token string, err error) {
	request := LoginRequest{
		Password: password,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Token, response.(LoginResponse).Err
}

func (e Endpoints) CreatePost(ctx context.Context, title string, content string, userId uint64) (post *entity.Post, err error) {
	request := CreatePostRequest{
		Content: content,
		Title:   title,
		UserId:  userId,
	}
	response, err := e.CreatePostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreatePostResponse).Post, response.(CreatePostResponse).Err
}

func (e Endpoints) ListPost(ctx context.Context, userId uint64) (posts []*entity.Post, err error) {
	request := ListPostRequest{}
	response, err := e.ListPostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListPostResponse).Posts, response.(ListPostResponse).Err
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"err"`
}

func MakeRegisterEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		v := validate.Struct(req)
		if !v.Validate() {
			return nil, v.Errors
		}
		user, token, err := s.Register(ctx, req.Username, req.Password)
		return RegisterResponse{
			Err:   err,
			Token: token,
			User:  user,
		}, nil
	}
}

func (r RegisterResponse) Failed() error {
	return r.Err
}

func (e Endpoints) Register(ctx context.Context, username string, password string) (user *user.User, token string, err error) {
	request := RegisterRequest{
		Password: password,
		Username: username,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).User, response.(RegisterResponse).Token, response.(RegisterResponse).Err
}
