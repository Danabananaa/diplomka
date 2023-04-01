package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type income_type struct {
	model.IncomeTypeRepo
}

func NewIncomeTypeService(it model.IncomeTypeRepo) *income_type {
	return &income_type{
		IncomeTypeRepo: it,
	}
}

func (s *income_type) GetAllIncomeRepo(ctx context.Context) ([]*model.IncomeType, error) {
	obj, err := s.IncomeTypeRepo.GetIncomeType(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
