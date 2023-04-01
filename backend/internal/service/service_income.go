package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
	"time"
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
	var err error
	inc.Date, err = time.Parse("01-02-2006 15:04", time.Now().Format("01-02-2006 15:04"))
	if err != nil {
		return nil, fmt.Errorf("Error was ocured during convertind time format %v", err)
	}

	obj, err := s.IncomeRepo.AddIncome(ctx, inc)
	if err != nil {
		return nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, nil
}
