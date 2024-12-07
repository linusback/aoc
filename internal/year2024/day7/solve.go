package day7

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
)

type symbol byte

const (
	exampleFile        = "./internal/year2024/day7/example"
	inputFile          = "./internal/year2024/day7/input"
	plus        symbol = '+'
	multi       symbol = '*'
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
	var rules [][]uint64
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {

		rules = append(rules, util.ParseUint64ArrNoError(row))
		return nil
	})
	if err != nil {
		return
	}
	var acc1 uint64
	for _, r := range rules {
		if isValid1(r) {
			acc1 += r[0]
		}
	}
	log.Println(rules)
	return
}

func isValid1(r []uint64) bool {
	res := r[0]
	acc := r[1]
	rest := r[2:]

	return false
}
