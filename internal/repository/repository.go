package repository

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository/mdb"
	"gorm.io/gorm"
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
type Repository struct {
	Posts    Posts
	Comments Comments
	Users    Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Posts:    mdb.NewPostsRepository(db),
		Comments: mdb.NewCommentsRepository(db),
		Users:    mdb.NewUsersRepository(db),
	}
}
