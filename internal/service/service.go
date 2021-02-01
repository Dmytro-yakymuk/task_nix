package service

import "github.com/Dmytro-yakymuk/task_nix/internal/repository"

type Posts interface {
}

type Comments interface {
}

type Service struct {
	Posts    Posts
	Comments Comments
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Posts:    NewPostsService(repository),
		Comments: NewCommentsService(repository),
	}
}
