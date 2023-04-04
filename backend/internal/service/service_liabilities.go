package service

// import (
// 	"context"
// 	"diplomka/internal/model"
// 	"fmt"
// )

// type liasbilities struct {
// 	model.LiabilitiesRepo
// }

// func NewLiabilitiesService(liab model.LiabilitiesRepo) *liasbilities {
// 	return &liasbilities{
// 		LiabilitiesRepo: liab,
// 	}
// }

// func (liab *liasbilities) AddLiabService(ctx context.Context, lia model.Liabilities) (*model.Liabilities, error) {
// 	obj, err := liab.LiabilitiesRepo.AddLiabilities(ctx, lia)
// 	if err != nil {
// 		return nil, fmt.Errorf("get spending repo: %v", err)
// 	}

// 	return obj, nil
// }
