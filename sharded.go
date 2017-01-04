package syncgroup

import (
	"github.com/GitbookIO/syncgroup/quickhash"
)

// ShardedMutexGroup provides a group of sync.RWMutex, that be locked/unlocked by key
type ShardedMutexGroup struct {
	shards []*MutexGroup
	n      uint64
}

func NewShardedMutexGroup() *ShardedMutexGroup {
	var n uint64 = 8
	shards := make([]*MutexGroup, n)
	// Init shard
	for idx, _ := range shards {
		shards[idx] = NewMutexGroup()
	}

	return &ShardedMutexGroup{
		shards: shards,
		n:      n,
	}
}

func (sg *ShardedMutexGroup) Lock(key string) {
	sg.getShard(key).Lock(key)
}

func (sg *ShardedMutexGroup) RLock(key string) {
	sg.getShard(key).RLock(key)
}

func (sg *ShardedMutexGroup) Unlock(key string) {
	sg.getShard(key).Unlock(key)
}

func (sg *ShardedMutexGroup) RUnlock(key string) {
	sg.getShard(key).RUnlock(key)
}

func (sg *ShardedMutexGroup) Has(key string) bool {
	return sg.getShard(key).Has(key)
}

func (sg *ShardedMutexGroup) getShard(key string) *MutexGroup {
	hash := quickhash.StrHash(key)
	return sg.shards[hash%sg.n]
}

func getShardIdx(key string, N uint64) uint64 {
	hash := quickhash.StrHash(key)
	return hash % N
}
