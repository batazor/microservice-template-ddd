package rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	billing_rpc "robovoice-template/internal/billing/domain"
	"robovoice-template/internal/di"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (billing_rpc.BillingRPCClient, error) {
	// Register clients
	client := billing_rpc.NewBillingRPCClient(rpcClient)

	return client, nil
}

type BillingServer struct {
	log *zap.Logger

	billing_rpc.UnimplementedBillingRPCServer
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*BillingServer, error) {
	server := &BillingServer{
		log: log,
	}

	// Register services
	billing_rpc.RegisterBillingRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *BillingServer) Get(ctx context.Context, in *billing_rpc.GetRequest) (*billing_rpc.GetResponse, error) {
	return &billing_rpc.GetResponse{
		Billing: &billing_rpc.Billing{
			Balance: 100.00,
		},
	}, nil
}
