package syncgroup

import (
	"fmt"
	"sync"
)

// MutexGroup provides a group of sync.RWMutex, that be locked/unlocked by key
type MutexGroup struct {
	lock    *sync.RWMutex
	active  *ActiveGroup
	mutexes map[string]*sync.RWMutex
}

func NewMutexGroup() *MutexGroup {
	return &MutexGroup{
		lock:    &sync.RWMutex{},
		active:  NewActiveGroup(),
		mutexes: map[string]*sync.RWMutex{},
	}
}

func (mg *MutexGroup) Lock(key string) {
	mg.inc(key)
	mg.getOrCreate(key).Lock()
}

func (mg *MutexGroup) RLock(key string) {
	mg.inc(key)
	mg.getOrCreate(key).RLock()
}

func (mg *MutexGroup) Unlock(key string) {
	mg.getOrFail(key).Unlock()
	mg.dec(key)
}

func (mg *MutexGroup) RUnlock(key string) {
	mg.getOrFail(key).RUnlock()
	mg.dec(key)
}

func (mg *MutexGroup) Has(key string) bool {
	return mg.active.Has(key)
}

func (mg *MutexGroup) inc(key string) {
	mg.lock.RLock()
	defer mg.lock.RUnlock()
	mg.active.Inc(key)
}

func (mg *MutexGroup) dec(key string) {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	// No longer active
	if mg.active.Dec(key) == 0 {
		delete(mg.mutexes, key)
	}
}

func (mg *MutexGroup) getOrFail(key string) *sync.RWMutex {
	mg.lock.RLock()
	defer mg.lock.RUnlock()
	// Get
	if mutex, ok := mg.mutexes[key]; ok {
		return mutex
	}
	panic(fmt.Sprintf(`MutexGroup.getOrFail("%s"): Tried to perform an Unlock on a key that was never Locked`, key))
}

func (mg *MutexGroup) getOrCreate(key string) *sync.RWMutex {
	mg.lock.Lock()
	defer mg.lock.Unlock()

	// Create if doesn't exist
	if _, ok := mg.mutexes[key]; !ok {
		mg.mutexes[key] = &sync.RWMutex{}
	}

	return mg.mutexes[key]
}
