package model

import (
	"context"
)

type AuthSerivice interface {
	SignUp(context.Context, User) error
	LogOut(ctx context.Context) error
	LogIn(context.Context, Authentication) (*Token, error)
	Refresh(ctr context.Context, token string) (string, error)
	JWTService
}

type JWTService interface {
	GenerateJWT(context.Context, int64, string) (*Token, error)
	Verification(signedToken string) error
}

type SpendingService interface {
	GetAllSpendingRepo(ctx context.Context) ([]*SpendingType, error)
}

type IncomeService interface {
	GetAllIncomeRepo(ctx context.Context) ([]*IncomeType, error)
}
