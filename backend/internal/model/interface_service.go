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
	GenerateJWT(context.Context, User) (*Token, error)
	Verification(signedToken string) (int64, error)
}

type SpendingService interface {
	GetAllSpendingRepo(ctx context.Context) ([]*SpendingType, error)
}

type IncomeTypeService interface {
	GetAllIncomeRepo(ctx context.Context) ([]*IncomeType, error)
}

type IncomeService interface {
	InsertIncomeService(ctx context.Context, inc Income) (*Income, error)
}
