package syncgroup

import (
	"strconv"
	"sync"
	"time"

	"testing"
)

func BenchmarkGoLockUnlock(b *testing.B) {
	rw := sync.RWMutex{}
	for n := 0; n < b.N; n++ {
		rw.Lock()
		rw.Unlock()
	}
}

func BenchmarkGroupLockUnlock(b *testing.B) {
	rw := NewMutexGroup()
	for n := 0; n < b.N; n++ {
		rw.Lock("a")
		rw.Unlock("a")
	}
}

func BenchmarkReadGoLockUnlock(b *testing.B) {
	rw := sync.RWMutex{}
	for n := 0; n < b.N; n++ {
		rw.RLock()
		rw.RUnlock()
	}
}

func BenchmarkReadGroupLockUnlock(b *testing.B) {
	rw := NewMutexGroup()
	for n := 0; n < b.N; n++ {
		rw.RLock("a")
		rw.RUnlock("a")
	}
}

func TestRwMutexes(t *testing.T) {
	N := 100
	M := 500

	locks := NewMutexGroup()
	wg := sync.WaitGroup{}

	for i := 0; i < N; i++ {
		key := strconv.Itoa(i)
		for j := 0; j < M; j++ {
			wg.Add(1)
			// Lock, sleep
			go func(key string) {
				locks.RLock(key)
				time.Sleep(1 * time.Millisecond)
				locks.RUnlock(key)
				wg.Done()
			}(key)
		}
	}

	// Wait for all to finish
	wg.Wait()
}
