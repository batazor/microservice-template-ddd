//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	billing_rpc "robovoice-template/internal/billing/infrastructure/rpc"
	book "robovoice-template/internal/book/application"
	book_rpc "robovoice-template/internal/book/infrastructure/rpc"
	"robovoice-template/internal/book/infrastructure/store"
	"robovoice-template/internal/db"
	user_rpc "robovoice-template/internal/user/infrastructure/rpc"
	"robovoice-template/pkg/rpc"
)

type BookService struct {
	Log *zap.Logger

	bookRPCServer *book_rpc.BookServer
}

// BookService =========================================================================================================
var BookSet = wire.NewSet(
	// log, tracer
	DefaultSet,

	// gRPC server
	runGRPCServer,
	NewBookRPCServer,

	// gRPC client
	runGRPCClient,
	NewBillingRPCClient,
	NewUserRPCClient,

	// store
	InitStore,
	InitBookStore,

	// applications
	NewBookApplication,

	// CMD
	NewBookService,
)

// InitConstructor =====================================================================================================
func InitBookStore(ctx context.Context, log *zap.Logger, conn *db.Store) (*store.BookStore, error) {
	st := store.BookStore{}
	bookStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return bookStore, nil
}

func NewBookApplication(store *store.BookStore, billingRPC billing_rpc.BillingRPCClient, userRPC user_rpc.UserRPCClient) (*book.Service, error) {
	bookService, err := book.New(store, userRPC, billingRPC)
	if err != nil {
		return nil, err
	}

	return bookService, nil
}

func NewBookRPCClient(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) (book_rpc.BookRPCClient, error) {
	bookService, err := book_rpc.Use(ctx, rpcClient)
	if err != nil {
		return nil, err
	}

	return bookService, nil
}

func NewBookRPCServer(bookService *book.Service, log *zap.Logger, serverRPC *rpc.RPCServer) (*book_rpc.BookServer, error) {
	bookRPCServer, err := book_rpc.New(serverRPC, log, bookService)
	if err != nil {
		return nil, err
	}

	return bookRPCServer, nil
}

func NewBookService(log *zap.Logger, bookRPCServer *book_rpc.BookServer) (*BookService, error) {
	return &BookService{
		Log: log,

		bookRPCServer: bookRPCServer,
	}, nil
}

func InitializeBookService(ctx context.Context) (*BookService, func(), error) {
	panic(wire.Build(BookSet))
}
