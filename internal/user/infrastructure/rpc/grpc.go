package user_rpc

import (
	"context"

	"google.golang.org/grpc"

	"robovoice-template/internal/user/domain"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (user_rpc.UserServiceClient, error) {
	// Register clients
	client := user_rpc.NewUserServiceClient(rpcClient)

	return client, nil
}
