package endpoint

import (
	"go-kit-reddit-demo/internal/post/entity"
	user "go-kit-reddit-demo/internal/user/entity"
)

type Failure interface {
	Failed() error
}

type LoginResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"err"`
}

func (r LoginResponse) Failed() error {
	return r.Err
}

type CreatePostResponse struct {
	Post *entity.Post `json:"post"`
	Err  error        `json:"err"`
}

func (r CreatePostResponse) Failed() error {
	return r.Err
}

type ListPostResponse struct {
	Posts []*entity.Post `json:"posts"`
	Err   error          `json:"err"`
}

func (r ListPostResponse) Failed() error {
	return r.Err
}

type RegisterResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
	Err   error      `json:"err"`
}

func (r RegisterResponse) Failed() error {
	return r.Err
}
