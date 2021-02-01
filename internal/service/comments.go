package service

import "github.com/Dmytro-yakymuk/task_nix/internal/repository"

type CommentsService struct {
	repository repository.Comments
}

func NewCommentsService(repository repository.Comments) *CommentsService {
	return &CommentsService{repository: repository}
}
