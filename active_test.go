package syncgroup

import (
	"testing"
)

func TestActiveGroupSimple(t *testing.T) {
	ag := NewActiveGroup()

	// Do basic inc/dec
	ag.Inc("abc")
	go func() {
		ag.Dec("abc")
	}()
	ag.WaitUntilFree("abc")

	// Key should no longer exist
	if ag.Has("abc") {
		t.Errorf("'abc' key should no longer exist")
	}

	// Dummy key
	if ag.Has("never_created") {
		t.Errorf("Never created keys should not exist either")
	}
}
