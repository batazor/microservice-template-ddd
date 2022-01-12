//go:generate protoc -I. -I../../domain --go-grpc_out=Minternal/billing/domain/billing.proto=.:. --go_out=Minternal/billing/domain/billing.proto=.:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative billing_rpc.proto

package billing_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"microservice-template-ddd/internal/billing/application"
	"microservice-template-ddd/pkg/rpc"
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
	service *billing.Service
}

func New(runRPCServer *rpc.RPCServer, log *zap.Logger, billingService *billing.Service) (*BillingServer, error) {
	server := &BillingServer{
		log: log,

		service: billingService,
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
		Billing: &Payload{
			Billing: billing,
		},
	}, nil
}
