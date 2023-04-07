package service

import (
	"diplomka/internal/model"
	"diplomka/internal/service/service_financial_type"
	"diplomka/internal/service/service_income"
	"diplomka/internal/service/service_liability"
	"diplomka/internal/service/service_spending"
	"diplomka/internal/service/service_statistics"
)

type financialTracker struct {
	model.IncomeService
	model.SpendingService
	model.StatisticsService
	model.FinancialTypeService
	model.FinancialLiabilityService
}

func NewFinancialTrackerService(f model.FinancialRepo) model.FinancialTrackerService {
	return &financialTracker{
		IncomeService:             service_income.NewIncomeService(f),
		SpendingService:           service_spending.NewSpendingService(f),
		StatisticsService:         service_statistics.NewStatisticsService(f),
		FinancialTypeService:      service_financial_type.NewFinancialTypeService(f),
		FinancialLiabilityService: service_liability.NewLoanDebtTypeService(f),
	}
}
