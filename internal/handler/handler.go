package handler

import (
	"golang-postgres/internal/handler/user"
	"golang-postgres/pkg/repository"
)

type Handler struct {
	User *user.UserHandler
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		User: user.NewUserHandler(repo),
	}
}
