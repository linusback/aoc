package day7

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"strconv"
)

const (
	exampleFile        = "./internal/year2024/day7/example.txt"
	inputFile          = "./internal/year2024/day7/input.txt"
	plus        symbol = '+'
	multi       symbol = '*'
	concatenate symbol = '|'
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type symbol byte

func (s symbol) apply(a, b uint64) uint64 {
	switch s {
	case multi:
		return a * b
	case plus:
		return a + b
	case concatenate:
		str := fmt.Sprintf("%d%d", a, b)
		res, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		return res
	default:
		log.Fatalf("unknown symbol %c", s)
		return 0
	}
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

	var acc1, acc2 uint64
	for _, r := range rules {
		if isValid(r, plus, multi) {
			//log.Println(r)
			acc1 += r[0]
			acc2 += r[0]
		} else if isValid(r, plus, multi, concatenate) {
			acc2 += r[0]
		}
	}

	solution1 = strconv.FormatUint(acc1, 10)
	solution2 = strconv.FormatUint(acc2, 10)

	return
}

func isValid(r []uint64, symbols ...symbol) bool {
	//fmt.Println(r)
	res := r[0]
	acc := r[1]
	rest := r[2:]

	for o := range util.Combinate(len(rest), symbols...) {
		for i, s := range o {
			acc = s.apply(acc, rest[i])
			//fmt.Println("acc is: ", acc)
			if acc > res {
				break
			}
		}
		if acc == res {
			return true
		}
		acc = r[1]
	}
	return false
}
