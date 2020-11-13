package user_rpc

import (
	"context"

	"google.golang.org/grpc"

	"robovoice-template/internal/user/domain"
)

type rpc struct {
	client     *grpc.ClientConn
	UserClient user_rpc.UserServiceClient
}

func Use(_ context.Context, rpcClient *grpc.ClientConn) (*rpc, error) {
	r := &rpc{
		client: rpcClient,

		// Register clients
		UserClient: user_rpc.NewUserServiceClient(rpcClient),
	}

	return r, nil
}
