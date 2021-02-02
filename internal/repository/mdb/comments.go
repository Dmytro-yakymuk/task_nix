package mdb

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"gorm.io/gorm"
)

type CommentsRepository struct {
	db *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

func (r *CommentsRepository) GetAll() ([]models.Comment, error) {
	var comments []models.Comment
	result := r.db.Find(&comments)
	return comments, result.Error
}

func (r *CommentsRepository) Create(comment *models.Comment) error {
	result := r.db.Create(comment)
	return result.Error
}

func (r *CommentsRepository) GetOne(id int) (*models.Comment, error) {
	var comment = new(models.Comment)
	result := r.db.First(&comment, id)
	return comment, result.Error
}

func (r *CommentsRepository) Update(comment *models.Comment) error {
	check_comment := r.db.First(&models.Comment{}, comment.Id)
	if check_comment.Error != nil {
		return check_comment.Error
	}

	result := r.db.Model(comment).Updates(models.Comment{PostId: comment.PostId, Name: comment.Name, Email: comment.Email, Body: comment.Body})
	return result.Error
}

func (r *CommentsRepository) Delete(id int) error {
	Comment := r.db.First(&models.Comment{}, id)
	if Comment.Error != nil {
		return Comment.Error
	}
	result := r.db.Delete(&models.Comment{}, id)
	return result.Error
}
