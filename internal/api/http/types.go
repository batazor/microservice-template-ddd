package http

import (
	"context"

	"go.uber.org/zap"

	"microservice-template-ddd/internal/billing/infrastructure/rpc"
	"microservice-template-ddd/internal/book/infrastructure/rpc"
	"microservice-template-ddd/internal/user/infrastructure/rpc"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
	Log *zap.Logger

	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
	BookService    book_rpc.BookRPCClient
}
