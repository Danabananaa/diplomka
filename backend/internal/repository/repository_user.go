package repository

import (
	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type user struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) model.UserRepo {
	return &user{}
}

func (u *user) GetUser() {
}
