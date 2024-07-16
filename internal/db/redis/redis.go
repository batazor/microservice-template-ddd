package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *redis.Client
	config Config
}

// Init ...
func (r *Store) Init(ctx context.Context) error {
	// Set configuration
	r.setConfig()

	// Connect to Redis
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.config.URI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := r.client.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (r *Store) Close() error {
	return r.client.Close()
}

// Migrate ...
func (r *Store) migrate() error { // nolint unused
	return nil
}

// setConfig - set configuration
func (r *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_REDIS_URI", "localhost:6379") // Redis URI

	r.config = Config{
		URI: viper.GetString("STORE_REDIS_URI"),
	}
}
