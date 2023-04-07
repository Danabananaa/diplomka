package service_liability

import (
	"context"
	"fmt"

	"diplomka/internal/model"
)

func (l *liability) GetAllLoanDebtTypeService(ctx context.Context) ([]*model.LoanDebtType, error) {
	obj, err := l.GetLoandebtTypeRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
