package day6

import (
	"io"
	"log"
	"testing"
)

const testFilename = "./input.txt"

func Test_Solve(t *testing.T) {
	s1, s2, err := solve(testFilename)
	if err != nil {
		t.Error(err)
	}
	t.Log(s1)
	t.Log(s2)
}

func Benchmark_Day6_Solve(b *testing.B) {
	log.SetOutput(io.Discard)
	b.ReportAllocs()
	b.SetBytes(17030)
	b.ResetTimer()
	var err error
	for i := 0; i < b.N; i++ {
		_, _, err = solve(testFilename)
		if err != nil {
			b.Error(err)
		}
	}
}
