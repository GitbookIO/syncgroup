package syncgroup

import (
	"sync"
	"sync/atomic"
)

// ActiveGroup is like a smart grouped version of sync.WaitGroup
type ActiveGroup struct {
	groups map[string]*ActiveCounter
	lock   sync.RWMutex
}

func NewActiveGroup() *ActiveGroup {
	return &ActiveGroup{
		groups: map[string]*ActiveCounter{},
		lock:   sync.RWMutex{},
	}
}

func (ag *ActiveGroup) IsActive(key string) bool {
	// False for inexistant key
	if !ag.Has(key) {
		return false
	}
	return ag.get(key).IsActive()
}

func (ag *ActiveGroup) WaitUntilFree(key string) {
	ag.get(key).WaitUntilFree()
}

func (ag *ActiveGroup) Inc(key string) int64 {
	return ag.get(key).Inc()
}

func (ag *ActiveGroup) Dec(key string) int64 {
	// Get counter
	ac := ag.get(key)
	// Decrement counter
	count := ac.Dec()
	if count == 0 {
		return ag.del(key)
	}
	return count
}

func (ag *ActiveGroup) get(key string) *ActiveCounter {
	ag.lock.Lock()
	defer ag.lock.Unlock()

	// Create if doesn't exist
	if !ag.has(key) {
		ag.groups[key] = &ActiveCounter{}
	}

	return ag.groups[key]
}

func (ag *ActiveGroup) Has(key string) bool {
	ag.lock.RLock()
	defer ag.lock.RUnlock()
	return ag.has(key)
}

func (ag *ActiveGroup) has(key string) bool {
	_, ok := ag.groups[key]
	return ok
}

func (ag *ActiveGroup) del(key string) int64 {
	ag.lock.Lock()
	defer ag.lock.Unlock()

	if counter, ok := ag.groups[key]; ok && counter.cnt == 0 {
		delete(ag.groups, key)
	} else if ok {
		// Ok but not zero counter
		return counter.cnt
	}

	// Already deleted
	return 0
}

type ActiveCounter struct {
	cnt int64
	grp sync.WaitGroup
}

func (ac *ActiveCounter) Inc() int64 {
	ac.grp.Add(1)
	return atomic.AddInt64(&ac.cnt, 1)
}

func (ac *ActiveCounter) Dec() int64 {
	count := atomic.AddInt64(&ac.cnt, -1)
	ac.grp.Done()
	return count
}

func (ac *ActiveCounter) IsActive() bool {
	return atomic.LoadInt64(&ac.cnt) > 0
}

func (ac *ActiveCounter) WaitUntilFree() {
	ac.grp.Wait()
}
