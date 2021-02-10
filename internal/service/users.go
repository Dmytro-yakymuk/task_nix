package service

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
)

type UsersService struct {
	repository repository.Users
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{repository: repository}
}

func (s *UsersService) Create(user *models.User) error {
	return s.repository.Create(user)
}

func (s *UsersService) GetOne(id string) (*models.User, error) {
	return s.repository.GetOne(id)
}
func (s *UsersService) Delete(id string) error {
	return s.repository.Delete(id)
}
