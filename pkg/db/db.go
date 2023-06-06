package db

import (
	"github.com/jmoiron/sqlx"
	"os"
	"path"
)

type DB struct {
	db *sqlx.DB
}

type contextKey int

const txKey contextKey = iota

func New(db *sqlx.DB) *DB {
	return &DB{db}
}

func (db *DB) DB() *sqlx.DB {
	return db.db
}

func (db *DB) Migrate(migrationsDir string) error {
	if entries, err := os.ReadDir(migrationsDir); err != nil {
		return err
	} else {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			fullPath := path.Join(migrationsDir, entry.Name())
			if content, err := os.ReadFile(fullPath); err != nil {
				return err
			} else {
				db.db.MustExec(string(content))
			}
		}
		return nil
	}
}
