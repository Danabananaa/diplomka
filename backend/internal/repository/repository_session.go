package repository

import (
	"database/sql"

	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type session struct {
	DB *sqlx.DB
}

func NewSessionRepo(db *sql.DB) model.SessionRepo {
	return &session{}
}
