//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	user "microservice-template-ddd/internal/user/application"
	user_rpc "microservice-template-ddd/internal/user/infrastructure/rpc"
	"microservice-template-ddd/pkg/rpc"
)

type UserService struct {
	Log *zap.Logger

	userRPCServer *user_rpc.UserServer
}

// UserService =========================================================================================================
var UserSet = wire.NewSet(
	// log, tracer
	DefaultSet,

	// gRPC server
	runGRPCServer,
	NewUserRPCServer,

	// applications
	NewUserApplication,

	// CMD
	NewUserService,
)

// InitConstructor =====================================================================================================
func NewUserApplication() (*user.Service, error) {
	userService, err := user.New()
	if err != nil {
		return nil, err
	}

	return userService, nil
}

func NewUserRPCClient(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) (user_rpc.UserRPCClient, error) {
	userService, err := user_rpc.Use(ctx, rpcClient)
	if err != nil {
		return nil, err
	}

	return userService, nil
}

func NewUserRPCServer(userService *user.Service, log *zap.Logger, serverRPC *rpc.RPCServer) (*user_rpc.UserServer, error) {
	userRPCServer, err := user_rpc.New(serverRPC, log, userService)
	if err != nil {
		return nil, err
	}

	return userRPCServer, nil
}

func NewUserService(log *zap.Logger, userRPCServer *user_rpc.UserServer) (*UserService, error) {
	return &UserService{
		Log: log,

		userRPCServer: userRPCServer,
	}, nil
}

func InitializeUserService(ctx context.Context) (*UserService, func(), error) {
	panic(wire.Build(UserSet))
}
