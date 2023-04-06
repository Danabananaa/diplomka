package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type debt struct {
	DB *sqlx.DB
}

func NewDebtRepo(db *sqlx.DB) *debt {
	return &debt{
		DB: db,
	}
}

func (lia *debt) AddLiabilities(ctx context.Context, ast model.Loan_Debt) (*model.Loan_Debt, error) {
	query := `INSERT INTO debt (user_id, type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := lia.DB.ExecContext(ctx, query, ast.UserID, ast.TypeID, ast.Amount, ast.Date.Format("2006-01-02"), ast.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}
	liab := model.Loan_Debt{}
	liab.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &liab, nil
}

func (lia *debt) GetLiabilities(ctx context.Context, bet model.Between) ([]*model.Loan_Debt, error) {
	liaarr := make([]*model.Loan_Debt, 0)
	query := `SELECT * FROM debt where user_id=? and date BETWEEN ? and ?`

	row, err := lia.DB.QueryxContext(ctx, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		lia := &model.Loan_Debt{}

		err := row.Scan(&lia.ID, &lia.UserID, &lia.TypeID, &lia.Amount, &lia.Description, &lia.Date)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		liaarr = append(liaarr, lia)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}
	return liaarr, nil
}
