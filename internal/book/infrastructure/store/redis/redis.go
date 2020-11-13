package redis

import (
	"context"

	"github.com/go-redis/redis"

	book_rpc "robovoice-template/internal/book/domain"
	"robovoice-template/internal/db"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *redis.Client
}

// Init ...
func (_ *Store) Init(_ context.Context, _ *db.Store) error {
	return nil
}

// Get ...
func (r *Store) Get(ctx context.Context, id string) (*book_rpc.Book, error) {
	panic("implement me")
}

// List ...
func (r *Store) List(ctx context.Context, filter interface{}) ([]*book_rpc.Book, error) { // nolint unused
	panic("implement me")
}

// Add ...
func (r *Store) Add(ctx context.Context, data *book_rpc.Book) (*book_rpc.Book, error) {
	panic("implement me")
}

// Update ...
func (r *Store) Update(ctx context.Context, data *book_rpc.Book) (*book_rpc.Book, error) {
	panic("implement me")
}

// Delete ...
func (r *Store) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
