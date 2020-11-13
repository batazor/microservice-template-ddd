// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"robovoice-template/internal/book/infrastructure/store"
	"robovoice-template/internal/db"
	"robovoice-template/pkg/traicing"
	"time"
)

// Injectors from wire.go:

func InitializeAPIService(ctx context.Context) (*Service, func(), error) {
	logger, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	tracer, cleanup, err := InitTracer(ctx, logger)
	if err != nil {
		return nil, nil, err
	}
	clientConn, cleanup2, err := runGRPCClient(logger, tracer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewAPIService(logger, clientConn)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup2()
		cleanup()
	}, nil
}

func InitializeUserService(ctx context.Context) (*Service, func(), error) {
	logger, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	tracer, cleanup, err := InitTracer(ctx, logger)
	if err != nil {
		return nil, nil, err
	}
	rpcServer, cleanup2, err := runGRPCServer(logger, tracer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewUserService(logger, rpcServer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup2()
		cleanup()
	}, nil
}

func InitializeBillingService(ctx context.Context) (*Service, func(), error) {
	logger, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	tracer, cleanup, err := InitTracer(ctx, logger)
	if err != nil {
		return nil, nil, err
	}
	rpcServer, cleanup2, err := runGRPCServer(logger, tracer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewBillingService(logger, rpcServer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup2()
		cleanup()
	}, nil
}

func InitializeBookService(ctx context.Context) (*Service, func(), error) {
	logger, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	store, cleanup, err := InitStore(ctx, logger)
	if err != nil {
		return nil, nil, err
	}
	bookStore, err := InitBookStore(ctx, logger, store)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup2, err := InitTracer(ctx, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup3, err := runGRPCServer(logger, tracer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup4, err := runGRPCClient(logger, tracer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewBookService(logger, bookStore, rpcServer, clientConn)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	Log *zap.Logger
	DB  *db.Store

	BookStore *store.BookStore

	ClientRPC *grpc.ClientConn
	ServerRPC *RPCServer
}

type RPCServer struct {
	Run      func()
	Server   *grpc.Server
	Endpoint string
}

// InitStore return db
func InitStore(ctx context.Context, log *zap.Logger) (*db.Store, func(), error) {
	var st db.Store
	db2, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db2.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db2, cleanup, nil
}

// InitMetaStore
func InitBookStore(ctx context.Context, log *zap.Logger, conn *db.Store) (*store.BookStore, error) {
	st := store.BookStore{}
	bookStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return bookStore, nil
}

func InitTracer(ctx context.Context, log *zap.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink")
	viper.SetDefault("TRACER_URI", "localhost:6831")

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

// TODO: Move to inside package
// runGRPCServer ...
func runGRPCServer(log *zap.Logger, tracer opentracing.Tracer) (*RPCServer, func(), error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051")
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	endpoint := fmt.Sprintf("0.0.0.0:%d", grpc_port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, nil, err
	}

	rpc := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()), grpc_prometheus.UnaryServerInterceptor)), grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()), grpc_prometheus.StreamServerInterceptor)))

	r := &RPCServer{
		Server: rpc,
		Run: func() {
			grpc_prometheus.Register(rpc)

			go rpc.Serve(lis)
			log.Info("Run gRPC server", zap.Int("port", grpc_port))
		},
		Endpoint: endpoint,
	}

	cleanup := func() {
		rpc.GracefulStop()
	}

	return r, cleanup, err
}

// TODO: Move to inside package
// runGRPCClient - set up a connection to the server.
func runGRPCClient(log *zap.Logger, tracer opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	viper.SetDefault("GRPC_CLIENT_PORT", "50051")
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", grpc_port), grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()), grpc_prometheus.UnaryClientInterceptor)), grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads()), grpc_prometheus.StreamClientInterceptor)))
	if err != nil {
		return nil, nil, err
	}

	log.Info("Run gRPC client", zap.Int("port", grpc_port))

	cleanup := func() {
		conn.Close()
	}

	return conn, cleanup, nil
}

func InitLogger(ctx context.Context) (*zap.Logger, error) {
	viper.SetDefault("LOG_LEVEL", zap.InfoLevel)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return log, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(InitLogger, InitTracer)

// APIService ==========================================================================================================
var APISet = wire.NewSet(DefaultSet, runGRPCClient, NewAPIService)

func NewAPIService(log *zap.Logger, clientRPC *grpc.ClientConn) (*Service, error) {
	return &Service{
		Log:       log,
		ClientRPC: clientRPC,
	}, nil
}

// UserService =========================================================================================================
var UserSet = wire.NewSet(DefaultSet, runGRPCServer, NewUserService)

func NewUserService(log *zap.Logger, serverRPC *RPCServer) (*Service, error) {
	return &Service{
		Log:       log,
		ServerRPC: serverRPC,
	}, nil
}

// UserService =========================================================================================================
var BillingSet = wire.NewSet(DefaultSet, runGRPCServer, NewBillingService)

func NewBillingService(log *zap.Logger, serverRPC *RPCServer) (*Service, error) {
	return &Service{
		Log:       log,
		ServerRPC: serverRPC,
	}, nil
}

// BookService =========================================================================================================
var BookSet = wire.NewSet(DefaultSet, runGRPCServer, runGRPCClient, InitStore, InitBookStore, NewBookService)

func NewBookService(log *zap.Logger, bookStore *store.BookStore, serverRPC *RPCServer, clientRPC *grpc.ClientConn) (*Service, error) {
	return &Service{
		Log: log,

		BookStore: bookStore,

		ServerRPC: serverRPC,
		ClientRPC: clientRPC,
	}, nil
}
