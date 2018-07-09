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
		return uint64(aeshashstr(unsafe.Pointer(&str), 0))
	}
	return uint64(strhash(unsafe.Pointer(&str), 0))
}

func AesHash(str string) uint64 {
	return uint64(aeshashstr(unsafe.Pointer(&str), 0))
}

func ByteHash(data []byte) uint64 {
	if useAeshash {
		return uint64(aeshashstr(unsafe.Pointer(&data), 0))
	}
	return uint64(strhash(unsafe.Pointer(&data), 0))
}

func AesByteHash(data []byte) uint64 {
	return uint64(aeshashstr(unsafe.Pointer(&data), 0))
}
