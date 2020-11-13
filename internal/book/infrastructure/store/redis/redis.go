package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/jsonpb"

	book_rpc "robovoice-template/internal/book/domain"
	"robovoice-template/internal/db"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *redis.Client
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(*redis.Client)

	return nil
}

// Get ...
func (r *Store) Get(ctx context.Context, id string) (*book_rpc.Book, error) {
	val, err := r.client.Get(id).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, fmt.Errorf("Not found id: %s", id)
		}

		return nil, err
	}

	var book book_rpc.Book
	err = jsonpb.UnmarshalString(val, &book)
	if err != nil {
		return nil, fmt.Errorf("Error parse book by id: %s", id)
	}

	return &book, nil
}

// List ...
func (r *Store) List(ctx context.Context, filter interface{}) ([]*book_rpc.Book, error) { // nolint unused
	panic("implement me")
}

// Add ...
func (r *Store) Add(ctx context.Context, in *book_rpc.Book) (*book_rpc.Book, error) {
	m := jsonpb.Marshaler{}
	json, err := m.MarshalToString(in)
	if err != nil {
		return nil, fmt.Errorf("Error convert to JSON id: %s", in.Title)
	}

	err = r.client.Set(in.Title, json, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("Error update book by id: %s", in.Title)
	}

	return in, nil
}

// Update ...
func (r *Store) Update(ctx context.Context, in *book_rpc.Book) (*book_rpc.Book, error) {
	m := jsonpb.Marshaler{}
	json, err := m.MarshalToString(in)
	if err != nil {
		return nil, fmt.Errorf("Error convert to JSON id: %s", in.Title)
	}

	err = r.client.Set(in.Title, json, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("Error update book by id: %s", in.Title)
	}

	return in, nil
}

// Delete ...
func (r *Store) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
