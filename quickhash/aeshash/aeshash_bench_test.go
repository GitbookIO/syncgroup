package aeshash

import (
	"testing"
)

const B = 1
const K = 1024
const M = K * K

func BenchmarkAES_16B(b *testing.B) {
	benchHash(b, 16)
}
func BenchmarkAES_32B(b *testing.B) {
	benchHash(b, 32)
}
func BenchmarkAES_64B(b *testing.B) {
	benchHash(b, 64)
}
func BenchmarkAES_128B(b *testing.B) {
	benchHash(b, 128)
}

func BenchmarkAES_256B(b *testing.B) {
	benchHash(b, 256)
}

func BenchmarkAES_512B(b *testing.B) {
	benchHash(b, 512)
}

func BenchmarkAES_1KB(b *testing.B) {
	benchHash(b, K)
}
func BenchmarkAES_4KB(b *testing.B) {
	benchHash(b, 4*K)
}

func BenchmarkAES_16KB(b *testing.B) {
	benchHash(b, 16*K)
}

func BenchmarkAES_128KB(b *testing.B) {
	benchHash(b, 128*K)
}

func BenchmarkAES_256KB(b *testing.B) {
	benchHash(b, 256*K)
}

func BenchmarkAES_512KB(b *testing.B) {
	benchHash(b, 512*K)
}
func BenchmarkAES_1MB(b *testing.B) {
	benchHash(b, 1*M)
}

func BenchmarkAES_2MB(b *testing.B) {
	benchHash(b, 2*M)
}

func BenchmarkAES_4MB(b *testing.B) {
	benchHash(b, 4*M)
}

func benchHash(b *testing.B, bufferSize int64) {
	buffer := make([]byte, bufferSize)
	b.SetBytes(bufferSize)
	for n := 0; n < b.N; n++ {
		Hash(buffer, 42)
	}
}
