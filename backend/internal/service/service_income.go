package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type income struct {
	model.IncomeRepo
}

func NewIncomeService(inc model.IncomeRepo) *income {
	return &income{
		IncomeRepo: inc,
	}
}

func (s *income) InsertIncomeService(ctx context.Context, inc model.Income) (*model.Income, error) {
	obj, err := s.IncomeRepo.AddIncome(ctx, inc)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}

func (s *income) GetIncomeService(ctx context.Context, bet model.Between) ([]*model.Income, error) {
	obj, err := s.IncomeRepo.GetIncome(ctx, bet)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}
