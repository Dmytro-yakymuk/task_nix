package mdb

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UsersRepository) GetOne(id string) (*models.User, error) {
	var user = new(models.User)
	result := r.db.First(&user, "id = ?", id)
	return user, result.Error
}

func (r *UsersRepository) Delete(id string) error {
	user := r.db.First(&models.User{}, "id = ?", id)
	if user.Error != nil {
		return user.Error
	}
	result := r.db.Delete(&models.User{}, "id = ?", id)
	return result.Error
}
