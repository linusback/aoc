package day22

import (
	"github.com/linusback/aoc/pkg/util"
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
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		numbers = append(numbers, util.ParseUint64ArrNoError(row)...)
		return nil
	})
	if err != nil {
		return
	}
	for _, number := range numbers {
		for range 2000 {
			number = evolveSecretNumber(number)
		}
		acc1 += number
	}
	solution1 = strconv.FormatUint(acc1, 10)
	return
}

func evolveSecretNumber(in uint64) (out uint64) {
	out = in * 64
	in = mix(out, in)
	in = prune(in)

	out = in / 32
	in = mix(out, in)
	in = prune(in)

	out = in * 2048
	in = mix(out, in)
	in = prune(in)
	return in

}

func mix(secret, in uint64) (out uint64) {
	return secret ^ in
}

func prune(in uint64) (out uint64) {
	return in % 16777216
}
