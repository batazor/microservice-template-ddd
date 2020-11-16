package billing_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/billing/application"
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

	// Application
	service *application.Service
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*BillingServer, error) {
	server := &BillingServer{
		log: log,

		service: &application.Service{},
	}

	// Register services
	RegisterBillingRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *BillingServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	billing, err := m.service.Get()
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		Billing: billing,
	}, nil
}
