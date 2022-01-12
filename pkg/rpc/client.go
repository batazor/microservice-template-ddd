package rpc

import (
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitClient(log *zap.Logger, tracer opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	viper.SetDefault("GRPC_CLIENT_HOST", "0.0.0.0") // gRPC host
	grpc_host := viper.GetString("GRPC_CLIENT_HOST")

	// Set up a connection to the server peer
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", grpc_host, grpc_port),
		grpc.WithInsecure(), // nolint staticcheck

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

	log.Info("Run gRPC client", zap.String("host", grpc_host), zap.Int("port", grpc_port))

	cleanup := func() {
		conn.Close()
	}

	return conn, cleanup, nil
}
