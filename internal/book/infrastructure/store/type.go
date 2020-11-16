package store

import (
	"context"

	"robovoice-template/internal/book/domain"
	"robovoice-template/internal/db"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error

	// CRUD
	Get(ctx context.Context, id string) (*domain.Book, error)
	List(ctx context.Context, filter interface{}) ([]*domain.Book, error)
	Add(ctx context.Context, data *domain.Book) (*domain.Book, error)
	Update(ctx context.Context, data *domain.Book) (*domain.Book, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type BookStore struct { // nolint unused
	typeStore string
	Store     Repository
}
