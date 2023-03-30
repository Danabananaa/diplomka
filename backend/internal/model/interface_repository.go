package model

import "context"

type UserRepo interface {
	AddUser(context.Context, User) (*User, error)
}

type SessionRepo interface{}

type SpendingRepo interface{}

type IncomeRepo interface{}

type SpendingTypeRepo interface{}

type IncomeTypeRepo interface{}
