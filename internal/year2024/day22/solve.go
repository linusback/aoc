package day22

import (
	"github.com/linusback/aoc/pkg/util"
	"os"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day22/example.txt"
	exampleFile2 = "./internal/year2024/day22/exemple2.txt"
	inputFile    = "./internal/year2024/day22/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

var (
	numbers []uint64
	acc1    uint64
)

func solve(filename string) (solution1, solution2 string, err error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	numbers = util.ParseUint64ArrNoError(b)
	for _, number := range numbers {
		for range 2000 {
			number = evolveSecretNumber(number)
		}
		acc1 += number
	}
	solution1 = strconv.FormatUint(acc1, 10)
	return
}

func evolveSecretNumber(in uint64) uint64 {
	in = ((in << 6) ^ in) % 16777216
	in = ((in >> 5) ^ in) % 16777216
	in = ((in << 11) ^ in) % 16777216
	return in
}

//
//func mix(secret, in uint64) (out uint64) {
//	return secret ^ in
//}
//
//func prune(in uint64) (out uint64) {
//	return in % 16777216
//}
