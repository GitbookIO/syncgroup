package syncgroup

import (
	"sync"
	"sync/atomic"
)

// rwmutex is a RWMutex that tracks it's use
// and after each lock/unlock reports it's current usage
type rwmutex struct {
	count counter
	mutex sync.RWMutex
}

func (rw *rwmutex) RLock() int64 {
	active := rw.count.Inc()
	rw.mutex.RLock()
	return active
}

func (rw *rwmutex) Lock() int64 {
	count := rw.count.Inc()
	rw.mutex.Lock()
	return count
}

func (rw *rwmutex) RUnlock() int64 {
	rw.mutex.RUnlock()
	return rw.count.Dec()
}

func (rw *rwmutex) Unlock() int64 {
	rw.mutex.Unlock()
	return rw.count.Dec()
}

// counter is a simple atomically safe counter
type counter struct {
	count int64
}

func (c *counter) Inc() int64 {
	return atomic.AddInt64(&c.count, +1)
}

func (c *counter) Dec() int64 {
	return atomic.AddInt64(&c.count, -1)
}

func (c *counter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}
