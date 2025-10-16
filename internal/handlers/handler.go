package handlers

import (
	"database/sql"
	"url-shortener/internal/repository"
)

type Handler struct {
	repo *repository.LinkRepository
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{repo: repository.NewLinkRepository(db)}
}
