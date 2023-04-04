package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
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

func (i *income) GetIncome(ctx context.Context, bet model.Between) ([]*model.Income, error) {
	incarr := make([]*model.Income, 0)

	query := `SELECT * FROM income where user_id=? and date BETWEEN ? and ?`

	// err := i.DB.SelectContext(ctx, &incarr, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	row, err := i.DB.QueryxContext(ctx, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		inc := &model.Income{}

		err := row.Scan(&inc.ID, &inc.UserID, &inc.IncomeTypeID, &inc.Amount, &inc.Description, &inc.Date)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		incarr = append(incarr, inc)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}
	return incarr, nil
}
