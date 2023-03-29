package repository

import (
	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type session struct {
	DB *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) model.SessionRepo {
	return &session{}
}
