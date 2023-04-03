package repository

import (
	"context"
	"diplomka/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type spending struct {
	DB *sqlx.DB
}

func NewSpendingRepo(db *sqlx.DB) *spending {
	return &spending{
		DB: db,
	}
}

func (i *spending) AddSpending(ctx context.Context, spn model.Spending) (*model.Spending, error) {
	query := `INSERT INTO spending (user_id, spending_type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := i.DB.ExecContext(ctx, query, spn.UserID, spn.SpendingTypeID, spn.Amount, spn.Date.Format("2006-01-02"), spn.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	spn.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &spn, nil
}
