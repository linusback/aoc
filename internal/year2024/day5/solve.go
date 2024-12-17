package day5

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"slices"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day5/example.txt"
	inputFile   = "./internal/year2024/day5/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile, getComparer1)
}

func solve(filename string, cmpFunc func([][2]int64) func(int64, int64) int) (solution1, solution2 string, err error) {
	var (
		rules   [][2]int64
		updates [][]int64
	)
	parseRules := func(row []byte, nr int) error {
		res := util.ParseInt64ArrNoError(row)
		if len(res) != 2 {
			return fmt.Errorf("expected len 2 got %d", len(res))
		}
		rules = append(rules, *(*[2]int64)(res))
		return nil
	}
	parseUpdates := func(row []byte, nr int) error {
		updates = append(updates, util.ParseInt64ArrNoError(row))
		return nil
	}

	err = util.DoEachRowFile(filename, parseRules, parseUpdates)
	if err != nil {
		return
	}

	cmp := cmpFunc(rules)
	var acc1, acc2 int64
	for _, u := range updates {
		if slices.IsSortedFunc(u, cmp) {
			acc1 += u[len(u)/2]
			continue
		}
		slices.SortFunc(u, cmp)
		acc2 += u[len(u)/2]
	}

	solution1 = strconv.FormatInt(acc1, 10)
	solution2 = strconv.FormatInt(acc2, 10)
	return
}

func getComparer1(rules [][2]int64) func(int64, int64) int {
	return func(a int64, b int64) int {
		for _, r := range rules {
			if a == r[0] && b == r[1] {
				return -1
			}
			if a == r[1] && b == r[0] {
				return 1
			}
		}
		return 0
	}
}

func getComparer2(rules [][2]int64) func(int64, int64) int {
	var x, y int64
	return func(a int64, b int64) int {
		for _, r := range rules {
			x = r[0]
			y = r[1]
			if a == x && b == y {
				return -1
			}
			if a == y && b == x {
				return 1
			}
		}
		return 0
	}
}

func getComparer3(rules [][2]int64) func(int64, int64) int {
	slices.SortFunc(rules, sortRules)
	return func(a int64, b int64) int {
		for _, r := range rules {
			if r[0] > a && r[0] > b {
				return 0
			}
			if a == r[0] && b == r[1] {
				return -1
			}
			if a == r[1] && b == r[0] {
				return 1
			}
		}
		return 0
	}
}

func sortRules(a, b [2]int64) int {
	if a[0] == b[0] {
		if a[1] < b[1] {
			return -1
		}
		if a[1] > b[1] {
			return 1
		}
		return 0
	}
	if a[0] < b[0] {
		return -1
	}
	if a[0] > b[0] {
		return 1
	}
	return 0
}
