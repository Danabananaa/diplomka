package model

import "context"

type UserRepo interface {
	AddUser(context.Context, User) (*User, error)
	GetUserforAuth(ctx context.Context, email, password string) (*User, error)
}

type SpendingRepo interface {
	GetSpendingType(ctx context.Context) ([]*SpendingType, error)
}

type IncomeRepo interface{}

type SpendingTypeRepo interface{}

type IncomeTypeRepo interface{}
