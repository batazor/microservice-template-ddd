package rpc

import (
	"context"

	book_rpc "robovoice-template/internal/book/domain"
)

func (m *BookServer) Rent(ctx context.Context, in *book_rpc.RentRequest) (*book_rpc.RentResponse, error) {
	return m.service.Rent(ctx, in)
}
