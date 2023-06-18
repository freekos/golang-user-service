package repository

import (
	"golang-postgres/pkg/repository/user"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User *user.UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: user.NewUserRepository(&user.Config{DB: db}),
	}
}
