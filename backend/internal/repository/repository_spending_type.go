package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type spending_type struct {
	DB *sqlx.DB
}

func NewSpendingTypeRepo(db *sqlx.DB) *spending_type {
	return &spending_type{
		DB: db,
	}
}

func (s *spending_type) GetSpendingType(ctx context.Context) ([]*model.SpendingType, error) {
	starr := make([]*model.SpendingType, 0)
	query := `SELECT * FROM spendingtype`
	row, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error from QueryContext")
	}
	defer row.Close()

	for row.Next() {
		st := &model.SpendingType{}
		err := row.Scan(&st.ID, &st.SpendingType)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		starr = append(starr, st)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}

	return starr, nil
}
