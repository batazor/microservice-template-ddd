package rpc

import (
	"context"
	"fmt"

	billing_rpc "robovoice-template/internal/billing/domain"
	book_rpc "robovoice-template/internal/book/domain"
	user_rpc "robovoice-template/internal/user/domain"
)

func (m *BookServer) Rent(ctx context.Context, in *book_rpc.RentRequest) (*book_rpc.RentResponse, error) {
	// Get user
	user, err := m.UserService.Get(ctx, &user_rpc.GetRequest{Id: "user@test"})
	if err != nil {
		return nil, err
	}

	// Check billing
	billing, err := m.BillingService.Get(ctx, &billing_rpc.GetRequest{Id: user.User.Email})
	if err != nil {
		return nil, err
	}

	if billing.Billing.Balance <= 0 {
		return nil, fmt.Errorf("Problem with balance. Current balance %f", billing.Billing.Balance)
	}

	return &book_rpc.RentResponse{
		Book: &book_rpc.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: true,
		},
	}, nil
}
