package util

import (
	"iter"
	"runtime"
	"slices"
	"sync"
)

func SliceToChannel[S ~[]V, V any](s S, buffer int) (*sync.WaitGroup, <-chan V) {
	return Seq2ToChannel(slices.All(s), buffer)
}

func Seq2ToChannel[K comparable, V any](s iter.Seq2[K, V], buffer int) (*sync.WaitGroup, <-chan V) {
	if buffer < 0 {
		buffer = runtime.NumCPU()
	}
	wg := new(sync.WaitGroup)
	ch := make(chan V, buffer)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return wg, ch
}

func SeqToChannel[V any](s iter.Seq[V], buffer int) (*sync.WaitGroup, <-chan V) {
	if buffer < 0 {
		buffer = runtime.NumCPU()
	}
	wg := new(sync.WaitGroup)
	ch := make(chan V, buffer)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range s {
			ch <- v
		}
		close(ch)
	}()
	return wg, ch
}
