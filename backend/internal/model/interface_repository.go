package model

import (
	"context"
)

type UserRepo interface {
	AddUser(context.Context, User) (*User, error)
	GetUserforAuth(ctx context.Context, auth Authentication) (*User, error)
}
type FinancialRepo interface {
	SpendingRepo
	IncomeRepo
	SpendingTypeRepo
	IncomeTypeRepo
	StatisticsRepo
	DebtorRepo
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
type StatisticsRepo interface {
	GetStatistics(context.Context, Between) (Statistics, error)
}

type DebtorRepo interface {
	LoanDebtTypeRepo
	LoanRepo
	DeptRepo
}

type LoanDebtTypeRepo interface {
	GetLoandebtTypeRepo(ctx context.Context) ([]*LoanDebtType, error)
}

type LoanRepo interface {
	AddAssets(ctx context.Context, asl Loan_Debt) (*Loan_Debt, error)
	GetAssets(ctx context.Context, bet Between) ([]*Loan_Debt, error)
}

type DeptRepo interface {
	AddLiabilities(ctx context.Context, ast Loan_Debt) (*Loan_Debt, error)
	GetLiabilities(ctx context.Context, bet Between) ([]*Loan_Debt, error)
}
