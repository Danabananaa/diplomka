package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

type assets_liab struct {
	model.AssetsRepo
	model.LiabilitiesRepo
}

func NewAssetsService(ass model.AssetsRepo, liab model.LiabilitiesRepo) *assets_liab {
	return &assets_liab{
		AssetsRepo:      ass,
		LiabilitiesRepo: liab,
	}
}

func (alt *assets_liab) AddAssetsLiabsService(ctx context.Context, asl model.Assets_or_Liabs) (*model.Assets_and_Liabs, error) {
	var obj *model.Assets_and_Liabs
	var err error
	if asl.Status {
		obj, err = alt.AssetsRepo.AddAssets(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}
	} else {
		obj, err = alt.LiabilitiesRepo.AddLiabilities(ctx, asl)
		if err != nil {
			return nil, fmt.Errorf("get spending repo: %v", err)
		}

	}

	return obj, nil
}

func (alt *assets_liab) GetAssetsLiabsService(ctx context.Context, bet model.Between) ([]*model.Assets, []*model.Liabilities, error) {
	obj, err := alt.AssetsRepo.GetAssets(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}
	obj2, err := alt.LiabilitiesRepo.GetLiabilities(ctx, bet)
	if err != nil {
		return nil, nil, fmt.Errorf("get income repo: %v", err)
	}

	return obj, obj2, nil
}
