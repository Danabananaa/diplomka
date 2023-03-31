package repository

import (
	"context"
	"diplomka/internal/model"

	"github.com/jmoiron/sqlx"
)

type spending_type struct {
	DB *sqlx.DB
}

func NewSpendingRepo(db *sqlx.DB) *spending_type {
	return &spending_type{
		DB: db,
	}
}

func (s *spending_type) GetSpendingType(ctx context.Context) ([]*model.SpendingType, error) {
	st := make([]*model.SpendingType, 0)
	// st := &model.Spendingtype{}
	query := `SELECT * FROM spendindtype`
	err := s.DB.SelectContext(ctx, nil, query, st)
	if err != nil {
		return nil, err
	}
	// if err := row.Err(); err != nil {
	// 	return nil, fmt.Errorf("error was ocured during executing method QueryRowContext: %v", err)
	// }
	// err := row.Scan(&st.ID, &st.SpendingType)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, sql.ErrNoRows
	// 	} else {
	// 		return nil, err
	// 	}
	// }
	return st, nil
}
