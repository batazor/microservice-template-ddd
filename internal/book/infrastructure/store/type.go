package store

import (
	"context"

	book_rpc "robovoice-template/internal/book/domain"
	"robovoice-template/internal/db"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error

	// CRUD
	Get(ctx context.Context, id string) (*book_rpc.Book, error)
	List(ctx context.Context, filter interface{}) ([]*book_rpc.Book, error)
	Add(ctx context.Context, data *book_rpc.Book) (*book_rpc.Book, error)
	Update(ctx context.Context, data *book_rpc.Book) (*book_rpc.Book, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type BookStore struct { // nolint unused
	typeStore string
	Store     Repository
}
