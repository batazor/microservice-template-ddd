package http

import (
	"context"

	"go.uber.org/zap"

	billing_rpc "robovoice-template/internal/billing/domain"
	"robovoice-template/internal/user/domain"
)

// API ...
type API struct { // nolint unused
	ctx context.Context
	Log *zap.Logger

	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
}
