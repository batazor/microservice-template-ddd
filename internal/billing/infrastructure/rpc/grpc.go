package billing_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	billing "robovoice-template/internal/billing/domain"
	"robovoice-template/internal/di"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (BillingRPCClient, error) {
	// Register clients
	client := NewBillingRPCClient(rpcClient)

	return client, nil
}

type BillingServer struct {
	log *zap.Logger

	UnimplementedBillingRPCServer
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*BillingServer, error) {
	server := &BillingServer{
		log: log,
	}

	// Register services
	RegisterBillingRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *BillingServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	return &GetResponse{
		Billing: &billing.Billing{
			Balance: 100.00,
		},
	}, nil
}
