package api

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/api/api_type"
	"robovoice-template/internal/api/http"
	user_rpc2 "robovoice-template/internal/user/domain"
	user_rpc "robovoice-template/internal/user/infrastructure/rpc"
)

type Server struct {
	UserService user_rpc2.UserServiceClient
} // nolint unused

// runAPIServer - start HTTP-server
func (s *Server) RunAPIServer(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) {
	var server api_type.API

	viper.SetDefault("API_PORT", 7070)  // API port
	viper.SetDefault("API_TIMEOUT", 60) // Request Timeout

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	// Register clients
	userService, err := user_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	server = &http.API{
		Log:         log,
		UserService: userService,
	}

	if err := server.Run(ctx, config); err != nil {
		log.Fatal(err.Error())
	}
}
