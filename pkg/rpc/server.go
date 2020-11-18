package rpc

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitServer(log *zap.Logger, tracer opentracing.Tracer) (*RPCServer, func(), error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	endpoint := fmt.Sprintf("0.0.0.0:%d", grpc_port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, nil, err
	}

	// Initialize the gRPC server.
	newRPCServer := grpc.NewServer(
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
		Server: newRPCServer,
		Run: func() {
			// After all your registrations, make sure all of the Prometheus metrics are initialized.
			grpc_prometheus.Register(newRPCServer)

			go newRPCServer.Serve(lis)
			log.Info("Run gRPC server", zap.Int("port", grpc_port))
		},
		Endpoint: endpoint,
	}

	cleanup := func() {
		newRPCServer.GracefulStop()
	}

	return r, cleanup, err
}
