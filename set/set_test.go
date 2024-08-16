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
BenchmarkMaps/Set-14             7641049               163.6 ns/op             7 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.418s
*/
func benchmarkSet(b *testing.B, set_ *set.Set) {
	for i := 0; i < b.N; i++ {
		set_.Add(i)
		set_.Remove(i)
	}
}

func BenchmarkMaps(b *testing.B) {
	set_ := set.NewSet()

	b.Run("Set", func(b *testing.B) {
		benchmarkSet(b, set_)
	})
}
