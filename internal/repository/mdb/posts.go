package mdb

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"gorm.io/gorm"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

func (r *PostsRepository) GetAll() ([]models.Post, error) {
	var lists []models.Post

	result := r.db.Find(&lists)

	return lists, result.Error
}

func (r *PostsRepository) Create(post *models.Post) error {
	result := r.db.Create(post)
	return result.Error
}
