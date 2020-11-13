//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

/*
Main DI-package
*/
package di

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/pkg/traicing"
)

// Service - heplers
type Service struct {
	Log       *zap.Logger
	ClientRPC *grpc.ClientConn
	ServerRPC *RPCServer
}

type RPCServer struct {
	Run      func()
	Server   *grpc.Server
	Endpoint string
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

// TODO: Move to inside package
// runGRPCServer ...
func runGRPCServer(log *zap.Logger, tracer opentracing.Tracer) (*RPCServer, func(), error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	endpoint := fmt.Sprintf("0.0.0.0:%d", grpc_port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, nil, err
	}

	// Initialize the gRPC server.
	rpc := grpc.NewServer(
		// Initialize your gRPC server's interceptor.
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.UnaryServerInterceptor,
		)),

		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.StreamServerInterceptor,
		)),
	)

	r := &RPCServer{
		Server: rpc,
		Run: func() {
			// After all your registrations, make sure all of the Prometheus metrics are initialized.
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
	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	// Set up a connection to the server peer
	conn, err := grpc.Dial(
		fmt.Sprintf("0.0.0.0:%d", grpc_port),
		grpc.WithInsecure(),

		// Initialize your gRPC server's interceptor.
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.UnaryClientInterceptor,
		)),

		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_prometheus.StreamClientInterceptor,
		)),
	)
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

func InitializeAPIService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(APISet))
}

// UserService =========================================================================================================
var UserSet = wire.NewSet(DefaultSet, runGRPCServer, NewUserService)

func NewUserService(log *zap.Logger, serverRPC *RPCServer) (*Service, error) {
	return &Service{
		Log:       log,
		ServerRPC: serverRPC,
	}, nil
}

func InitializeUserService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(UserSet))
}

// UserService =========================================================================================================
var BillingSet = wire.NewSet(DefaultSet, runGRPCServer, NewBillingService)

func NewBillingService(log *zap.Logger, serverRPC *RPCServer) (*Service, error) {
	return &Service{
		Log:       log,
		ServerRPC: serverRPC,
	}, nil
}

func InitializeBillingService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(BillingSet))
}
