package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type Spending struct {
	model.SpendingRepo
}

func NewSpendingService(sr model.SpendingRepo) *Spending {
	return &Spending{
		SpendingRepo: sr,
	}
}

func (s *Spending) GetAllSpendingRepo(ctx context.Context) ([]*model.SpendingType, error) {
	obj, err := s.SpendingRepo.GetSpendingType(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
