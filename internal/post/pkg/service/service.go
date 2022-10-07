package service

import (
	"context"
	"go-kit-reddit-demo/internal/post/pkg/entity"
	"go-kit-reddit-demo/internal/post/pkg/repository"

	"github.com/google/uuid"
)

// PostService describes the service.
type PostService interface {
	Create(ctx context.Context, title string, content string, userId uint64) (res *entity.Post, err error)
	List(ctx context.Context) (res []*entity.Post, err error)
	ListById(ctx context.Context, id uint64) (res []*entity.Post, err error)
}

type basicPostService struct {
	repo repository.PostRepository
}

func (b *basicPostService) Create(ctx context.Context, title string, content string, userId uint64) (res *entity.Post, err error) {
	post := &entity.Post{
		UUID:    uuid.New().String(),
		Title:   title,
		Content: content,
		UserID:  userId,
	}
	err = b.repo.Create(ctx, post)
	return post, err
}
func (b *basicPostService) List(ctx context.Context) (res []*entity.Post, err error) {
	return b.repo.List(ctx)
}

func (b *basicPostService) ListById(ctx context.Context, id uint64) (res []*entity.Post, err error) {
	return b.repo.ListById(ctx, id)
}

// NewBasicPostService returns a naive, stateless implementation of PostService.
func NewBasicPostService(repo repository.PostRepository) PostService {
	return &basicPostService{
		repo,
	}
}

// New returns a PostService with all of the expected middleware wired in.
func New(middleware []Middleware, repo repository.PostRepository) PostService {
	var svc PostService = NewBasicPostService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
