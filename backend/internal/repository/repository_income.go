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

// func (i income) GetIncome(ctx context.Context, inc model.Income, begindate string) ([]*model.Income, error) {
// 	incarr := []*model.Income{}
// 	query := `SELECT * FROM income where user_id=? and date BETWEEN ? and ?`
// 	row, err := i.DB.QueryContext(ctx, query, inc.UserID, begindate, inc.Date.Format("2006-01-02"))
// 	if err != nil {
// 		return nil, fmt.Errorf("Error from QueryContext")
// 	}
// 	defer row.Close()
// 	fmt.Println(row)

// 	for row.Next() {
// 		st := &model.Income{}

// 		err := row.Scan(&st.ID, &st.UserID, &st.IncomeTypeID, &st.Amount, &st.Description, model.CustomTime{&st.Date})
// 		if err != nil {
// 			if errors.Is(err, sql.ErrNoRows) {
// 				return nil, fmt.Errorf("No rows into table")
// 			} else {
// 				return nil, err
// 			}
// 		}
// 		incarr = append(incarr, st)
// 	}

// 	if err = row.Err(); err != nil {
// 		return nil, fmt.Errorf("Error from row")
// 	}
// 	return incarr, nil
// }
