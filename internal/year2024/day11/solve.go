package day11

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
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

func solve(filename string) (solution1, solution2 string, err error) {
	const blinks = 40
	//const blinks = 1
	// 18 446 744 073 709 551 615
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	arr := util.ParseUint64ArrNoError(b)
	log.Printf("bytes: %s", b)
	log.Printf("arr: %v", arr)
	res := arr
	for i := range blinks {
		res = blink(res)
		if i == 24 {
			solution1 = strconv.Itoa(len(res))
		}
		log.Println(i, "total", len(res))
	}
	
	return
}

func blink(arr []uint64) []uint64 {
	res := make([]uint64, 0, len(arr)*2)
	for _, u := range arr {
		res = newValuesAppend(res, u)
	}
	return res
}

func newValuesAppend(res []uint64, u uint64) []uint64 {
	var count, div uint64
	if u == 0 {
		return append(res, 1)
	}

	if count, div = CountDigitsDivisor(u); count%2 == 0 {
		return append(res, u/div, u%div)
	}

	return append(res, u*2024)
}

// 52 74
// 5274 / 100 = 52
// 5274 % 100 74

/// 7494 9947

func CountDigitsDivisor(u uint64) (count, div uint64) {
	// 18446744073709551615 <- max amount
	// could be done in a loop but would most likelly be slower
	switch {
	case u >= 10_000_000_000_000_000_000:
		return 20, 10_000_000_000
	case u >= 10_000_000_000_000_000_00:
		return 19, 1_000_000_000
	case u >= 10_000_000_000_000_000_0:
		return 18, 1_000_000_000
	case u >= 10_000_000_000_000_000:
		return 17, 100_000_000
	case u >= 10_000_000_000_000_00:
		return 16, 100_000_000
	case u >= 10_000_000_000_000_0:
		return 15, 10_000_000
	case u >= 10_000_000_000_000:
		return 14, 10_000_000
	case u >= 10_000_000_000_00:
		return 13, 1_000_000
	case u >= 10_000_000_000_0:
		return 12, 1_000_000
	case u >= 10_000_000_000:
		return 11, 100_000
	case u >= 10_000_000_00:
		return 10, 100_000
	case u >= 10_000_000_0:
		return 9, 10_000
	case u >= 10_000_000:
		return 8, 10_000
	case u >= 10_000_00:
		return 7, 1000
	case u >= 10_000_0:
		return 6, 1000
	case u >= 10_000:
		return 5, 100
	case u >= 10_00:
		return 4, 100
	case u >= 10_0:
		return 3, 10
	case u >= 10:
		return 2, 10
	default:
		return 1, 1
	}
	//18 446 744 073 709 551 615
}
