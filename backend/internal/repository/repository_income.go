package repository

import (
	"context"
	"diplomka/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type income struct {
	DB *sqlx.DB
}

func NewIncomeRepo(db *sqlx.DB) *income {
	return &income{
		DB: db,
	}
}

func (i *income) AddIncome(ctx context.Context, inc model.Income) (*model.Income, error) {
	query := `INSERT INTO income (user_id, income_type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := i.DB.ExecContext(ctx, query, inc.UserID, inc.IncomeTypeID, inc.Amount, inc.Date.Format("2006-01-02"), inc.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	inc.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &inc, nil
}
