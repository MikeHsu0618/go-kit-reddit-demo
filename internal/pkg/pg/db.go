package pg

import (
	"fmt"

	"go-kit-reddit-demo/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	params := " sslmode=disable TimeZone=Asia/Taipei"
	dsn = dsn + params
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func NewUserDB(conf *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v",
		conf.User.DB.Host,
		conf.User.DB.User,
		conf.User.DB.Password,
		conf.User.DB.Name,
		conf.User.DB.Port,
	)
	return NewDB(dsn)
}

func NewPostDB(conf *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v",
		conf.Post.DB.Host,
		conf.Post.DB.User,
		conf.Post.DB.Password,
		conf.Post.DB.Name,
		conf.Post.DB.Port,
	)
	return NewDB(dsn)
}
