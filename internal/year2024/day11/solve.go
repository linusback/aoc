package day11

import (
	"github.com/linusback/aoc/pkg/util"
	"os"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day11/example.txt"
	example2File = "./internal/year2024/day11/example2.txt"
	inputFile    = "./internal/year2024/day11/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

const (
	blinks1 = 25
	blinks2 = 75
)

type key struct {
	stone uint64
	depth uint8
}

var cache = map[key]uint64{}

func solve(filename string) (solution1, solution2 string, err error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	arr := util.ParseUint64ArrNoError(b)

	var total uint64
	for _, u := range arr {
		total += countValuesCached(u, key{}, blinks1)
	}

	solution1 = strconv.FormatUint(total, 10)

	total = 0
	for _, u := range arr {
		total += countValuesCached(u, key{}, blinks2)
	}
	solution2 = strconv.FormatUint(total, 10)

	return
}

func countValuesCached(u uint64, k key, depth uint8) uint64 {
	if depth == 0 {
		return 1
	}
	depth--
	k.stone = u
	k.depth = depth
	if cached, ok := cache[k]; ok {
		return cached
	}
	var val uint64
	if u == 0 {
		val = countValuesCached(1, k, depth)
	} else if count, div := CountDigitsDivisor2(u); count%2 == 0 {
		val = countValuesCached(u/div, k, depth) + countValuesCached(u%div, k, depth)
	} else {
		val = countValuesCached(u*2024, k, depth)
	}
	cache[k] = val
	return val
}

func CountDigitsDivisor2(u uint64) (count, div uint64) {
	// 18446744073709551615 <- max amount
	// could be done in a loop but would most likely be slower
	switch {
	case u < 10:
		return 1, 1
	case u < 100:
		return 2, 10
	case u < 1_000:
		return 3, 10
	case u < 10_000:
		return 4, 100
	case u < 100_000:
		return 5, 100
	case u < 1_000_000:
		return 6, 1_000
	case u < 10_000_000:
		return 7, 1_000
	case u < 100_000_000:
		return 8, 10_000
	case u < 1_000_000_000:
		return 9, 10_000
	case u < 10_000_000_000:
		return 10, 100_000
	case u < 100_000_000_000:
		return 11, 100_000
	case u < 1_000_000_000_000:
		return 12, 1_000_000
	case u < 10_000_000_000_000:
		return 13, 1_000_000
	case u < 100_000_000_000_000:
		return 14, 10_000_000
	case u < 1_000_000_000_000_000:
		return 15, 10_000_000
	case u < 10_000_000_000_000_000:
		return 16, 100_000_000
	case u < 100_000_000_000_000_000:
		return 17, 100_000_000
	case u < 1_000_000_000_000_000_000:
		return 18, 1_000_000_000
	case u < 10_000_000_000_000_000_000:
		return 19, 1_000_000_000

	default:
		return 20, 10_000_000_000
	}
}
