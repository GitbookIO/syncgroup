package quickhash

import (
	"reflect"
	_ "runtime"
	"unsafe" // required to use //go:linkname
)

// Fixed seed for deterministic outcomes
const shSeed uint64 = 42

// Modified from runtime/alg.go
func rthash(ptr unsafe.Pointer, size int, seed uint64) uint64 {
	if size == 0 {
		return seed
	}
	// The runtime hasher only works on uintptr. For 64-bit
	// architectures, we use the hasher directly. Otherwise,
	// we use two parallel hashers on the lower and upper 32 bits.
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return uint64(runtime_memhash(ptr, uintptr(seed), uintptr(size)))
	}
	lo := runtime_memhash(ptr, uintptr(seed), uintptr(size))
	hi := runtime_memhash(ptr, uintptr(seed>>32), uintptr(size))
	return uint64(hi)<<32 | uint64(lo)
}

//go:linkname runtime_memhash runtime.memhash
//go:noescape
func runtime_memhash(p unsafe.Pointer, seed, s uintptr) uintptr

func StrHash(str string) uint64 {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	ptr := unsafe.Pointer(hdr.Data)
	return uint64(rthash(ptr, len(str), shSeed))
}

func ByteHash(data []byte) uint64 {
	return uint64(rthash(unsafe.Pointer(&data[0]), len(data), shSeed))
}
