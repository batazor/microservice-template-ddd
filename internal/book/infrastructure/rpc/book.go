package book_rpc

import (
	"context"
)

func (m *BookServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	book, err := m.service.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		Book: book,
	}, nil
}
