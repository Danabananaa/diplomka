package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type loan_dept_type struct {
	model.LoanDebtTypeRepo
}

func NewLoanDebtTypeService(ldt model.LoanDebtTypeRepo) *loan_dept_type {
	return &loan_dept_type{
		LoanDebtTypeRepo: ldt,
	}
}

func (ldt *loan_dept_type) GetAllLoanDebtTypeService(ctx context.Context) ([]*model.LoanDebtType, error) {
	obj, err := ldt.LoanDebtTypeRepo.GetLoandebtTypeRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
