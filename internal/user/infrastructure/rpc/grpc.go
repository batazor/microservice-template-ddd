package user_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/di"
	"robovoice-template/internal/user/application"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (UserRPCClient, error) {
	// Register clients
	client := NewUserRPCClient(rpcClient)

	return client, nil
}

type UserServer struct {
	log *zap.Logger

	UnimplementedUserRPCServer

	// Application
	service *application.Service
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*UserServer, error) {
	server := &UserServer{
		log: log,

		service: &application.Service{},
	}

	// Register services
	RegisterUserRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *UserServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	user, err := m.service.Get()
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		User: user,
	}, nil
}
