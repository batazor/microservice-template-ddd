package rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/di"
	"robovoice-template/internal/user/domain"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (user_rpc.UserRPCClient, error) {
	// Register clients
	client := user_rpc.NewUserRPCClient(rpcClient)

	return client, nil
}

type UserServer struct {
	log *zap.Logger

	user_rpc.UnimplementedUserRPCServer
}

func New(runRPCServer *di.RPCServer, log *zap.Logger) (*UserServer, error) {
	server := &UserServer{
		log: log,
	}

	// Register services
	user_rpc.RegisterUserRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *UserServer) Get(ctx context.Context, in *user_rpc.GetRequest) (*user_rpc.GetResponse, error) {
	return &user_rpc.GetResponse{
		User: &user_rpc.User{
			Login:    "test@user",
			Password: "",
			Email:    "test@user.com",
			IsActive: true,
		},
	}, nil
}
