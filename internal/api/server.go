package api

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/api/api_type"
	"robovoice-template/internal/api/http"
	user_rpc "robovoice-template/internal/user/infrastructure/rpc"
)

type Server struct{} // nolint unused

// runAPIServer - start HTTP-server
func (*Server) RunAPIServer(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) {
	var server api_type.API

	viper.SetDefault("API_PORT", 7070)  // API port
	viper.SetDefault("API_TIMEOUT", 60) // Request Timeout

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	server = &http.API{}

	// Register clients
	_, err := user_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := server.Run(ctx, config, log); err != nil {
		log.Fatal(err.Error())
	}
}
