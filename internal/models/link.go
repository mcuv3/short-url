package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        uuid.UUID `json:"id"`
	FullURL   string    `json:"full_url"`
	Title     string    `json:"title"`
	Short     string    `json:"short"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type LinkService interface {
	CreateLink(params CreateLinkParams) (*Link, error)
	GetLinkByShort(short string) (*Link, error)
}

type LinkRepository struct {
	db *sql.DB
}

type CreateLinkParams struct {
	FullURL string `json:"full_url"`
	Title   string `json:"title"`
	Short   string `json:"short"`
}

func NewLinkRepo(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) CreateLink(params CreateLinkParams) (*Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stmt := `INSERT INTO link (id,full_url,title,short,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6);`

	created := time.Now()
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, stmt, id, params.FullURL, params.Title, params.Short, created, created)
	if err != nil {
		return nil, err
	}

	return &Link{
		ID:        id,
		FullURL:   params.FullURL,
		Title:     params.Title,
		Short:     params.Short,
		CreatedAt: created,
		UpdateAt:  created,
	}, nil
}

func (r *LinkRepository) GetLinkByShort(short string) (*Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stmt := `SELECT id,full_url,title,short,created_at,updated_at FROM link WHERE short=$1;`

	row := r.db.QueryRowContext(ctx, stmt, short)

	var l Link
	err := row.Scan(
		&l.ID,
		&l.FullURL,
		&l.Title,
		&l.Short,
		&l.CreatedAt,
		&l.UpdateAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &l, nil
}
