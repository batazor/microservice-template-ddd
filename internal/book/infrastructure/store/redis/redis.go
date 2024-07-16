package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"

	"microservice-template-ddd/internal/book/domain"
	"microservice-template-ddd/internal/db"
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
func (r *Store) Get(ctx context.Context, id string) (*domain.Book, error) {
	val, err := r.client.Get(id).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, fmt.Errorf("Not found id: %s", id)
		}

		return nil, err
	}

	var book domain.Book
	err = protojson.Unmarshal([]byte(val), &book)
	if err != nil {
		return nil, fmt.Errorf("Error parse book by id: %s", id)
	}

	return &book, nil
}

// List ...
func (r *Store) List(ctx context.Context, filter interface{}) ([]*domain.Book, error) { // nolint unused
	panic("implement me")
}

// Add ...
func (r *Store) Add(ctx context.Context, in *domain.Book) (*domain.Book, error) {
	m := protojson.MarshalOptions{}
	json, err := m.Marshal(in)
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
func (r *Store) Update(ctx context.Context, in *domain.Book) (*domain.Book, error) {
	m := protojson.MarshalOptions{}
	json, err := m.Marshal(in)
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
