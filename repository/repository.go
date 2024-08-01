package repository

import (
	"database/sql"
)

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

type Repository struct {
	db *sql.DB
}
