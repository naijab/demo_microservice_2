package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"post_service/post/model"
)

type IPostRepository interface {
	GetById(id string) (model.Post, error)
	GetAll() ([]model.Post, error)
}

type PostRepository struct {
	DB *gorm.DB
}

// Create New Instance of PostRepository
func NewPostRepository(DB *gorm.DB) *PostRepository {
	return &PostRepository{DB}
}

func (r *PostRepository) GetById(id string) (model.Post, error) {
	var p model.Post
	var e error
	if err := r.DB.Where("id = ?", id).First(&p).Error; err != nil {
		e = errors.New("post not found")
	}
	return p, e
}

func (r *PostRepository) GetAll() ([]model.Post, error) {
	var p []model.Post
	var e error
	if err := r.DB.Find(&p).Error; err != nil {
		e = errors.New("post not found")
	}
	return p, e
}
