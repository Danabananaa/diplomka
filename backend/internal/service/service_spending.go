package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type spending struct {
	model.SpendingRepo
}

func NewSpendingService(spn model.SpendingRepo) *spending {
	return &spending{
		SpendingRepo: spn,
	}
}

func (s *spending) InsertSpendingService(ctx context.Context, spn model.Spending) (*model.Spending, error) {
	obj, err := s.SpendingRepo.AddSpending(ctx, spn)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}

func (s *spending) GetSpendingService(ctx context.Context, bet model.Between) ([]*model.Spending, error) {
	obj, err := s.SpendingRepo.GetSpending(ctx, bet)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}
