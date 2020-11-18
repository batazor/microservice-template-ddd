/*
User-service
*/
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"robovoice-template/internal/di"
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

	// RUN HTTP HEALTH AS EXAMPLE
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{}")
	})
	go http.ListenAndServe(":8080", nil)

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
