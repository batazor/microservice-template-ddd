//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"robovoice-template/internal/db"
	"robovoice-template/pkg/rpc"
	"robovoice-template/pkg/traicing"
)

type DefaultService struct {
	Log *zap.Logger
}

// DefaultService ======================================================================================================
var DefaultSet = wire.NewSet(InitLogger, InitTracer, NewDefaultService)

// InitConstructor =====================================================================================================
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

// InitStore return db
func InitStore(ctx context.Context, log *zap.Logger) (*db.Store, func(), error) {
	var st db.Store
	db, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

// runGRPCServer ...
func runGRPCServer(log *zap.Logger, tracer opentracing.Tracer) (*rpc.RPCServer, func(), error) {
	return rpc.InitServer(log, tracer)
}

// runGRPCClient - set up a connection to the server.
func runGRPCClient(log *zap.Logger, tracer opentracing.Tracer) (*grpc.ClientConn, func(), error) {
	return rpc.InitClient(log, tracer)
}

func NewDefaultService(log *zap.Logger) (*DefaultService, error) {
	return &DefaultService{
		Log: log,
	}, nil
}

func InitializeDefaultService(ctx context.Context) (*DefaultService, func(), error) {
	panic(wire.Build(DefaultSet))
}
