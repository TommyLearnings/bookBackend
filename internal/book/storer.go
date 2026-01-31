package book

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/TommyLearning/bookBackend/internal/response"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Store struct {
	db bun.IDB
}

func NewStore(db bun.IDB) *Store {
	return &Store{db: db}
}

// Create news record.
func (s Store) Create(ctx context.Context, bi100I *Record) (*Record, error) {
	bi100I.ReferenceId = uuid.New()
	if err := s.db.NewInsert().Model(bi100I).Returning("*").Scan(ctx, bi100I); err != nil {
		return nil, response.NewCustomError(err, http.StatusInternalServerError)
	}
	return bi100I, nil
}

func (s Store) FindById(ctx context.Context, id uuid.UUID) (bi100I *Record, err error) {
	err = s.db.NewSelect().Model(&bi100I).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return bi100I, response.NewCustomError(err, http.StatusNotFound)
		}

		return bi100I, response.NewCustomError(err, http.StatusInternalServerError)
	}
	return bi100I, nil
}

func (s Store) FindAll(ctx context.Context) (bi100I []*Record, err error) {
	err = s.db.NewSelect().Model(&bi100I).Scan(ctx, &bi100I)
	return bi100I, err
}

func (s Store) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.NewDelete().Model(&Record{}).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return response.NewCustomError(err, http.StatusInternalServerError)
	}
	return nil
}

// UpdateById update bi100I by it's ID.
func (s Store) UpdateById(ctx context.Context, id uuid.UUID, bi100I *Record) (err error) {
	r, err := s.db.NewUpdate().Model(bi100I).Where("id = ?", id).Returning("NULL").Exec(ctx)
	if err != nil {
		return response.NewCustomError(err, http.StatusInternalServerError)
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return response.NewCustomError(err, http.StatusInternalServerError)
	}
	if rowsAffected == 0 {
		return response.NewCustomError(err, http.StatusNotFound)
	}
	return nil
}
