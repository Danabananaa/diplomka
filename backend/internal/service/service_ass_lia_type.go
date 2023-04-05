package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type assliatype struct {
	model.AssLiaTypeRepo
}

func NewAssLiaTypeService(alt model.AssLiaTypeRepo) *assliatype {
	return &assliatype{
		AssLiaTypeRepo: alt,
	}
}

func (alt *assliatype) GetAllAssLiaTypeService(ctx context.Context) ([]*model.AssLiaType, error) {
	obj, err := alt.AssLiaTypeRepo.GetAssLiaTypeRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
