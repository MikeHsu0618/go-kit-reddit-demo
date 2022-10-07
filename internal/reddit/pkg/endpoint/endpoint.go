package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	"go-kit-reddit-demo/internal/post/pkg/entity"
	service "go-kit-reddit-demo/internal/reddit/pkg/service"
	user "go-kit-reddit-demo/internal/user/pkg/entity"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
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

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// CreatePostRequest collects the request parameters for the CreatePost method.
type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint64 `json:"user_id"`
}

// CreatePostResponse collects the response parameters for the CreatePost method.
type CreatePostResponse struct {
	Post *entity.Post `json:"post"`
	Err  error        `json:"err"`
}

// MakeCreatePostEndpoint returns an endpoint that invokes CreatePost on the service.
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

// Failed implements Failer.
func (r CreatePostResponse) Failed() error {
	return r.Err
}

// ListPostRequest collects the request parameters for the ListPost method.
type ListPostRequest struct {
	UserId uint64 `json:"user_id"`
}

// ListPostResponse collects the response parameters for the ListPost method.
type ListPostResponse struct {
	Posts []*entity.Post `json:"posts"`
	Err   error          `json:"err"`
}

// MakeListPostEndpoint returns an endpoint that invokes ListPost on the service.
func MakeListPostEndpoint(s service.RedditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListPostRequest)
		posts, err := s.ListPost(ctx, req.UserId)
		return ListPostResponse{
			Err:   err,
			Posts: posts,
		}, nil
	}
}

// Failed implements Failer.
func (r ListPostResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
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

// CreatePost implements Service. Primarily useful in a client.
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

// ListPost implements Service. Primarily useful in a client.
func (e Endpoints) ListPost(ctx context.Context, userId uint64) (posts []*entity.Post, err error) {
	request := ListPostRequest{
		UserId: userId,
	}
	response, err := e.ListPostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListPostResponse).Posts, response.(ListPostResponse).Err
}
