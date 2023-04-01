package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type income_type struct {
	DB *sqlx.DB
}

func NewIncomeTypeRepo(db *sqlx.DB) *income_type {
	return &income_type{
		DB: db,
	}
}

func (i *income_type) GetIncomeType(ctx context.Context) ([]*model.IncomeType, error) {
	itarr := make([]*model.IncomeType, 0)
	query := `SELECT * FROM incometype`
	row, err := i.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error from QueryContext")
	}
	defer row.Close()

	for row.Next() {
		it := &model.IncomeType{}
		err := row.Scan(&it.ID, &it.IncomeType)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		itarr = append(itarr, it)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}

	return itarr, nil
}
