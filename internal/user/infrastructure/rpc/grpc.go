//go:generate protoc -I/usr/local/include -I. -I../../domain --go-grpc_out=Minternal/user/domain/user.proto=.:. --go_out=Minternal/user/domain/user.proto=.:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative user_rpc.proto

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

func New(runRPCServer *di.RPCServer, log *zap.Logger, userService *application.Service) (*UserServer, error) {
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
