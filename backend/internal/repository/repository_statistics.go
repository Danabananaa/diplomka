package repository

import (
	"context"
	"log"

	"diplomka/internal/model"
)

func (s *repo) GetStatistics(ctx context.Context, b model.Between) (model.Statistics, error) {
	stat := model.Statistics{}

	querySumIncome := `
	SELECT sum(amount) AS sum
	FROM income WHERE user_id = ?
	AND date  BETWEEN ? AND ?`

	row := s.DB.QueryRowxContext(ctx, querySumIncome, b.UserID, b.StartDate, b.EndDate)
	err := row.StructScan(&stat.ValueIncome)
	if err != nil {
		return stat, err
	}

	querySumSpending := `
	SELECT ifnull(sum(amount), 0) AS sum
	FROM spending WHERE user_id = ?
	AND date  BETWEEN ? AND ?`

	row = s.DB.QueryRowxContext(ctx, querySumSpending, b.UserID, b.StartDate, b.EndDate)
	err = row.StructScan(&stat.ValueSpending)
	if err != nil {
		return stat, err
	}

	if stat.ValueIncome.TotalAmount > 0 {

		queryIncomePercentag := `SELECT
			income_type AS type_id,
			sum(amount) as 'total',
			(sum(amount) * 100.0) / ? as 'percentage'
		FROM income
		WHERE
			user_id = ?
		AND
			date BETWEEN ? AND ?
		GROUP by income_type;`

		rows, err := s.DB.QueryxContext(ctx,
			queryIncomePercentag,
			stat.ValueIncome.TotalAmount,
			b.UserID,
			b.StartDate,
			b.EndDate)
		if err != nil {
			return stat, err
		}

		for rows.Next() {
			v := model.Value{}
			err := rows.StructScan(&v)
			if err != nil {
				log.Println(err)
				return stat, err
			}
			stat.ValueIncome.Incomes = append(stat.ValueIncome.Incomes, v)
		}
	}

	if stat.ValueSpending.TotalAmount > 0 {
		querySpendingPercentag := `
			SELECT
			spending_type AS type_id,
			sum(amount) as 'total',
			(sum(amount) * 100.0) / ? as 'percentage'
		FROM spending
		WHERE
			user_id = ?
		AND
			date BETWEEN ? AND ?
		GROUP by spending_type;`

		rows, err := s.DB.QueryxContext(ctx,
			querySpendingPercentag,
			stat.ValueSpending.TotalAmount,
			b.UserID,
			b.StartDate,
			b.EndDate)
		if err != nil {
			return stat, err
		}

		for rows.Next() {
			v := model.Value{}
			err := rows.StructScan(&v)
			if err != nil {
				return stat, err
			}
			stat.ValueSpending.Spendings = append(stat.ValueSpending.Spendings, v)
		}
	}

	return stat, nil
}
