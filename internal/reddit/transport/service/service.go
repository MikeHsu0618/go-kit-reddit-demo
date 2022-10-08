package service

import (
	"context"
	authSvc "go-kit-reddit-demo/internal/auth/service"
	post "go-kit-reddit-demo/internal/post/entity"
	postSvc "go-kit-reddit-demo/internal/post/service"
	user "go-kit-reddit-demo/internal/user/entity"
	userSvc "go-kit-reddit-demo/internal/user/service"
	"strings"
)

// RedditService describes the service.
type RedditService interface {
	Login(ctx context.Context, username string, password string) (user *user.User, token string, err error)
	Register(ctx context.Context, username string, password string) (user *user.User, token string, err error)
	CreatePost(ctx context.Context, title string, content string, userId uint64) (post *post.Post, err error)
	ListPost(ctx context.Context) (posts []*post.Post, err error)
}

type basicRedditService struct {
	authClient authSvc.AuthService
	userClient userSvc.UserService
	postClient postSvc.PostService
}

func (b *basicRedditService) Login(ctx context.Context, username string, password string) (user *user.User, token string, err error) {
	user, err = b.userClient.Login(ctx, username, password)
	if err != nil {
		return user, "", err
	}
	token, err = b.authClient.GenerateToken(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, err
}

func (b *basicRedditService) Register(ctx context.Context, username string, password string) (user *user.User, token string, err error) {
	_, err = b.userClient.Create(ctx, username, password)
	if err != nil {
		return user, "", err
	}
	return b.Login(ctx, username, password)
}

func (b *basicRedditService) CreatePost(ctx context.Context, title string, content string, userId uint64) (post *post.Post, err error) {
	token := ctx.Value("token")
	if token == "" {
		return nil, Forbidden
	}
	return b.postClient.Create(ctx, title, content, userId)
}
func (b *basicRedditService) ListPost(ctx context.Context) (posts []*post.Post, err error) {
	token := ctx.Value("token")
	if token == nil {
		return nil, Forbidden
	}
	_, err = b.authClient.ValidateToken(ctx, strings.Split(token.(string), " ")[1])
	if err != nil {
		return nil, err
	}
	return b.postClient.List(ctx)
}

// NewBasicRedditService returns a naive, stateless implementation of RedditService.
func NewBasicRedditService(authClient authSvc.AuthService, userClient userSvc.UserService, postClient postSvc.PostService) RedditService {
	return &basicRedditService{
		authClient,
		userClient,
		postClient,
	}
}

// New returns a RedditService with all of the expected middleware wired in.
func New(middleware []Middleware, authClient authSvc.AuthService, userClient userSvc.UserService, postClient postSvc.PostService) RedditService {
	var svc RedditService = NewBasicRedditService(authClient, userClient, postClient)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
