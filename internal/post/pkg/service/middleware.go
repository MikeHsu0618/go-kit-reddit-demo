package service

import (
	"context"
	"go-kit-reddit-demo/internal/post/pkg/entity"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(PostService) PostService

type loggingMiddleware struct {
	logger log.Logger
	next   PostService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PostService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PostService) PostService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, title string, content string, userId uint64) (res *entity.Post, err error) {
	defer func() {
		l.logger.Log("method", "Create", "title", title, "content", content, "userId", userId, "res", res, "err", err)
	}()
	return l.next.Create(ctx, title, content, userId)
}
func (l loggingMiddleware) List(ctx context.Context) (res []*entity.Post, err error) {
	defer func() {
		l.logger.Log("method", "List", "res", res, "err", err)
	}()
	return l.next.List(ctx)
}

func (l loggingMiddleware) ListById(ctx context.Context, id uint64) (res []*entity.Post, err error) {
	defer func() {
		l.logger.Log("method", "ListById", "id", id, "res", res, "err", err)
	}()
	return l.next.ListById(ctx, id)
}
