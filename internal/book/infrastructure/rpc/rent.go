package book_rpc

import (
	"context"
)

func (m *BookServer) Rent(ctx context.Context, in *RentRequest) (*RentResponse, error) {
	book, err := m.service.Rent(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &RentResponse{
		Book: book,
	}, nil
}
