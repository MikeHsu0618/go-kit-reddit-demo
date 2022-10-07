package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-kit-reddit-demo/internal/pkg/config"
	"go-kit-reddit-demo/internal/pkg/pg"
	"go-kit-reddit-demo/internal/post/pkg/entity"
	"math/rand"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	repo := getRepo()
	rand.Seed(time.Now().UnixNano())
	post := &entity.Post{
		UUID:    uuid.New().String(),
		Title:   "title",
		Content: "content",
		UserID:  uint64(rand.Int()),
	}
	err := repo.Create(context.Background(), post)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, post)
}

func TestGet(t *testing.T) {
	repo := getRepo()
	rand.Seed(time.Now().UnixNano())
	post := &entity.Post{
		UUID:    uuid.New().String(),
		Title:   "title",
		Content: "content",
		UserID:  uint64(rand.Int()),
	}
	err := repo.Create(context.Background(), post)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, post)
	res, err := repo.Get(context.Background(), post.ID)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, res)
}

func TestList(t *testing.T) {
	db := getDb()
	repo := getRepo()
	posts, err := repo.List(context.Background())
	assert.Equal(t, nil, err)

	var count int64
	err = db.Model(&entity.Post{}).Count(&count).Error
	assert.Equal(t, nil, err)
	assert.Equal(t, len(posts), int(count))
}

func getRepo() PostRepository {
	path := config.GetPath()
	conf, _ := config.Load(path)
	db, _ := pg.NewPostDB(conf)
	repo := NewPostRepository(db)
	return repo
}

func getDb() *pg.DB {
	path := config.GetPath()
	conf, _ := config.Load(path)
	db, _ := pg.NewPostDB(conf)
	return db
}
