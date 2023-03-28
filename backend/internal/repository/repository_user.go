package repository

import (
	"database/sql"

	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type user struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sql.DB) model.UserRepo {
	return &user{}
}

func (u *user) GetUser() {
}
