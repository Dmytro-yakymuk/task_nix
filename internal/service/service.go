package service

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
)

type Posts interface {
	GetAll() ([]models.Post, error)
	Create(post *models.Post) error
	GetOne(id int) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id int) error
}

type Comments interface {
	GetAll() ([]models.Comment, error)
	Create(comment *models.Comment) error
	GetOne(id int) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id int) error
}

type Users interface {
	Create(user *models.User) error
	GetOne(id string) (*models.User, error)
	Delete(id string) error
}

type Service struct {
	Posts    Posts
	Comments Comments
	Users    Users
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Posts:    NewPostsService(repository.Posts),
		Comments: NewCommentsService(repository.Comments),
		Users:    NewUsersService(repository.Users),
	}
}
