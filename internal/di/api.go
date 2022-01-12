//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"microservice-template-ddd/internal/api"
)

type APIService struct {
	Log *zap.Logger

	ClientRPC *grpc.ClientConn
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	// log, tracer
	DefaultSet,

	// gRPC client
	runGRPCClient,

	// CMD
	NewAPIService,
)

// InitConstructor =====================================================================================================
func NewAPIService(ctx context.Context, log *zap.Logger, clientRPC *grpc.ClientConn) (*APIService, error) {
	// Run API server
	var API api.Server
	API.RunAPIServer(ctx, log, clientRPC)

	return &APIService{
		Log:       log,
		ClientRPC: clientRPC,
	}, nil
}

func InitializeAPIService(ctx context.Context) (*APIService, func(), error) {
	panic(wire.Build(APISet))
}
