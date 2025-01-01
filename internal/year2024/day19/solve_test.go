package day19

import (
	"bytes"
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"slices"
	"testing"
)

const testFilename = "./input.txt"
const patternSize = 20644

var (
	matchTests   []matchTest
	towelMap     [826][]pattern
	oneStripeMap [26]uint64
	towels       []pattern
)

type matchTest struct {
	name   string
	towels []pattern
}

type hashFunc func(byte) uint16

func Benchmark_solveTowels(b *testing.B) {
	//err := parseInput()
	//if err != nil {
	//	b.Error(err)
	//}
	//b.ReportAllocs()
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//_, _ = solveTowels()
		_, _, _ = solve(testFilename)
	}
}

func Benchmark_Input1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		patterns = patterns[:0]
		err := util.DoEachRowFile(testFilename, parseTowels, parsePatterns)
		if err != nil {
			return
		}
	}
}

func Benchmark_Input2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		towelBuff = towelBuff[:0]
		patternBuff = patternBuff[:0]
		transform = transformTowels
		err := util.DoEachByteFile(testFilename, TransformTowel)
		if err != nil {
			return
		}
	}
}

func Benchmark_Hash(b *testing.B) {
	err := parseInput()
	if err != nil {
		b.Error(err)
	}
	b.Run("rust", benchHash(perfectHashRust))
	b.Run("switch", benchHash(perfectHashSwitch))
}

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
	for i, u := range oneStripeMap {
		if u == 0 {
			continue
		}
		towels = append(towels, []byte{byte(i)})
	}
	for _, t := range towelMap {
		if len(t) == 0 {
			continue
		}
		towels = append(towels, t...)
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

func notMatch(pattern, t pattern) bool {
	for i := 2; i < len(t); i++ {
		if pattern[i] != t[i] {
			return true
		}
	}
	return false
}

func notMatch2(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	return !bytes.Equal(t, pattern[:len(t)])
}

func notMatch3(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	return util.ToUnsafeString(t) != util.ToUnsafeString(pattern[:len(t)])
}

func notMatch4(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	switch len(t) {
	case 1:
		return false
	case 2:
		return pattern[1] != t[1]
	case 3:
		return pattern[1] != t[1] || pattern[2] != t[2]
	default:
		for i := 1; i < len(t); i++ {
			if pattern[i] != t[i] {
				return true
			}
		}
		return false
	}
}

func benchHash(f hashFunc) func(*testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()
		for _, p := range patterns {
			for i := 0; i < len(p); i++ {
				for _, by := range p[i:] {
					f(by)
				}
			}
		}
	}
}

func perfectHashRust(b byte) uint16 {
	return uint16((b ^ (b >> 4)) & 7)
}

func perfectHashSwitch(b byte) uint16 {
	switch b {
	case 'b':
		return 0
	case 'g':
		return 1
	case 'r':
		return 2
	case 'u':
		return 3
	case 'w':
		return 4
	default:
		panic("unknown pattern")
	}
}
