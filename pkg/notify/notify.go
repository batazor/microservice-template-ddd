/*
Notify system
*/

package notify

import (
	"go.uber.org/atomic"
)

var (
	eventCounter atomic.Uint32
)

func NewEventID() uint32 {
	eventCounter.Inc()
	return eventCounter.Load()
}
