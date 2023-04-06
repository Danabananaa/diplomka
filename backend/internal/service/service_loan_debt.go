package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type loan_debt struct {
	model.LoanRepo
	model.DeptRepo
}

func NewLoanDebtService(ass model.LoanRepo, liab model.DeptRepo) *loan_debt {
	return &loan_debt{
		LoanRepo: ass,
		DeptRepo: liab,
	}
}

func (alt *loan_debt) AddLoanDebtService(ctx context.Context, asl model.Loan_Debt) (*model.Loan_Debt, error) {
	var obj *model.Loan_Debt
	var err error
	if asl.Status {
		obj, err = alt.LoanRepo.AddAssets(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}
	} else {
		obj, err = alt.DeptRepo.AddLiabilities(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}

	}

	return obj, nil
}

func (alt *loan_debt) GetLoanDebtService(ctx context.Context, bet model.Between) ([]*model.Loan_Debt, []*model.Loan_Debt, error) {
	obj, err := alt.LoanRepo.GetAssets(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}
	obj2, err := alt.DeptRepo.GetLiabilities(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, obj2, nil
}
