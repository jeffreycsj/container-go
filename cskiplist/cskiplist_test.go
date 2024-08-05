package cskiplist_test

import (
	"container-go/cskiplist"
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
	BenchmarkMaps/ConcurrentSkipList1
	BenchmarkMaps/ConcurrentSkipList1-14               22560             61016 ns/op            2562 B/op        147 allocs/op
	BenchmarkMaps/ConcurrentSkipList2
	BenchmarkMaps/ConcurrentSkipList2-14              395006            104886 ns/op            4614 B/op        351 allocs/op
	PASS
	ok      command-line-arguments  49.485s
*/
func benchmarkCSkipList1(b *testing.B, cskiplist *cskiplist.ConcurrentSkipList) {
	for i := 0; i < b.N; i++ {
		go func() {
			for k := 0; k < Writer; k++ {
				cskiplist.Set(k, k+100)
			}
		}()

		go func() {
			for k := 0; k < Reader; k++ {
				cskiplist.Get(k)
			}
		}()
	}
}

func benchmarkCSkipList2(b *testing.B, cskiplist *cskiplist.ConcurrentSkipList) {
	for i := 0; i < b.N; i++ {
		go func() {
			for k := 0; k < Deleter; k++ {
				cskiplist.Get(k)
			}
		}()

		go func() {
			for k := 0; k < Writer; k++ {
				cskiplist.Set(k, k+1000)
			}
		}()

		go func() {
			for k := 0; k < Reader; k++ {
				cskiplist.Get(k)
			}
		}()
	}
}

func BenchmarkMaps(b *testing.B) {
	cskiplist := cskiplist.NewConcurrentSkipList(func(key1, key2 any) bool {
		return key1.(int) < key2.(int)
	})

	b.Run("ConcurrentSkipList1", func(b *testing.B) {
		benchmarkCSkipList1(b, cskiplist)
	})

	b.Run("ConcurrentSkipList2", func(b *testing.B) {
		benchmarkCSkipList2(b, cskiplist)
	})
}
