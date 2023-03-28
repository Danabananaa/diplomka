package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(nameDB string, path string) error {
	file, err := os.Create(fmt.Sprintf("%s.sqlite", nameDB))
	if err != nil {
		return err
	}
	return file.Close()
}
