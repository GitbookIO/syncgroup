package quickhash

import (
	"unsafe"

	"testing"
)

func TestStrHash(t *testing.T) {
	x0 := "abc"
	b := "b"
	b2 := "a" + b + "c"
	x1 := StrHash(x0)
	x2 := StrHash(x0)
	x3 := StrHash(b2)
	y0 := "def"
	y1 := StrHash(y0)
	y2 := StrHash(y0)

	if x1 != x2 || x2 != x3 {
		t.Errorf("x: should all be equal: %d - %d - %d", x1, x2, x3)
	}
	if !(unsafe.Pointer(&x0) != unsafe.Pointer(&b2) && x1 == x3) {
		t.Errorf("Hash should work on string value not pointer")
	}
	if y1 == x1 {
		t.Errorf("Different input strings should nearly 'always' have different hashes: %d - %d", x1, y1)
	}
	if y1 != y2 {
		t.Errorf("y: should all be equal: %d - %d", y1, y2)
	}
}

func TestStrVSByte(t *testing.T) {
	keys := []string{
		"a",
		"abc",
		"123456",
		"blablabla",
	}

	for _, k := range keys {
		h1 := StrHash(k)
		h2 := ByteHash([]byte(k))
		if h1 != h2 {
			t.Errorf("Expected str & byte hashes to be equal got: %d vs %d", h1, h2)
		}
	}
}

func BenchmarkStrHash(b *testing.B) {
	x0 := "abcdefghijklmnopqrstuvwxyz"
	for n := 0; n < b.N; n++ {
		StrHash(x0)
	}
}

func BenchmarkAesHash(b *testing.B) {
	x0 := "abcdefghijklmnopqrstuvwxyz"
	for n := 0; n < b.N; n++ {
		StrHash(x0)
	}
}

func BenchmarkFnvHash(b *testing.B) {
	x0 := "abcdefghijklmnopqrstuvwxyz"
	for n := 0; n < b.N; n++ {
		FnvStr(x0)
	}
}

func BenchmarkMapHash(b *testing.B) {
	key := "abcdefghijklmnopqrstuvwxyz"
	t := map[string]int64{
		key: 99,
	}
	for n := 0; n < b.N; n++ {
		_ = t[key]
	}
}

func FnvStr(str string) uint64 {
	return Fnv64a([]byte(str))
}
