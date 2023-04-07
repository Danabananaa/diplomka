package service_liability

import (
	"context"
	"fmt"

	"diplomka/internal/model"
)

func (l *liability) AddLoanDebtService(ctx context.Context, asl model.Loan_Debt) (*model.Loan_Debt, error) {
	var obj *model.Loan_Debt
	var err error
	if asl.Status {
		obj, err = l.AddAssets(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}
	} else {
		obj, err = l.AddLiabilities(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}

	}

	return obj, nil
}

func (l *liability) GetLoanDebtService(ctx context.Context, bet model.Between) ([]*model.Loan_Debt, []*model.Loan_Debt, error) {
	obj, err := l.GetAssets(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}
	obj2, err := l.GetLiabilities(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, obj2, nil
}
