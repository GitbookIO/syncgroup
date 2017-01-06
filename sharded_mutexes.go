package syncgroup

import (
	"sync"

	"github.com/GitbookIO/syncgroup/quickhash"
)

// ShardedMutexes shards calls to a list of sync.RWMutex by key
type ShardedMutexes struct {
	shards []*sync.RWMutex
	n      uint64
}

func NewShardedMutexes() *ShardedMutexes {
	var n uint64 = 2048
	shards := make([]*sync.RWMutex, n)
	// Init shard
	for idx, _ := range shards {
		shards[idx] = &sync.RWMutex{}
	}

	return &ShardedMutexes{
		shards: shards,
		n:      n,
	}
}

func (sg *ShardedMutexes) Lock(key string) {
	sg.getShard(key).Lock()
}

func (sg *ShardedMutexes) RLock(key string) {
	sg.getShard(key).RLock()
}

func (sg *ShardedMutexes) Unlock(key string) {
	sg.getShard(key).Unlock()
}

func (sg *ShardedMutexes) RUnlock(key string) {
	sg.getShard(key).RUnlock()
}

func (sg *ShardedMutexes) Locker(key string) *sync.RWMutex {
	return sg.getShard(key)
}

func (sg *ShardedMutexes) RLocker(key string) *sync.RWMutex {
	return sg.getShard(key)
}

func (sg *ShardedMutexes) getShard(key string) *sync.RWMutex {
	hash := quickhash.AesHash(key)
	return sg.shards[hash%2048]
}
