package endpoint

import "go-kit-reddit-demo/internal/user/entity"

type Failure interface {
	Failed() error
}

type CreateResponse struct {
	Res string `json:"res"`
	Err error  `json:"err"`
}

func (r CreateResponse) Failed() error {
	return r.Err
}

type LoginResponse struct {
	User *entity.User `json:"user"`
	Err  error        `json:"err"`
}

func (r LoginResponse) Failed() error {
	return r.Err
}
