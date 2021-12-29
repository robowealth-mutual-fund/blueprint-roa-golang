package postgres

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/database"
)

// this file should be a query db process if use must be install gorm lib

// Repository ...
type Repository struct {
	Db *database.DB
}

// NewRepository ...
func NewRepository(db *database.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
