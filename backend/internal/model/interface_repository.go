package model

import (
	"context"
)

type UserRepo interface {
	AddUser(context.Context, User) (*User, error)
	GetUserforAuth(ctx context.Context, auth Authentication) (*User, error)
}

type SpendingRepo interface {
	AddSpending(ctx context.Context, inc Spending) (*Spending, error)
}

type IncomeRepo interface {
	AddIncome(ctx context.Context, inc Income) (*Income, error)
}

type SpendingTypeRepo interface {
	GetSpendingType(ctx context.Context) ([]*SpendingType, error)
}

type IncomeTypeRepo interface {
	GetIncomeType(ctx context.Context) ([]*IncomeType, error)
}
