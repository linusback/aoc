package day19

import (
	"cmp"
	"github.com/linusback/aoc/pkg/util"
	"slices"
	"testing"
)

const testFilename = "./input.txt"

func Benchmark_solveTowels(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solveTowels()
	}
}

func parseInput() (err error) {
	err = util.DoEachRowFile(testFilename, parseTowels, parsePatterns)
	if err != nil {
		return
	}
	slices.SortFunc(towels, func(a, b pattern) int {
		if len(a) < len(b) {
			return -1
		}
		if len(a) > len(b) {
			return 1
		}
		return cmp.Compare(string(a), string(b))
	})
	return nil
}
