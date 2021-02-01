package service

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
)

type Posts interface {
	GetAll() ([]models.Post, error)
	Create(post *models.Post) error
}

type Comments interface {
}

type Service struct {
	Posts    Posts
	Comments Comments
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Posts:    NewPostsService(repository.Posts),
		Comments: NewCommentsService(repository.Comments),
	}
}
