package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"diplomka/internal/model"
)

func (alt *repo) GetLoandebtTypeRepo(ctx context.Context) ([]*model.LoanDebtType, error) {
	ldarr := make([]*model.LoanDebtType, 0)
	query := `SELECT * FROM loandebttype`
	row, err := alt.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error from QueryContext")
	}
	defer row.Close()

	for row.Next() {
		ld := &model.LoanDebtType{}
		err := row.Scan(&ld.ID, &ld.Type)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		ldarr = append(ldarr, ld)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}

	return ldarr, nil
}
