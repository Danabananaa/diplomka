package service_statistics

import (
	"context"
	"fmt"

	"diplomka/internal/model"
)

type stat struct {
	model.StatisticsRepo
}

func NewStatisticsService(s model.StatisticsRepo) *stat {
	return &stat{
		StatisticsRepo: s,
	}
}

func (s *stat) GetStatistics(ctx context.Context, b model.Between) (model.Statistics, error) {
	st, err := s.StatisticsRepo.GetStatistics(ctx, b)
	if err != nil {
		return st, fmt.Errorf("get statistics: %v", err)
	}
	return st, nil
}
