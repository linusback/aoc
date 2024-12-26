package day19

import (
	"fmt"
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

type matchTest struct {
	name   string
	towels []pattern
}

var matchTests []matchTest

func Benchmark_notMatch(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	for _, test := range matchTests {
		b.Run(test.name, runBenchmarkMatcher(test, notMatch))
	}
}

func Benchmark_notMatch2(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	for _, test := range matchTests {
		b.Run(test.name, runBenchmarkMatcher(test, notMatch2))
	}
}

func Benchmark_notMatch3(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	for _, test := range matchTests {
		b.Run(test.name, runBenchmarkMatcher(test, notMatch3))
	}
}

func Benchmark_notMatch4(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	for _, test := range matchTests {
		b.Run(test.name, runBenchmarkMatcher(test, notMatch4))
	}
}

func runBenchmarkMatcher(mt matchTest, m func(pattern, pattern) bool) func(b *testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, t := range mt.towels {
				tLen := len(t)
				for _, p := range patterns {
					for k := 0; k < len(p); k++ {
						pp := p[k:]
						if len(pp) < tLen {
							break
						}
						_ = m(pp, t)
					}
				}
			}
		}
	}
}

func parseInput() (err error) {
	if len(towels) > 0 {
		return
	}
	matchTests = matchTests[:0]
	err = util.DoEachRowFile(testFilename, parseTowels, parsePatterns)
	if err != nil {
		return
	}
	if len(matchTests) > 0 {
		return nil
	}
	var (
		idx  int
		name string
	)
	for _, towel := range towels {
		name = fmt.Sprintf("len-%d", len(towel))
		idx = slices.IndexFunc(matchTests, func(test matchTest) bool {
			return test.name == name
		})
		if idx >= 0 {
			matchTests[idx].towels = append(matchTests[idx].towels, towel)
		} else {
			matchTests = append(matchTests, matchTest{
				name:   name,
				towels: []pattern{towel},
			})
		}
	}
	return nil
}
