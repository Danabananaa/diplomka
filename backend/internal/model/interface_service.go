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

// type SpendingTypeService interface {
// 	GetAllSpendingRepo(ctx context.Context) ([]*SpendingType, error)
// }

type IncomeSpendTypeService interface {
	GetAllIncSpendingService(ctx context.Context) ([]*IncomeType, []*SpendingType, error)
}

type IncomeService interface {
	InsertIncomeService(ctx context.Context, inc Income) (*Income, error)
	GetIncomeService(ctx context.Context, inc Between) ([]*Income, error)
}

type SpendingService interface {
	InsertSpendingService(ctx context.Context, spn Spending) (*Spending, error)
	GetSpendingService(ctx context.Context, bet Between) ([]*Spending, error)
}

type LoanDebtTypeService interface {
	GetAllLoanDebtTypeService(ctx context.Context) ([]*LoanDebtType, error)
}

type LoanDebtService interface {
	AddLoanDebtService(ctx context.Context, asl Loan_Debt) (*Loan_Debt, error)
	GetLoanDebtService(ctx context.Context, bet Between) ([]*Loan_Debt, []*Loan_Debt, error)
}

type StatisticsService interface {
	GetStatistics(context.Context, Between) (Statistics, error)
}
