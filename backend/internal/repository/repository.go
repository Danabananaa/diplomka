package repository

import "github.com/jmoiron/sqlx"

type repo struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repo {
	return &repo{
		DB: db,
	}
}
