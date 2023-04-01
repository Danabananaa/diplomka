package repository

import (
	"github.com/jmoiron/sqlx"
)

type test struct {
	DB *sqlx.DB
}

func NewTestRepo(db *sqlx.DB) *test {
	return &test{
		DB: db,
	}
}
