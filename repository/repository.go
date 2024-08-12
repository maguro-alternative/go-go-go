package repository

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

type Repository struct {
	db *gorm.DB
}
