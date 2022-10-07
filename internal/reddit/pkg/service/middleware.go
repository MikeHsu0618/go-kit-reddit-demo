package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	"go-kit-reddit-demo/internal/post/pkg/entity"
	user "go-kit-reddit-demo/internal/user/pkg/entity"
)

// Middleware describes a service middleware.
type Middleware func(RedditService) RedditService

type loggingMiddleware struct {
	logger log.Logger
	next   RedditService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a RedditService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next RedditService) RedditService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Login(ctx context.Context, username string, password string) (user *user.User, token string, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "password", "user", user, password, "token", token, "err", err)
	}()
	return l.next.Login(ctx, username, password)
}
func (l loggingMiddleware) CreatePost(ctx context.Context, title string, content string, userId uint64) (post *entity.Post, err error) {
	defer func() {
		l.logger.Log("method", "CreatePost", "title", title, "content", content, "userId", userId, "post", post, "err", err)
	}()
	return l.next.CreatePost(ctx, title, content, userId)
}
func (l loggingMiddleware) ListPost(ctx context.Context, userId uint64) (posts []*entity.Post, err error) {
	defer func() {
		l.logger.Log("method", "ListPost", "posts", posts, "err", err)
	}()
	return l.next.ListPost(ctx, userId)
}
