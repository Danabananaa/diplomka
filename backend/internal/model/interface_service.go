package model

import (
	"context"
)

type AuthService interface {
	AuthenticationService
}

type AuthenticationService interface {
	SignUp(ctx context.Context, user User) error
	LogIn(ctx context.Context, auth Authentication) (*Token, error)
	Refresh(ctx context.Context, token string) (string, error)
	JWTService
}

type JWTService interface {
	GenerateJWT(context.Context, User) (*Token, error)
	Verification(signedToken string) (int64, error)
}

type FinancialTrackerService interface {
	IncomeService
	SpendingService
	StatisticsService
	FinancialTypeService
	FinancialLiabilityService
}

type FinancialTypeService interface {
	GetAllIncomeSpendingService(ctx context.Context) ([]*IncomeType, []*SpendingType, error)
}

type IncomeService interface {
	InsertIncomeService(ctx context.Context, inc Income) (*Income, error)
	GetIncomeService(ctx context.Context, inc Between) ([]*Income, error)
}

type SpendingService interface {
	InsertSpendingService(ctx context.Context, spn Spending) (*Spending, error)
	GetSpendingService(ctx context.Context, bet Between) ([]*Spending, error)
}

type StatisticsService interface {
	GetStatistics(context.Context, Between) (Statistics, error)
}

type FinancialLiabilityService interface {
	GetAllLoanDebtTypeService(ctx context.Context) ([]*LoanDebtType, error)
	AddLoanDebtService(ctx context.Context, asl Loan_Debt) (*Loan_Debt, error)
	GetLoanDebtService(ctx context.Context, bet Between) ([]*Loan_Debt, []*Loan_Debt, error)
}
