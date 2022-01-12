/*
API-service
*/
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"microservice-template-ddd/internal/di"
	"microservice-template-ddd/pkg/config"
	"microservice-template-ddd/pkg/error/status"
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
	s, cleanup, err := di.InitializeAPIService(ctx)
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

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Context close
	ctx.Done()

	// Close our other dependencies
	cleanup()
}
