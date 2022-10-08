package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-kit-reddit-demo/internal/user/entity"
	"go-kit-reddit-demo/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService describes the service.
type UserService interface {
	Create(ctx context.Context, username string, pwd string) (res string, err error)
	Login(ctx context.Context, username string, pwd string) (user *entity.User, err error)
}

type basicUserService struct {
	repo repository.UserRepository
}

func (b *basicUserService) Create(ctx context.Context, username string, pwd string) (res string, err error) {
	password, err := generateFromPassword(pwd)
	if err != nil {
		return "", err
	}
	user := &entity.User{
		UUID:     uuid.New().String(),
		Username: username,
		Password: password,
	}
	err = b.repo.Create(ctx, user)
	return "success", err
}

func (b *basicUserService) Login(ctx context.Context, username string, pwd string) (user *entity.User, err error) {
	user, err = b.repo.GetByUsername(ctx, username)
	if err != nil {
		return user, errors.New("user not found")
	}

	if !isCorrectPassword(user.Password, pwd) {
		return user, errors.New("wrong password")
	}

	return user, nil
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService(repo repository.UserRepository) UserService {
	return &basicUserService{
		repo,
	}
}

func generateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func isCorrectPassword(password string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	return err == nil
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware, repo repository.UserRepository) UserService {
	var svc UserService = NewBasicUserService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
