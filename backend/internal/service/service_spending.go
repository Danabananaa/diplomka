package service

import (
	"diplomka/internal/model"
)

type spending struct {
	model.SpendingRepo
}

func NewSpendingService(sr model.SpendingRepo) *spending {
	return &spending{
		SpendingRepo: sr,
	}
}

// func (s *Spending) GetAllSpendingRepo(ctx context.Context) ([]*model.SpendingType, error) {
// 	obj, err := s.SpendingRepo.GetSpendingType(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("get spending repo: %v", err)
// 	}

// 	return obj, nil
// }
