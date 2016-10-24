package syncgroup

import (
	"testing"
)

func TestMutexSimple(t *testing.T) {
	group := NewMutexGroup()

	// Check that we can lock multiple keys at once
	group.Lock("a")
	group.Lock("b")
	group.Lock("c")
	group.Unlock("a")
	group.Unlock("b")
	group.Unlock("c")
}

func TestMutexDeleted(t *testing.T) {
	group := NewMutexGroup()
	keys := []string{"a", "b", "c"}

	// Lock
	for _, key := range keys {
		group.Lock(key)
	}

	// TODO: Ensure keys exist

	// Unlock
	for _, key := range keys {
		group.Unlock(key)
	}

	// Ensure keys are deleted once unlocked
	for _, key := range keys {
		if group.Has(key) {
			t.Errorf("MutexGroup: '%s' should be deleted since all of it's instances have been unlocked", key)
		}
	}
}
