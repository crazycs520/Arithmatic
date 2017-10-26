package bitcount

import (
	"math/rand"
	"testing"
)

func TestBitCount(t *testing.T) {
	f := []func(x uint64) (n int){}
	f = append(f, BitCount0)
	f = append(f, BitCount1)
	f = append(f, BitCount2)
	f = append(f, BitCount3)
	//	f = append(f, BitCount4)
	fnum := len(f)
	result := make([]int, fnum)
	for i := 0; i < 100; i++ {
		num := rand.Uint64()
		for index, fc := range f {
			result[index] = fc(num)
		}
		for j := 0; j < fnum-1; j++ {
			if result[j] != result[j+1] {
				t.Errorf("num=0x%x\n", num)
				for in, v := range result {
					t.Errorf(" result[%d]=%d", in, v)
				}

			}
		}
	}
}

func BenchmarkBitCount0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount0(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount1(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount3(0x1234567890ABCDEF)
	}
}

/*
func BenchmarkBitCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount4(0x1234567890ABCDEF)
	}
}
*/

// go version go1.9 darwin/amd64   2.9G Intel Core i7 7820
//go test -bench=.
//goos: darwin
//goarch: amd64
//BenchmarkBitCount0-8    30000000                41.5 ns/op
//BenchmarkBitCount1-8    100000000               15.0 ns/op
//BenchmarkBitCount2-8    2000000000               0.57 ns/op
//BenchmarkBitCount3-8    2000000000               0.28 ns/op
