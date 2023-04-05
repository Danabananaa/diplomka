package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type income_spending_type struct {
	model.IncomeTypeRepo
	model.SpendingTypeRepo
}

func NewIncSpnTypeService(it model.IncomeTypeRepo, sp model.SpendingTypeRepo) *income_spending_type {
	return &income_spending_type{
		IncomeTypeRepo:   it,
		SpendingTypeRepo: sp,
	}
}

func (s *income_spending_type) GetAllIncSpendingService(ctx context.Context) ([]*model.IncomeType, []*model.SpendingType, error) {
	inctype, err := s.IncomeTypeRepo.GetIncomeType(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("get spending repo: %v", err)
	}
	spntype, err := s.SpendingTypeRepo.GetSpendingType(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("get spending repo: %v", err)
	}

	return inctype, spntype, nil
}
