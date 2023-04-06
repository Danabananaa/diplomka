package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type loan_debt_type struct {
	DB *sqlx.DB
}

func NewLoanDebtTypeRepo(db *sqlx.DB) *loan_debt_type {
	return &loan_debt_type{
		DB: db,
	}
}

func (alt *loan_debt_type) GetLoandebtTypeRepo(ctx context.Context) ([]*model.LoanDebtType, error) {
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
