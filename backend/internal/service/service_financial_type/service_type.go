package service_financial_type

import (
	"context"
	"fmt"

	"diplomka/internal/model"
)

type financialType struct {
	model.FinancialRepo
}

func NewFinancialTypeService(f model.FinancialRepo) *financialType {
	return &financialType{
		FinancialRepo: f,
	}
}

func (s *financialType) GetAllIncomeSpendingService(ctx context.Context) ([]*model.IncomeType, []*model.SpendingType, error) {
	inctype, err := s.FinancialRepo.GetIncomeType(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("get spending repo: %v", err)
	}
	spntype, err := s.FinancialRepo.GetSpendingType(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("get spending repo: %v", err)
	}

	return inctype, spntype, nil
}
