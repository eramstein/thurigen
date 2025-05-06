package ng

import "sync/atomic"

var nextID uint64

func getNextID() uint64 {
	return atomic.AddUint64(&nextID, 1)
}
