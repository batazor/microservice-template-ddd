//go:generate protoc -I/usr/local/include -I. -I../../domain --go-grpc_out=Minternal/book/domain/book.proto=.:. --go_out=Minternal/book/domain/book.proto=.:. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative book_rpc.proto

package book_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/billing/infrastructure/rpc"
	"robovoice-template/internal/book/application"
	"robovoice-template/internal/book/infrastructure/store"
	"robovoice-template/internal/di"
	"robovoice-template/internal/user/infrastructure/rpc"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (BookRPCClient, error) {
	// Register clients
	client := NewBookRPCClient(rpcClient)

	return client, nil
}

type BookServer struct {
	log *zap.Logger

	UnimplementedBookRPCServer

	// Application
	service *application.Service

	// ServiceClients
	UserService    user_rpc.UserRPCClient
	BillingService billing_rpc.BillingRPCClient
}

func New(runRPCServer *di.RPCServer, log *zap.Logger, bookStore *store.BookStore, userService user_rpc.UserRPCClient, billingService billing_rpc.BillingRPCClient) (*BookServer, error) {
	server := &BookServer{
		log: log,

		service: &application.Service{
			Store: bookStore,

			UserService:    userService,
			BillingService: billingService,
		},

		UserService:    userService,
		BillingService: billingService,
	}

	// Register services
	RegisterBookRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}
