package syncgroup

import (
	"fmt"
	"sync"
)

// MutexGroup provides a group of sync.RWMutex, that be locked/unlocked by key
type MutexGroup struct {
	lock    *sync.RWMutex
	active  *ActiveGroup
	mutexes map[string]*rwmutex
}

func NewMutexGroup() *MutexGroup {
	return &MutexGroup{
		lock:    &sync.RWMutex{},
		active:  NewActiveGroup(),
		mutexes: map[string]*rwmutex{},
	}
}

func (mg *MutexGroup) Lock(key string) {
	mg.getOrCreate(key).Lock()
}

func (mg *MutexGroup) RLock(key string) {
	mg.getOrCreate(key).RLock()
}

func (mg *MutexGroup) Unlock(key string) {
	mutex := mg.getOrFail(key)
	usage := mutex.Unlock()
	if usage == 0 {
		mg.maybeDelete(key, mutex)
	}
}

func (mg *MutexGroup) RUnlock(key string) {
	mutex := mg.getOrFail(key)
	usage := mutex.RUnlock()
	if usage == 0 {
		mg.maybeDelete(key, mutex)
	}
}

func (mg *MutexGroup) Has(key string) bool {
	return mg.active.Has(key)
}

func (mg *MutexGroup) maybeDelete(key string, mutex *rwmutex) bool {
	// Is in use
	if mutex.count.Get() != 0 {
		return false
	}
	// Track if deleted
	deleted := false
	// Lock
	mutex.Lock()
	// Delete if only used by us
	if mutex.count.Get() == 1 {
		deleted = true
		// Delete from map
		mg.lock.Lock()
		delete(mg.mutexes, key)
		mg.lock.Unlock()
	}
	mutex.Unlock()

	return deleted
}

func (mg *MutexGroup) getOrFail(key string) *rwmutex {
	mg.lock.RLock()
	mutex, ok := mg.mutexes[key]
	mg.lock.RUnlock()
	// Get
	if ok {
		return mutex
	}
	panic(fmt.Sprintf(`MutexGroup.getOrFail("%s"): Tried to perform an Unlock on a key that was never Locked`, key))
}

func (mg *MutexGroup) getOrCreate(key string) *rwmutex {
	// Get
	mg.lock.RLock()
	mutex, ok := mg.mutexes[key]
	mg.lock.RUnlock()

	// Exists, so early exit
	if ok {
		return mutex
	}

	// Create
	mutex = &rwmutex{}
	mg.lock.Lock()
	mg.mutexes[key] = mutex
	mg.lock.Unlock()

	return mutex
}
