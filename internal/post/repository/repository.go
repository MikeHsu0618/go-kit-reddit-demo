package repository

import (
	"context"
	"go-kit-reddit-demo/internal/pkg/pg"
	"go-kit-reddit-demo/internal/post/entity"
)

type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) error
	Get(ctx context.Context, id uint64) (*entity.Post, error)
	List(ctx context.Context) ([]*entity.Post, error)
	ListById(ctx context.Context, id uint64) ([]*entity.Post, error)
}

type postRepository struct {
	db *pg.DB
}

func NewPostRepository(db *pg.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(ctx context.Context, post *entity.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Get(ctx context.Context, id uint64) (*entity.Post, error) {
	post := &entity.Post{}
	err := r.db.Where("id = ?", id).First(post).Error
	return post, err
}

func (r *postRepository) List(ctx context.Context) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *postRepository) ListById(ctx context.Context, id uint64) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.Where("user_id = ?", id).Find(&posts).Error
	return posts, err
}
