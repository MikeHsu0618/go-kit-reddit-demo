package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-kit-reddit-demo/internal/pkg/config"
	"go-kit-reddit-demo/internal/pkg/pg"
	"go-kit-reddit-demo/internal/user/entity"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	repo := getRepo()
	rand.Seed(time.Now().UnixNano())
	err := repo.Create(context.Background(), &entity.User{
		UUID:     uuid.New().String(),
		Username: strconv.Itoa(rand.Int()),
		Password: "AA",
	})
	assert.Equal(t, nil, err)
}

func TestGetByLogin(t *testing.T) {
	repo := getRepo()
	rand.Seed(time.Now().UnixNano())
	user := &entity.User{
		UUID:     uuid.New().String(),
		Username: strconv.Itoa(rand.Int()),
		Password: "AA",
	}
	err := repo.Create(context.Background(), user)
	assert.Equal(t, nil, err)

	res, err := repo.GetByUsername(context.Background(), user.Username)
	assert.Equal(t, nil, err)
	assert.Equal(t, res.Username, user.Username)
	assert.Equal(t, res.Password, user.Password)
}

func TestGet(t *testing.T) {
	repo := getRepo()
	rand.Seed(time.Now().UnixNano())
	user := &entity.User{
		UUID:     uuid.New().String(),
		Username: strconv.Itoa(rand.Int()),
		Password: "AA",
	}
	err := repo.Create(context.Background(), user)
	assert.Equal(t, nil, err)
	res, err := repo.Get(context.Background(), user.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, user.ID, res.ID)
}

func getRepo() UserRepository {
	path := config.GetPath()
	conf, _ := config.Load(path)
	db, _ := pg.NewUserDB(conf)
	repo := NewUserRepository(db)
	return repo
}
