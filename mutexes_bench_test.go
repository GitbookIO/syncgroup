package syncgroup

import (
	"strconv"
	"sync"
	//	"time"

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

// BenchmarkGroupMutexes
// Runs 100 go-routines read locking/unlocing on a single key (at a time)of a group mutex
func BenchmarkGroupMutexes(b *testing.B) {
	M := 100

	locks := NewMutexGroup()
	wg := sync.WaitGroup{}

	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		for j := 0; j < M; j++ {
			wg.Add(1)
			// Lock, sleep
			go func(key string) {
				locks.RLock(key)
				//time.Sleep(1 * time.Millisecond)
				locks.RUnlock(key)
				wg.Done()
			}(key)
		}

		// Wait for all to finish
		wg.Wait()
	}
}

// BenchmarkEmptyParallel
// Runs 100 go-routines read locking/unlocing on a single mutex
func BenchmarkEmptyParallel(b *testing.B) {
	M := 100

	lock := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		for j := 0; j < M; j++ {
			wg.Add(1)
			// Lock, sleep
			go func(key string) {
				lock.RLock()
				//time.Sleep(1 * time.Millisecond)
				lock.RUnlock()
				wg.Done()
			}(key)
		}

		// Wait for all to finish
		wg.Wait()
	}
}
