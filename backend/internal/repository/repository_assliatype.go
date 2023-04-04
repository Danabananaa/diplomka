package repository

import (
	"context"
	"database/sql"
	"diplomka/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type assliatype struct {
	DB *sqlx.DB
}

func NewAssLiaTypeRepo(db *sqlx.DB) *assliatype {
	return &assliatype{
		DB: db,
	}
}

func (alt *assliatype) GetAssLiaTypeRepo(ctx context.Context) ([]*model.AssLiaType, error) {
	assliaarr := make([]*model.AssLiaType, 0)
	query := `SELECT * FROM incometype`
	row, err := alt.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error from QueryContext")
	}
	defer row.Close()

	for row.Next() {
		asslia := &model.AssLiaType{}
		err := row.Scan(&asslia.ID, &asslia.Type)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("No rows into table")
			} else {
				return nil, err
			}
		}
		assliaarr = append(assliaarr, asslia)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("Error from row")
	}

	return assliaarr, nil
}
