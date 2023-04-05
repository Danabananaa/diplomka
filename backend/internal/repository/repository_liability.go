package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type liabilities struct {
	DB *sqlx.DB
}

func NewLiabilitiesRepo(db *sqlx.DB) *liabilities {
	return &liabilities{
		DB: db,
	}
}

func (lia *liabilities) AddLiabilities(ctx context.Context, ast model.Assets_or_Liabs) (*model.Assets_and_Liabs, error) {
	query := `INSERT INTO liabilities (user_id, type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := lia.DB.ExecContext(ctx, query, ast.UserID, ast.TypeID, ast.Amount, ast.Date.Format("2006-01-02"), ast.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}
	liab := model.Assets_and_Liabs{}
	liab.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &liab, nil
}

func (lia *liabilities) GetLiabilities(ctx context.Context, bet model.Between) ([]*model.Liabilities, error) {
	liaarr := make([]*model.Liabilities, 0)
	query := `SELECT * FROM liabilities where user_id=? and date BETWEEN ? and ?`

	row, err := lia.DB.QueryxContext(ctx, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		lia := &model.Liabilities{}

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
