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
	fmt.Printf("inc.Date.Time: %v\n", inc.Date.Time)
	obj, err := s.IncomeRepo.AddIncome(ctx, inc)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}
