package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type spending_type struct {
	model.SpendingTypeRepo
}

func NewSpendingTypeService(sr model.SpendingTypeRepo) *spending_type {
	return &spending_type{
		SpendingTypeRepo: sr,
	}
}

func (s *spending_type) GetAllSpendingRepo(ctx context.Context) ([]*model.SpendingType, error) {
	obj, err := s.SpendingTypeRepo.GetSpendingType(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
