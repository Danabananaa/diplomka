package repository

// import (
// 	"context"
// 	"diplomka/internal/model"
// 	"fmt"

// 	"github.com/jmoiron/sqlx"
// )

// type liabilities struct {
// 	DB *sqlx.DB
// }

// func NewLiabilitiesRepo(db *sqlx.DB) *liabilities {
// 	return &liabilities{
// 		DB: db,
// 	}
// }

// func (lia *liabilities) AddLiabilities(ctx context.Context, ast model.Liabilities) (*model.Liabilities, error) {
// 	query := `INSERT INTO liabilities (user_id, type, amount, date, description) VALUES (?,?,?,?,?);`

// 	res, err := lia.DB.ExecContext(ctx, query, ast.UserID, ast.TypeID, ast.Amount, ast.Date.Format("2006-01-02"), ast.Description)
// 	if err != nil {
// 		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
// 	}

// 	ast.ID, err = res.LastInsertId()
// 	if err != nil {
// 		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
// 	}

// 	return &ast, nil
// }
//
