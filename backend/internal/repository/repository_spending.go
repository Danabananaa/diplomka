package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"diplomka/internal/model"
)

func (s *repo) AddSpending(ctx context.Context, spn model.Spending) (*model.Spending, error) {
	query := `INSERT INTO spending (user_id, spending_type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := s.DB.ExecContext(ctx, query, spn.UserID, spn.SpendingTypeID, spn.Amount, spn.Date.Format("2006-01-02"), spn.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	spn.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &spn, nil
}

func (s *repo) GetSpending(ctx context.Context, bet model.Between) ([]*model.Spending, error) {
	incarr := make([]*model.Spending, 0)

	query := `SELECT * FROM spending where user_id=? and date BETWEEN ? and ?`

	// err := s.DB.SelectContext(ctx, &incarr, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	row, err := s.DB.QueryxContext(ctx, query, bet.UserID, bet.StartDate.Format("2006-01-02"), bet.EndDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		spn := &model.Spending{}

		err := row.Scan(&spn.ID, &spn.UserID, &spn.SpendingTypeID, &spn.Amount, &spn.Description, &spn.Date)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		incarr = append(incarr, spn)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}
	return incarr, nil
}
