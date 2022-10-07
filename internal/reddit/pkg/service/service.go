package service

import (
	"context"
	authSvc "go-kit-reddit-demo/internal/auth/pkg/service"
	postSvc "go-kit-reddit-demo/internal/post/pkg/service"
	userSvc "go-kit-reddit-demo/internal/user/pkg/service"

	post "go-kit-reddit-demo/internal/post/pkg/entity"
	user "go-kit-reddit-demo/internal/user/pkg/entity"
)

// RedditService describes the service.
type RedditService interface {
	Login(ctx context.Context, username string, password string) (user *user.User, token string, err error)
	CreatePost(ctx context.Context, title string, content string, userId uint64) (post *post.Post, err error)
	ListPost(ctx context.Context, userId uint64) (posts []*post.Post, err error)
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
func (b *basicRedditService) CreatePost(ctx context.Context, title string, content string, userId uint64) (post *post.Post, err error) {
	return b.postClient.Create(ctx, title, content, userId)
}
func (b *basicRedditService) ListPost(ctx context.Context, userId uint64) (posts []*post.Post, err error) {
	if userId == 0 {
		return b.postClient.List(ctx)
	}
	return b.postClient.ListById(ctx, userId)
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
