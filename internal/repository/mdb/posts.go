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
	var posts []models.Post
	result := r.db.Find(&posts)
	return posts, result.Error
}

func (r *PostsRepository) Create(post *models.Post) error {
	result := r.db.Create(post)
	return result.Error
}

func (r *PostsRepository) GetOne(id int) (*models.Post, error) {
	var post = new(models.Post)
	result := r.db.First(&post, id)
	return post, result.Error
}

func (r *PostsRepository) Update(post *models.Post) error {
	check_post := r.db.First(&models.Post{}, post.Id)
	if check_post.Error != nil {
		return check_post.Error
	}

	result := r.db.Model(post).Updates(models.Post{UserId: post.UserId, Title: post.Title, Body: post.Body})
	return result.Error
}

func (r *PostsRepository) Delete(id int) error {
	post := r.db.First(&models.Post{}, id)
	if post.Error != nil {
		return post.Error
	}
	result := r.db.Delete(&models.Post{}, id)
	return result.Error
}
