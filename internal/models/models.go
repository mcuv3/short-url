package models

import (
	"database/sql"
)

type Store struct {
	LinkService
}

// NewModels returns a model type with database connection pool
func NewStore(db *sql.DB) Store {
	return Store{
		LinkService: NewLinkRepo(db),
	}
}
