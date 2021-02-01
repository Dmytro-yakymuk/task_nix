package repository

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/repository/mdb"
	"gorm.io/gorm"
)

type Posts interface {
}

type Comments interface {
}

type Repository struct {
	Posts    Posts
	Comments Comments
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Posts:    mdb.NewPostsRepository(db),
		Comments: mdb.NewCommentsRepository(db),
	}
}
