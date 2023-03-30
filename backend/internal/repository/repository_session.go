package repository

import (
	"context"
	"diplomka/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type session struct {
	DB *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) *session {
	return &session{DB: db}
}

func (s *session) AddSession(ctx context.Context, session model.Session) error {
	query := `insert into sessions (token,userid) VALUES(?,?)`
	_, err := s.DB.ExecContext(ctx, query, session.Token, session.UserID)
	if err != nil {
		return fmt.Errorf("error was ocured during executing method ExecContext: %v", err)
	}
	return nil
}
