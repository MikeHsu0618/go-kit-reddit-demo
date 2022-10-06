package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtManager interface {
	Generate(id uint64) (string, error)
	Validate(tokenStr string) (*UserClaims, error)
}

type jwtManager struct {
	config *Config
}

type UserClaims struct {
	UserId uint64 `json:"user_id,omitempty"`
	jwt.StandardClaims
}

type Config struct {
	Secret         string
	ExpirationTime time.Duration
}

func NewJwtManager(config *Config) JwtManager {
	return &jwtManager{
		config: config,
	}
}

func (j *jwtManager) Generate(id uint64) (string, error) {
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.config.ExpirationTime * time.Second).Unix(),
		},
	})
	return tokenObj.SignedString([]byte(j.config.Secret))
}

func (j *jwtManager) Validate(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(j.config.Secret), nil
		},
	)
	if err != nil {
		return &UserClaims{}, fmt.Errorf("parse token error %v", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return &UserClaims{}, fmt.Errorf("invalid token")
	}

	return claims, nil
}
