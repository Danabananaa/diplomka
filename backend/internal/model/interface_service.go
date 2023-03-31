package model

import "context"

type AuthSerivice interface {
	SignUp(context.Context, User) error
	LogOut()
	LogIn(ctx context.Context, email, password string) (string, error)
	Refresh(ctr context.Context, token string) (string, error)
}

type SpendingService interface {
	GetAllSpendingRepo(ctx context.Context) ([]*SpendingType, error)
}
