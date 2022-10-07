package repository

import (
	"context"
	"go-kit-reddit-demo/internal/pkg/pg"
	"go-kit-reddit-demo/internal/user/pkg/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id uint64) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Where("username = ?", username).First(user).Error
	return user, err
}

func (r *userRepository) Get(ctx context.Context, id uint64) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Where("id = ?", id).First(user).Error
	return user, err
}
