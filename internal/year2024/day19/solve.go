package day19

import (
	"cmp"
	"github.com/linusback/aoc/pkg/util"
	"slices"
	"strconv"
	"strings"
)

const (
	exampleFile = "./internal/year2024/day19/example.txt"
	inputFile   = "./internal/year2024/day19/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type pattern []byte

func (t pattern) String() string {
	return util.ToUnsafeString(t)
}

func ToString(t []pattern) string {
	if len(t) == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteByte('[')
	sb.Write(t[0])
	for _, to := range t[1:] {
		sb.WriteByte(' ')
		sb.Write(to)
	}
	sb.WriteByte(']')
	return sb.String()
}

var (
	towels       []pattern
	patterns     []pattern
	knownPattern = make(map[string]uint64, 19000)
)

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseTowels, parsePatterns)
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
	solution1, solution2 = solveTowels()

	return
}

func solveTowels() (solution1, solution2 string) {
	var res1, res2 uint64
	for _, t := range patterns {
		if ways := canBeMade(t, 0); ways > 0 {
			res1++
			res2 += ways
		}
	}
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func canBeMade(pattern pattern, res uint64) uint64 {
	if len(pattern) == 0 {
		return res + 1
	}
	var (
		newWays uint64
		key     string
	)
	for _, t := range towels {
		if len(pattern) < len(t) {
			break
		}
		if notMatch(pattern, t) {
			continue
		}
		key = util.ToUnsafeString(pattern[len(t):])
		if ways, ok := knownPattern[key]; ok {
			newWays += ways
			continue
		}
		ways := canBeMade(pattern[len(t):], res)
		newWays += ways
		knownPattern[key] = ways
	}
	return res + newWays
}

func notMatch(pattern, t pattern) bool {
	for i, b := range t {
		if pattern[i] != b {
			return true
		}
	}
	return false
}

func parsePatterns(row []byte, _ int) error {
	patterns = append(patterns, row)
	return nil
}

func parseTowels(row []byte, _ int) error {
	var (
		start, i int
		b        byte
	)
	for i, b = range row {
		if util.AsciiSpace[b] == 1 {
			start++
			continue
		}
		if b == ',' {
			towels = append(towels, row[start:i])
			start = i + 1
		}
	}
	if start < i {
		towels = append(towels, row[start:])
	}
	return nil
}
