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

type SpendingTypeService interface {
	GetAllSpendingRepo(ctx context.Context) ([]*SpendingType, error)
}

type IncomeTypeService interface {
	GetAllIncomeRepo(ctx context.Context) ([]*IncomeType, error)
}

type IncomeService interface {
	InsertIncomeService(ctx context.Context, inc Income) (*Income, error)
	GetIncomeService(ctx context.Context, inc Between) ([]*Income, error)
}

type SpendingService interface {
	InsertSpendingService(ctx context.Context, spn Spending) (*Spending, error)
	GetSpendingService(ctx context.Context, bet Between) ([]*Spending, error)
}

type AssLiaTypeService interface {
	GetAllAssLiaTypeService(ctx context.Context) ([]*AssLiaType, error)
}

type AssetsLiabService interface {
	AddAssetsLiabsService(ctx context.Context, asl Assets_or_Liabs) (*Assets_and_Liabs, error)
	GetAssetsLiabsService(ctx context.Context, bet Between) ([]*Assets, []*Liabilities, error)
}
