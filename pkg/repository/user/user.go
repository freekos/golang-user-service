package user

import (
	"database/sql"
	"errors"
	_errors "golang-postgres/pkg/errors"
	"golang-postgres/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	DB *sqlx.DB
}

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(args *Config) *UserRepository {
	return &UserRepository{
		DB: args.DB,
	}
}

func (u *UserRepository) GetMany() ([]model.User, error) {
	users := make([]model.User, 0)
	err := u.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) GetOne(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Create(user *model.User) error {
	_, err := u.DB.Exec("INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)", user.FirstName, user.LastName, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Update(user *model.User) error {
	_, err := u.DB.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4", user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
