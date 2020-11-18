//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

/*
Main DI-package
*/
package di

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	// common
	"robovoice-template/internal/db"
	"robovoice-template/pkg/rpc"
	"robovoice-template/pkg/traicing"

	// billing
	"robovoice-template/internal/billing/application"
	"robovoice-template/internal/billing/infrastructure/rpc"

	// user
	"robovoice-template/internal/user/application"
	"robovoice-template/internal/user/infrastructure/rpc"

	// book
	"robovoice-template/internal/book/application"
	"robovoice-template/internal/book/infrastructure/rpc"
	"robovoice-template/internal/book/infrastructure/store"
)

// TODO: Move to inside package
// runGRPCServer ...
func runGRPCServer(log *zap.Logger, tracer opentracing.Tracer) (*rpc.RPCServer, func(), error) {
	return rpc.InitServer(log, tracer)
}

// TODO: Move to inside package
// runGRPCClient - set up a connection to the server.
func runGRPCClient(log *zap.Logger, tracer opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	return rpc.InitClient(log, tracer)
}

// DefaultService ======================================================================================================
type DefaultService struct {
	Log *zap.Logger
}

var DefaultSet = wire.NewSet(InitLogger, InitTracer)

func InitLogger(ctx context.Context) (*zap.Logger, error) {
	viper.SetDefault("LOG_LEVEL", zap.InfoLevel)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	// TODO: add conf
	//conf := logger.Configuration{
	//	Level:      viper.GetInt("LOG_LEVEL"),
	//	TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	//}

	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return log, nil
}

func InitTracer(ctx context.Context, log *zap.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink") // Service Name
	viper.SetDefault("TRACER_URI", "localhost:6831")     // Tracing addr:host

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(config)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}

// InitStore return db
func InitStore(ctx context.Context, log *zap.Logger) (*db.Store, func(), error) {
	var st db.Store
	db, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

// APIService ==========================================================================================================
type APIService struct {
	Log *zap.Logger

	ClientRPC *grpc.ClientConn
}

var APISet = wire.NewSet(DefaultSet, runGRPCClient, NewAPIService)

func NewAPIService(log *zap.Logger, clientRPC *grpc.ClientConn) (*APIService, error) {
	return &APIService{
		Log:       log,
		ClientRPC: clientRPC,
	}, nil
}

func InitializeAPIService(ctx context.Context) (*APIService, func(), error) {
	panic(wire.Build(APISet))
}

// UserService =========================================================================================================
type UserService struct {
	Log *zap.Logger

	userRPCServer *user_rpc.UserServer
}

var UserSet = wire.NewSet(DefaultSet, runGRPCServer, NewUserService, NewUserApplication, NewUserRPCServer)

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

// BillingService ======================================================================================================
type BillingService struct {
	Log *zap.Logger

	billingRPCServer *billing_rpc.BillingServer
}

var BillingSet = wire.NewSet(DefaultSet, runGRPCServer, InitBillingService, NewBillingApplication, NewBillingRPCServer)

func NewBillingApplication() (*billing.Service, error) {
	billingService, err := billing.New()
	if err != nil {
		return nil, err
	}

	return billingService, nil
}

func NewBillingRPCClient(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) (billing_rpc.BillingRPCClient, error) {
	billingService, err := billing_rpc.Use(ctx, rpcClient)
	if err != nil {
		return nil, err
	}

	return billingService, nil
}

func NewBillingRPCServer(billingService *billing.Service, log *zap.Logger, serverRPC *rpc.RPCServer) (*billing_rpc.BillingServer, error) {
	billingRPCServer, err := billing_rpc.New(serverRPC, log, billingService)
	if err != nil {
		return nil, err
	}

	return billingRPCServer, nil
}

func InitBillingService(log *zap.Logger, billingRPCServer *billing_rpc.BillingServer) (*BillingService, error) {
	return &BillingService{
		Log: log,

		billingRPCServer: billingRPCServer,
	}, nil
}

func InitializeBillingService(ctx context.Context) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}

// BookService =========================================================================================================
type BookService struct {
	Log *zap.Logger

	bookRPCServer *book_rpc.BookServer
}

var BookSet = wire.NewSet(DefaultSet, runGRPCServer, runGRPCClient, NewBookApplication, NewBillingRPCClient, NewUserRPCClient, InitStore, InitBookStore, NewBookRPCServer, NewBookService)

// InitMetaStore
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
