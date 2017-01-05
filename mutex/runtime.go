package mutex

import (
	"unsafe"
)

//go:noescape
//go:linkname race_Enabled internal.race.Enabled
var race_Enabled bool

//go:noescape
//go:linkname race_Enable internal.race.Enable
func race_Enable()

//go:noescape
//go:linkname race_Disable internal.race.Disable
func race_Disable()

//go:noescape
//go:linkname race_Acquire internal.race.Acquire
func race_Acquire(unsafe.Pointer)

//go:noescape
//go:linkname race_Release internal.race.Release
func race_Release(unsafe.Pointer)

//go:noescape
//go:linkname race_ReleaseMerge internal.race.ReleaseMerge
func race_ReleaseMerge(unsafe.Pointer)

//go:noescape
//go:linkname runtime_canSpin runtime.runtime_canSpin
func runtime_canSpin(i int) bool

//go:noescape
//go:linkname runtime_doSpin runtime.runtime_doSpin
func runtime_doSpin()

//go:noescape
//go:linkname runtime_Semacquire runtime.runtime_Semacquire
func runtime_Semacquire(s *uint32)

//go:noescape
//go:linkname runtime_Semrelease runtime.runtime_Semrelease
func runtime_Semrelease(s *uint32)
