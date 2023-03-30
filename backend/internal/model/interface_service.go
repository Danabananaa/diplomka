package model

import "context"

type AuthSerivice interface {
	SignUp(context.Context, User) error
	LogOut()
	LogIn(ctx context.Context, email, password string) (string, error)
	Refresh()
}

type Analysis interface{}
