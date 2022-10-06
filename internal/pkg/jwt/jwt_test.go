package jwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	jwtManager := getJwtManage()
	token, err := jwtManager.Generate(123)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, token)
	fmt.Println(token)
}

func TestValidate(t *testing.T) {
	jwtManager := getJwtManage()

	token, err := jwtManager.Generate(111)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, token)

	claims, err := jwtManager.Validate(token)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, claims)

	claims, err = jwtManager.Validate("wrong token")
	assert.NotEmpty(t, err)
}

func getJwtManage() JwtManager {
	return NewJwtManager(&Config{
		Secret:         "secret",
		ExpirationTime: 100,
	})
}
