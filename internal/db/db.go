/*
Data Base package
*/
package db

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"microservice-template-ddd/internal/db/mongo"
	"microservice-template-ddd/internal/db/redis"
)

// Use return implementation of db
func (store *Store) Use(ctx context.Context, log *zap.Logger) (*Store, error) {
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "mongo":
		store.Store = &mongo.Store{}
	case "redis":
		store.Store = &redis.Store{}
	default:
		store.Store = &redis.Store{}
	}

	if err := store.Store.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run db", zap.String("db", store.typeStore))

	return store, nil
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "redis") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
