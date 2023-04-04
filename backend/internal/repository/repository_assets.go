package repository

import (
	"context"
	"diplomka/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type assets struct {
	DB *sqlx.DB
}

func NewAssetsRepo(db *sqlx.DB) *assets {
	return &assets{
		DB: db,
	}
}

func (ass *assets) AddAssets(ctx context.Context, ast model.Assets) (*model.Assets, error) {
	query := `INSERT INTO assests (user_id, type, amount, date, description) VALUES (?,?,?,?,?);`

	res, err := ass.DB.ExecContext(ctx, query, ast.UserID, ast.TypeID, ast.Amount, ast.Date.Format("2006-01-02"), ast.Description)
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}

	ast.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error was ocured during executing method LastInsertId: %v", err)
	}

	return &ast, nil
}
