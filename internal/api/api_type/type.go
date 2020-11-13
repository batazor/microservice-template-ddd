/*
API
*/

package api_type

import (
	"context"
	"time"

	"go.uber.org/zap"

	"robovoice-template/pkg/notify"
)

var (
	METHOD_USER_ADD    = notify.NewEventID()
	METHOD_USER_GET    = notify.NewEventID()
	METHOD_USER_LIST   = notify.NewEventID()
	METHOD_USER_UPDATE = notify.NewEventID()
	METHOD_USER_DELETE = notify.NewEventID()
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config Config, log *zap.Logger) error
}

// Config - base configuration for API
type Config struct { // nolint unused
	Port    int
	Timeout time.Duration
}
