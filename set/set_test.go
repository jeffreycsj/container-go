package set_test

import (
	"container-go/set"
	"testing"
)

/*
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) Silver 4210R CPU @ 2.40GHz
BenchmarkMaps
BenchmarkMaps/Set
BenchmarkMaps/Set-14            39211989                34.30 ns/op            0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.384s
*/
func benchmarkSet(b *testing.B, set_ *set.Set[int]) {
	for i := 0; i < b.N; i++ {
		set_.Add(i)
		set_.Remove(i)
	}
}

func BenchmarkMaps(b *testing.B) {
	set_ := set.NewSet[int]()

	b.Run("Set", func(b *testing.B) {
		benchmarkSet(b, set_)
	})
}
