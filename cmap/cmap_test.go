package cmap_test

import (
	"container-go/cmap"
	"testing"
)

const Writer = 100
const Reader = 100
const Deleter = 10

/*
*

	goos: linux
	goarch: amd64
	cpu: Intel(R) Xeon(R) Silver 4210R CPU @ 2.40GHz
	BenchmarkMaps
	BenchmarkMaps/ConcurrentMap
	BenchmarkMaps/ConcurrentMap-14            168572              9680 ns/op            1342 B/op         97 allocs/op
	PASS
	ok      command-line-arguments  1.820s
*/
func benchmarkConcurrentMap(b *testing.B, cmap_ *cmap.ConcurrentMap[int, *int]) {
	for i := 0; i < b.N; i++ {
		go func() {
			for k := 0; k < Writer; k++ {
				cmap_.Set(k, &k)
			}
		}()

		go func() {
			for k := 0; k < Reader; k++ {
				cmap_.Get(k)
			}
		}()

		go func() {
			for k := 0; k < Deleter; k++ {
				cmap_.Remove(k)
			}
		}()
	}
}

func BenchmarkMaps(b *testing.B) {
	cmap_ := cmap.New[int, *int]()

	b.Run("ConcurrentMap", func(b *testing.B) {
		benchmarkConcurrentMap(b, &cmap_)
	})
}
