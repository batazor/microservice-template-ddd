/*
Book Service. Application layer
*/
package application

import (
	"context"

	book_rpc "robovoice-template/internal/book/domain"
	"robovoice-template/internal/book/infrastructure/store"
)

type Service struct {
	Store *store.BookStore
}

func (s *Service) Rent(ctx context.Context, in *book_rpc.RentRequest) (*book_rpc.GetResponse, error) {
	return &book_rpc.GetResponse{
		Book: &book_rpc.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		},
	}, nil
}
