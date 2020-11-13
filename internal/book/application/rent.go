/*
Book Service. Application layer
*/
package application

import (
	"context"
	"fmt"

	billing_rpc "robovoice-template/internal/billing/domain"
	book_rpc "robovoice-template/internal/book/domain"
	"robovoice-template/internal/book/infrastructure/store"
	user_rpc "robovoice-template/internal/user/domain"
)

type Service struct {
	Store *store.BookStore

	// ServiceClients
	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
}

func (s *Service) Rent(ctx context.Context, in *book_rpc.RentRequest) (*book_rpc.RentResponse, error) {
	// Get user
	user, err := s.UserService.Get(ctx, &user_rpc.GetRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	// Check billing
	billing, err := s.BillingService.Get(ctx, &billing_rpc.GetRequest{Id: user.User.Email})
	if err != nil {
		return nil, err
	}

	if billing.Billing.Balance <= 0 {
		return nil, fmt.Errorf("Problem with balance. Current balance %f", billing.Billing.Balance)
	}

	// Get book from store
	book, err := s.Store.Store.Get(ctx, "Hello World")
	if err != nil {
		// For example create book
		s.Store.Store.Add(ctx, &book_rpc.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		})

		return nil, err
	}

	// Change state in DB
	book.IsRent = !book.IsRent
	book, err = s.Store.Store.Update(ctx, book)
	if err != nil {
		return nil, err
	}

	return &book_rpc.RentResponse{
		Book: book,
	}, nil
}
