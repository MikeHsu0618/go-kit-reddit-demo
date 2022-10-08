package endpoint

import "go-kit-reddit-demo/internal/pkg/jwt"

type Failure interface {
	Failed() error
}

type ValidateTokenResponse struct {
	Claims *jwt.UserClaims `json:"claims"`
	Err    error           `json:"err"`
}

func (r ValidateTokenResponse) Failed() error {
	return r.Err
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

func (r GenerateTokenResponse) Failed() error {
	return r.Err
}
