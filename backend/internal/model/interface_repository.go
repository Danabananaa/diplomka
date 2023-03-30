package model

import "context"

type UserRepo interface {
	AddUser(context.Context, User) (*User, error)
	GetUserforAuth(ctx context.Context, email, password string) (*User, error)
}

type SessionRepo interface {
	AddSession(ctx context.Context, session Session) error
	IsSessionExist(ctx context.Context, session Session) (bool, error)
}

type SpendingRepo interface{}

type IncomeRepo interface{}

type SpendingTypeRepo interface{}

type IncomeTypeRepo interface{}
