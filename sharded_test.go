package syncgroup

import (
	"testing"

	"github.com/GitbookIO/syncgroup/quickhash"
)

func BenchmarkShardIdx(b *testing.B) {
	key := "abc"
	var N uint64 = 4
	for i := 0; i < b.N; i++ {
		_ = getShardIdx(key, N)
	}
}

func BenchmarkShardModulo(b *testing.B) {
	key := "abc"
	var N uint64 = 4
	hash := quickhash.StrHash(key)
	for i := 0; i < b.N; i++ {
		_ = hash % N
	}
}

func BenchmarkShardHash(b *testing.B) {
	key := "abc"
	for i := 0; i < b.N; i++ {
		_ = quickhash.StrHash(key)
	}
}
