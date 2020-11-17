/*
Book Service. Application layer
*/
package application

import (
	"context"
	"fmt"

	"robovoice-template/internal/billing/infrastructure/rpc"
	"robovoice-template/internal/book/domain"
	"robovoice-template/internal/book/infrastructure/store"
	"robovoice-template/internal/user/infrastructure/rpc"
)

type Service struct {
	Store *store.BookStore

	// ServiceClients
	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
}

func New(store *store.BookStore, userService user_rpc.UserRPCClient, billingService billing_rpc.BillingRPCClient) (*Service, error) {
	return &Service{
		Store: store,

		UserService:    userService,
		BillingService: billingService,
	}, nil
}

// Get - get book from store
func (s *Service) Get(ctx context.Context, bookId string) (*domain.Book, error) {
	// Get book from store
	book, err := s.Store.Store.Get(ctx, "Hello World")
	if err != nil {
		// For example create book
		s.Store.Store.Add(ctx, &domain.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		})

		return nil, err
	}

	return book, nil
}

func (s *Service) Rent(ctx context.Context, bookId string) (*domain.Book, error) {
	// Get user
	user, err := s.UserService.Get(ctx, &user_rpc.GetRequest{Id: bookId})
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
	book, err := s.Get(ctx, "Hello World")
	if err != nil {
		return nil, err
	}

	// Change state in DB
	book.IsRent = !book.IsRent
	book, err = s.Store.Store.Update(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
