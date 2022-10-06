package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type UserClaims struct {
	UserId uint64 `json:"user_id,omitempty"`
	jwt.StandardClaims
}

type Config struct {
	Secret         string
	ExpirationTime time.Duration
}

type JwtManager interface {
	Generate(id uint64) (string, error)
	Validate(tokenStr string) (*UserClaims, error)
}

type jwtManager struct {
	config *Config
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
			ExpiresAt: time.Now().Add(j.config.ExpirationTime).Unix(),
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
		return nil, fmt.Errorf("parse token error")
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("invalid token error")
			}
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token expired or not avaliable")
			}
		}
		return nil, fmt.Errorf("invalid token error")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
