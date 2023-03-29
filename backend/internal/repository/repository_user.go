package repository

import (
	"github.com/jmoiron/sqlx"
)

type user struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *user {
	return &user{}
}

func (u *user) GetUser() {
}
