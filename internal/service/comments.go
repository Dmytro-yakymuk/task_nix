package service

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
)

type CommentsService struct {
	repository repository.Comments
}

func NewCommentsService(repository repository.Comments) *CommentsService {
	return &CommentsService{repository: repository}
}

func (s *CommentsService) GetAll() ([]models.Comment, error) {
	return s.repository.GetAll()
}

func (s *CommentsService) Create(comment *models.Comment) error {
	return s.repository.Create(comment)
}

func (s *CommentsService) GetOne(id int) (*models.Comment, error) {
	return s.repository.GetOne(id)
}

func (s *CommentsService) Update(comment *models.Comment) error {
	return s.repository.Update(comment)
}

func (s *CommentsService) Delete(id int) error {
	return s.repository.Delete(id)
}
