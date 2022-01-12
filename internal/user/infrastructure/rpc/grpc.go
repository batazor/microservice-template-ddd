//go:generate protoc -I. -I../../domain --go-grpc_out=Minternal/user/domain/user.proto=.:. --go_out=Minternal/user/domain/user.proto=.:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative user_rpc.proto

package user_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"microservice-template-ddd/internal/user/application"
	"microservice-template-ddd/pkg/rpc"
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
	service *user.Service
}

func New(runRPCServer *rpc.RPCServer, log *zap.Logger, userService *user.Service) (*UserServer, error) {
	server := &UserServer{
		log: log,

		service: userService,
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
