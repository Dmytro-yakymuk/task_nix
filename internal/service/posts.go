package service

import "github.com/Dmytro-yakymuk/task_nix/internal/repository"

type PostsService struct {
	repository repository.Posts
}

func NewPostsService(repository repository.Posts) *PostsService {
	return &PostsService{repository: repository}
}
