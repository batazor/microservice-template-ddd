package rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	billing_rpc "robovoice-template/internal/billing/domain"
	"robovoice-template/internal/book/application"
	"robovoice-template/internal/book/domain"
	"robovoice-template/internal/book/infrastructure/store"
	"robovoice-template/internal/di"
	user_rpc "robovoice-template/internal/user/domain"
)

func Use(_ context.Context, rpcClient *grpc.ClientConn) (book_rpc.BookRPCClient, error) {
	// Register clients
	client := book_rpc.NewBookRPCClient(rpcClient)

	return client, nil
}

type BookServer struct {
	log *zap.Logger

	book_rpc.UnimplementedBookRPCServer

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
		},

		UserService:    userService,
		BillingService: billingService,
	}

	// Register services
	book_rpc.RegisterBookRPCServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *BookServer) Get(ctx context.Context, in *book_rpc.GetRequest) (*book_rpc.GetResponse, error) {
	return &book_rpc.GetResponse{
		Book: &book_rpc.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		},
	}, nil
}
