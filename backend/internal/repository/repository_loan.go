package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type loan struct {
	DB *sqlx.DB
}

func NewLoanRepo(db *sqlx.DB) *loan {
	return &loan{
		DB: db,
	}
}

func (ass *loan) AddAssets(ctx context.Context, asl model.Loan_Debt) (*model.Loan_Debt, error) {
	query := `INSERT INTO loan (user_id, type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := ass.DB.ExecContext(ctx, query, asl.UserID, asl.TypeID, asl.Amount, asl.Date.Format("2006-01-02"), asl.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}
	ast := model.Loan_Debt{}
	ast.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &ast, nil
}

func (ass *loan) GetAssets(ctx context.Context, bet model.Between) ([]*model.Loan_Debt, error) {
	astarr := make([]*model.Loan_Debt, 0)
	query := `SELECT * FROM loan where user_id=? and date BETWEEN ? and ?`

	row, err := ass.DB.QueryxContext(ctx, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		ast := &model.Loan_Debt{}

		err := row.Scan(&ast.ID, &ast.UserID, &ast.TypeID, &ast.Amount, &ast.Description, &ast.Date)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		ast.Status = true
		astarr = append(astarr, ast)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}
	return astarr, nil
}
