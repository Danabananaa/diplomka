package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

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
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &user, nil
}

func (u *user) GetUser(ctx context.Context, id int) (*model.User, error) {
	x := &model.User{}

	query := `select * from users where id=?`
	row := u.DB.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error was ocured during executing method QueryRowContext: %v", err)
	}

	err := row.Scan(&x.ID, &x.Name, &x.Surname, &x.Email, &x.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}
	return x, nil
}

func (u *user) GetUserforAuth(ctx context.Context, auth model.Authentication) (*model.User, error) {
	x := model.User{}
	query := `select * from users where email=? and password=?`
	row := u.DB.QueryRowContext(ctx, query, auth.Email, auth.Password)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error was ocured during executing method QueryRowContext: %v", err)
	}
	err := row.Scan(&x.ID, &x.Name, &x.Surname, &x.Email, &x.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}
	return &x, nil
}

func (u *user) AddUserImage(ctx context.Context, info model.UserImage) (*model.UserImage, error) {
	query := `INSERT INTO images (user_id, image_name) VALUES (?,?)`

	res, err := u.DB.ExecContext(ctx, query, info.UserID, info.ImageName)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	info.UserID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}
	return &info, nil
}

func (u *user) GetUserImage(ctx context.Context, id int) (*model.UserImage, error) {
	query := `Select * from images where user_id=?`

	row := u.DB.QueryRowxContext(ctx, query, id)

	info := model.UserImage{}
	err := row.Scan(&info.UserID, &info.ImageName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}

	return &info, nil
}
