package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type assets struct {
	model.AssetsRepo
}

func NewAssetsService(ass model.AssetsRepo) *assets {
	return &assets{
		AssetsRepo: ass,
	}
}

func (alt *assets) AddAssetsService(ctx context.Context, ass model.Assets) (*model.Assets, error) {
	obj, err := alt.AssetsRepo.AddAssets(ctx, ass)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("get spending repo: %v", err)
	}

	return obj, nil
}
