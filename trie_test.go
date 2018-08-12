package triebenchmark

import (
	"math/rand"
	"testing"

	dghubble "github.com/dghubble/trie"
	fvbock "github.com/fvbock/trie"
  gocollections "github.com/golang-collections/collections/trie"
)

var (
  dgtrie *dghubble.PathTrie
  fvtrie *fvbock.Trie
  gocollections *
	randstrings []string
)

func init() {
	randstrings = make([]string, 1000000)
	i := 0
	for i < 1000000 {
		rstr := []byte{}
		n := 0
		for n < 50 {
			rstr = append(rstr, byte(rand.Intn(255)))
			n++
		}
		randstrings[i] = string(rstr)
		i++
	}
}

func BenchmarkFVbockTrieBenchAdd(b *testing.B) {
	tr := fvbock.NewTrie()
	for x := 0; x < b.N; x++ {
		tr.Add(randstrings[x%500000])
	}
}

func BenchmarkBGHubbleTrieBenchAdd(b *testing.B) {
	tr := dghubble.NewPathTrie()
	for x := 0; x < b.N; x++ {
		tr.Put(randstrings[x%500000], struct{}{})
	}
}

func BenchmarkTrieBenchHasPrefix(b *testing.B) {
	tr := NewTrie()
	b.StopTimer()
	randstr := make([]string, 100)
	i := 0
	for i < 100000 {
		rstr := []byte{}
		n := 0
		for n < 100 {
			rstr = append(rstr, byte(rand.Intn(255)))
			n++
		}
		randstr = append(randstr, string(rstr))
		i++
	}

	for x := 0; x < 1000000; x++ {
		tr.Add(randstr[x%10000])
	}
	// fmt.Printf("Having %v distinct entries.\n", len(tr.Members()))
	b.StartTimer()
	for x := 0; x < b.N; x++ {
		tr.HasPrefix(randstr[x%100000])
	}
}

func BenchmarkTrieBenchHas(b *testing.B) {
	tr := NewTrie()
	b.StopTimer()
	randstr := make([]string, 100)
	i := 0
	for i < 100000 {
		rstr := []byte{}
		n := 0
		for n < 100 {
			rstr = append(rstr, byte(rand.Intn(255)))
			n++
		}
		randstr = append(randstr, string(rstr))
		i++
	}

	for x := 0; x < 1000000; x++ {
		tr.Add(randstr[x%10000])
	}
	// fmt.Printf("Having %v distinct entries.\n", len(tr.Members()))
	b.StartTimer()
	for x := 0; x < b.N; x++ {
		tr.Has(randstr[x%100000])
	}
}

func BenchmarkTrie1MBenchHasPrefix(b *testing.B) {
	for x := 0; x < b.N; x++ {
		tr1M.HasPrefix(randstrings[x%1000000])
	}
}

func BenchmarkTrie1MBenchHas(b *testing.B) {
	for x := 0; x < b.N; x++ {
		tr1M.Has(randstrings[x%1000000])
	}
}
