package model

import "context"

type AuthSerivice interface {
	SignUp(context.Context, User) error
	LogOut()
	LogIn()
	Refresh()
}

type Analysis interface{}
