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

func BenchmarkShardedLockUnlock(b *testing.B) {
	rw := NewShardedMutexes()
	lock := rw.Locker("a")
	for n := 0; n < b.N; n++ {
		lock.Lock()
		lock.Unlock()
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

func BenchmarkReadShardedLockUnlock(b *testing.B) {
	rw := NewShardedMutexes()
	lock := rw.RLocker("a")
	for n := 0; n < b.N; n++ {
		lock.RLock()
		lock.RUnlock()
	}
}

// BenchmarkParallelGroup
// Runs 100 go-routines read locking/unlocing on a single key (at a time)of a group mutex
func BenchmarkParallelGroup(b *testing.B) {
	M := 4096
	b.SetParallelism(M)

	locks := NewMutexGroup()
	keys := nkeys(M)

	b.RunParallel(func(pb *testing.PB) {
		j := 0
		for pb.Next() {
			y := j
			j++
			key := keys[y%M]
			// Lock, sleep
			locks.Lock(key)
			//time.Sleep(1 * time.Millisecond)
			locks.Unlock(key)
		}
	})
}

// BenchmarkParallelSharded
// Runs 100 go-routines read locking/unlocing on a single key (at a time) of a Sharded mutex
func BenchmarkParallelSharded(b *testing.B) {
	M := 4096
	b.SetParallelism(M)

	locks := NewShardedMutexes()
	keys := nkeys(M)

	b.RunParallel(func(pb *testing.PB) {
		j := 0
		for pb.Next() {
			y := j
			j++
			key := keys[y%M]
			lock := locks.Locker(key)
			// Lock, sleep
			lock.Lock()
			//time.Sleep(1 * time.Millisecond)
			lock.Unlock()
		}
	})
}

// BenchmarkParallelSingle
// Runs 100 go-routines read locking/unlocing on a single mutex
func BenchmarkParallelSingle(b *testing.B) {
	M := 4096
	b.SetParallelism(M)

	lock := sync.RWMutex{}
	keys := nkeys(M)

	b.RunParallel(func(pb *testing.PB) {
		j := 0
		for pb.Next() {
			y := j
			j++
			_ = keys[y%M]
			// Lock, sleep
			lock.Lock()
			//time.Sleep(1 * time.Millisecond)
			lock.Unlock()
		}
	})
}

func nkeys(n int) []string {
	strs := make([]string, n)

	for idx, _ := range strs {
		strs[idx] = strconv.Itoa(idx)
	}

	return strs
}
