package book_rpc

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/billing/infrastructure/rpc"
	"robovoice-template/internal/book/application"
	"robovoice-template/internal/book/domain"
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

func (m *BookServer) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	return &GetResponse{
		Book: &domain.Book{
			Title:  "Hello World",
			Author: "God",
			IsRent: false,
		},
	}, nil
}
