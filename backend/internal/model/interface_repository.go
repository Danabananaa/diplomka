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
	GetSpending(ctx context.Context, bet Between) ([]*Spending, error)
}

type IncomeRepo interface {
	AddIncome(ctx context.Context, inc Income) (*Income, error)
	GetIncome(context.Context, Between) ([]*Income, error)
}

type SpendingTypeRepo interface {
	GetSpendingType(ctx context.Context) ([]*SpendingType, error)
}

type IncomeTypeRepo interface {
	GetIncomeType(ctx context.Context) ([]*IncomeType, error)
}

type AssLiaTypeRepo interface {
	GetAssLiaType(ctx context.Context) ([]*AssLiaType, error)
}

type AssetsRepo interface {
	AddAssets(ctx context.Context, ast Assets) (*Assets, error)
}

type LiabilitiesRepo interface {
	AddLiabilities(ctx context.Context, lia Liabilities) (*Liabilities, error)
}
