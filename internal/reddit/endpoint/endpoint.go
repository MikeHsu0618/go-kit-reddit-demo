package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	"go-kit-reddit-demo/internal/post/entity"
	"go-kit-reddit-demo/internal/reddit/service"
	user "go-kit-reddit-demo/internal/user/entity"
)

// server

func MakeLoginEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		user, token, err := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			User:  user,
			Err:   err,
			Token: token,
		}, nil
	}
}

func MakeCreatePostEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePostRequest)
		post, err := s.CreatePost(ctx, req.Title, req.Content, req.UserId)
		return CreatePostResponse{
			Err:  err,
			Post: post,
		}, nil
	}
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

func MakeRegisterEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		user, token, err := s.Register(ctx, req.Username, req.Password)
		return RegisterResponse{
			Err:   err,
			Token: token,
			User:  user,
		}, nil
	}
}

// client

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
