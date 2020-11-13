/*
User-service
*/
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	billing_rpc "robovoice-template/internal/billing/infrastructure/rpc"
	book_rpc "robovoice-template/internal/book/infrastructure/rpc"
	"robovoice-template/internal/di"
	user_rpc "robovoice-template/internal/user/infrastructure/rpc"
	"robovoice-template/pkg/config"
	"robovoice-template/pkg/error/status"
)

func init() {
	// Read ENV variables
	if err := config.Init(); err != nil {
		fmt.Println(err.Error())
		os.Exit(status.ERROR_CONFIG)
	}
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeBookService(ctx)
	if err != nil { // TODO: use as helpers
		if r, ok := err.(*net.OpError); ok {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT enviroment", r.Addr.String()))
		}

		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string))
		}
	}()

	// Register user
	userService, err := user_rpc.Use(ctx, s.ClientRPC)
	if err != nil {
		s.Log.Fatal(err.Error())
	}

	billingService, err := billing_rpc.Use(ctx, s.ClientRPC)
	if err != nil {
		s.Log.Fatal(err.Error())
	}

	// Run user server
	book_rpc.New(s.ServerRPC, s.Log, s.BookStore, userService, billingService)

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Context close
	ctx.Done()

	// Close our other dependencies
	cleanup()
}
