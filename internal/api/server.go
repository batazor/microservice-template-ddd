package api

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"microservice-template-ddd/internal/api/api_type"
	"microservice-template-ddd/internal/api/http"
	"microservice-template-ddd/internal/billing/infrastructure/rpc"
	"microservice-template-ddd/internal/book/infrastructure/rpc"
	"microservice-template-ddd/internal/user/infrastructure/rpc"
)

type Server struct{} // nolint unused

// runAPIServer - start HTTP-server
func (s *Server) RunAPIServer(ctx context.Context, log *zap.Logger, rpcClient *grpc.ClientConn) {
	var server api_type.API

	viper.SetDefault("API_PORT", 7070)  // API port
	viper.SetDefault("API_TIMEOUT", 60) // Request Timeout

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	// Register user
	userService, err := user_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	billingService, err := billing_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	bookService, err := book_rpc.Use(ctx, rpcClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	server = &http.API{
		Log:            log,
		UserService:    userService,
		BillingService: billingService,
		BookService:    bookService,
	}

	if err := server.Run(ctx, config); err != nil {
		log.Fatal(err.Error())
	}
}
