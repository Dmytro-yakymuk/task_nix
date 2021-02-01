package service

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
)

type PostsService struct {
	repository repository.Posts
}

func NewPostsService(repository repository.Posts) *PostsService {
	return &PostsService{repository: repository}
}

func (s *PostsService) GetAll() ([]models.Post, error) {
	return s.repository.GetAll()
}

func (s *PostsService) Create(post *models.Post) error {
	return s.repository.Create(post)
}
