package repository

import (
	"context"
	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type user struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *user {
	return &user{DB: db}
}

func (u *user) AddUser(ctx context.Context, user model.User) (*model.User, error) {
	query := `INSERT INTO users (name, surname, email, password) VALUES (?,?,?,?);`

	res, err := u.DB.ExecContext(ctx, query, user.Name, user.Surname, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *user) GetUser() {
}
