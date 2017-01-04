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

// BenchmarkParallelGroup
// Runs 100 go-routines read locking/unlocing on a single key (at a time)of a group mutex
func BenchmarkParallelGroup(b *testing.B) {
	M := 100

	locks := NewMutexGroup()
	wg := sync.WaitGroup{}
	keys := nkeys(M)

	for i := 0; i < b.N; i++ {
		for j := 0; j < M; j++ {
			key := keys[j]
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

// BenchmarkParallelSharded
// Runs 100 go-routines read locking/unlocing on a single key (at a time) of a Sharded mutex
func BenchmarkParallelSharded(b *testing.B) {
	M := 100

	locks := NewShardedMutexGroup()
	wg := sync.WaitGroup{}
	keys := nkeys(M)

	for i := 0; i < b.N; i++ {
		for j := 0; j < M; j++ {
			key := keys[j]
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

// BenchmarkParallelSingle
// Runs 100 go-routines read locking/unlocing on a single mutex
func BenchmarkParallelSingle(b *testing.B) {
	M := 100

	lock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	keys := nkeys(M)

	for i := 0; i < b.N; i++ {
		for j := 0; j < M; j++ {
			key := keys[j]
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

func nkeys(n int) []string {
	strs := make([]string, n)

	for idx, _ := range strs {
		strs[idx] = strconv.Itoa(idx)
	}

	return strs
}
