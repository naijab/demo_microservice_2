package service

import (
	"post_service/post/model"
	"post_service/post/repository"
)

type IPostService interface {
	GetById(id string) (model.Post, error)
	GetAll() ([]model.Post, error)
}

type PostService struct {
	postRepo repository.IPostRepository
}

// Create New Instance of PostService
func NewPostService(postRepo repository.IPostRepository) *PostService {
	return  &PostService{postRepo}
}

func (s *PostService) GetById(id string) (model.Post, error) {
	return s.postRepo.GetById(id)
}

func (s *PostService) GetAll() ([]model.Post, error) {
	return s.postRepo.GetAll()
}
