package repository

import (
	"context"
	"database/sql"

	"diplomka/internal/model"
)

func (r *repo) AddPlannerSpending(ctx context.Context, planner model.SpendingPlanner) error {
	query := `INSERT INTO planner (user_id,spending_type_id,amount) VALUES (:user_id, :spending_type_id, :amount);`
	res, err := r.DB.NamedExecContext(ctx, query, planner)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count <= 0 {
		return sql.ErrNoRows
	}

	return nil
}
