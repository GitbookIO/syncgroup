package quickhash

import (
	_ "runtime"
	"unsafe" // required to use //go:linkname
)

//go:noescape
//go:linkname useAeshash runtime.useAeshash
var useAeshash bool

//go:noescape
//go:linkname strhash runtime.strhash
func strhash(a unsafe.Pointer, h uintptr) uintptr

//go:noescape
//go:linkname aeshashstr runtime.aeshashstr
func aeshashstr(p unsafe.Pointer, h uintptr) uintptr

func StrHash(str string) uint64 {
	if useAeshash {
		return uint64(aeshashstr(unsafe.Pointer(&str), 44))
	}
	return uint64(strhash(unsafe.Pointer(&str), 44))
}

func AesHash(str string) uint64 {
	return uint64(aeshashstr(unsafe.Pointer(&str), 44))
}
