package repository

import (
	"github.com/jmoiron/sqlx"
)

type session struct {
	DB *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) *session {
	return &session{}
}
