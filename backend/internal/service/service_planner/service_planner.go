package service_planner

import (
	"context"
	"fmt"
	"log"

	"diplomka/internal/model"
)

type planner struct {
	model.PlannerRepo
}

func New(r model.PlannerRepo) *planner {
	return &planner{
		PlannerRepo: r,
	}
}

func (p *planner) AddSpendingPLanner(ctx context.Context, planner model.SpendingPlanner) error {
	log.Println("yes")

	err := p.PlannerRepo.AddPlannerSpending(ctx, planner)
	if err != nil {
		return fmt.Errorf("planner repo: %v", err)
	}

	return nil
}
