package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"diplomka/internal/model"
)

func (u *repo) AddUser(ctx context.Context, user model.User) (*model.User, error) {
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

func (u *repo) GetUser(ctx context.Context, id int) (*model.User, error) {
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

func (u *repo) DeleteUser(ctx context.Context, userID int64) error {
	tx, err := u.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `DELETE FROM spending WHERE user_id = ?`, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM income WHERE user_id = ?`, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM loan WHERE user_id = ?`, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM debt WHERE user_id = ?`, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM images WHERE user_id = ?`, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *repo) GetUserforAuth(ctx context.Context, auth model.Authentication) (*model.User, error) {
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

func (u *repo) AddUserImage(ctx context.Context, info model.UserImage) (*model.UserImage, error) {
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

func (u *repo) GetUserImage(ctx context.Context, id int64) (*model.UserImage, error) {
	query := `Select * from images where user_id=?`

	row := u.DB.QueryRowxContext(ctx, query, id)

	info := model.UserImage{}
	err := row.Scan(&info.UserID, &info.ImageName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no avatar found for user") // Return a custom error
		} else {
			return nil, err
		}
	}

	return &info, nil
}

func (u *repo) DeleteImage(ctx context.Context, id int64) error {
	query := `DELETE FROM images WHERE user_id=?`
	_, err := u.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}
	return nil
}
