package user_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/di"
	"robovoice-template/internal/user/domain"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (UserRPCClient, error) {
	// Register clients
	client := NewUserRPCClient(rpcClient)

	return client, nil
}

type UserServer struct {
	log *zap.Logger

	UnimplementedUserRPCServer
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*UserServer, error) {
	server := &UserServer{
		log: log,
	}

	// Register services
	RegisterUserRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *UserServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	return &GetResponse{
		User: &domain.User{
			Login:    "test@user",
			Password: "",
			Email:    "test@user.com",
			IsActive: true,
		},
	}, nil
}
