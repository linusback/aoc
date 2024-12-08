package day7

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"maps"
	"slices"
)

const (
	exampleFile        = "./internal/year2024/day7/example"
	inputFile          = "./internal/year2024/day7/input"
	plus        symbol = '+'
	multi       symbol = '*'
)

type symbol byte

func (s symbol) apply(a, b uint64) uint64 {
	switch s {
	case multi:
		return a * b
	case plus:
		return a + b
	default:
		log.Fatalf("unknown symbol %c", s)
		return 0
	}
}

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
	var rules [][]uint64
	ruleLen := make(map[int]int)
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		rule := util.ParseUint64ArrNoError(row)
		rules = append(rules, rule)
		if _, ok := ruleLen[len(rule)]; !ok {
			ruleLen[len(rule)] = 0
		}
		ruleLen[len(rule)]++
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
	keys := slices.Collect(maps.Keys(ruleLen))
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println("len: ", k, ", count: ", ruleLen[k])
	}

	return
}

func isValid1(r []uint64) bool {
	//res := r[0]
	//acc := r[1]
	//rest := r[2:]

	return false
}
