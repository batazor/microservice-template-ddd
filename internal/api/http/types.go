package http

import (
	"context"

	"go.uber.org/zap"

	"robovoice-template/internal/billing/infrastructure/rpc"
	"robovoice-template/internal/book/infrastructure/rpc"
	"robovoice-template/internal/user/infrastructure/rpc"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
	Log *zap.Logger

	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
	BookService    book_rpc.BookRPCClient
}
