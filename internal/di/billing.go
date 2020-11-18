//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	billing "robovoice-template/internal/billing/application"
	billing_rpc "robovoice-template/internal/billing/infrastructure/rpc"
	"robovoice-template/pkg/rpc"
)

type BillingService struct {
	Log *zap.Logger

	billingRPCServer *billing_rpc.BillingServer
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(
	// log, tracer
	DefaultSet,

	// gRPC server
	runGRPCServer,
	NewBillingRPCServer,

	// applications
	NewBillingApplication,

	// CMD
	NewBillingService,
)

// InitConstructor =====================================================================================================
func NewBillingApplication() (*billing.Service, error) {
	billingService, err := billing.New()
	if err != nil {
		return nil, err
	}

	return billingService, nil
}

func NewBillingRPCClient(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) (billing_rpc.BillingRPCClient, error) {
	billingService, err := billing_rpc.Use(ctx, rpcClient)
	if err != nil {
		return nil, err
	}

	return billingService, nil
}

func NewBillingRPCServer(billingService *billing.Service, log *zap.Logger, serverRPC *rpc.RPCServer) (*billing_rpc.BillingServer, error) {
	billingRPCServer, err := billing_rpc.New(serverRPC, log, billingService)
	if err != nil {
		return nil, err
	}

	return billingRPCServer, nil
}

func NewBillingService(log *zap.Logger, billingRPCServer *billing_rpc.BillingServer) (*BillingService, error) {
	return &BillingService{
		Log: log,

		billingRPCServer: billingRPCServer,
	}, nil
}

func InitializeBillingService(ctx context.Context) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}
