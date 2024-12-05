package day1

import (
	"errors"
	"github.com/linusback/aoc/pkg/util"
	"slices"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day1/example"
	inputFile   = "./internal/year2024/day1/input"
)

type occurrences [][2]int64

func Solve() (solution1, solution2 string, err error) {
	col1, col2 := make([]int64, 0), make([]int64, 0)
	err = util.DoEachRowFile(inputFile, func(row []byte, nr int) error {
		u := util.ParseInt64ArrNoError(row)
		if len(u) < 2 {
			return nil
		}
		col1 = append(col1, u[0])
		col2 = append(col2, u[1])
		return nil
	})
	if err != nil {
		return
	}
	if len(col1) != len(col2) {
		err = errors.New("column length missmatch")
		return
	}
	if len(col1) == 0 {
		err = errors.New("0 values in columns")
		return
	}

	slices.Sort(col1)
	slices.Sort(col2)
	occ := countOccurrences(col2)
	//fmt.Println(col2)
	//fmt.Println(occ)
	var acc, acc2 int64
	for i := 0; i < len(col1); i++ {
		a := col1[i]
		b := col2[i]
		if a < b {
			acc += b - a
		}
		if b < a {
			acc += a - b
		}
		acc2 += a * occ.getMultiple(a)
	}
	solution1 = strconv.FormatInt(acc, 10)
	solution2 = strconv.FormatInt(acc2, 10)

	return
}

func (o occurrences) getMultiple(val int64) int64 {
	for _, occ := range o {
		if occ[0] == val {
			return occ[1]
		}
	}
	return 0
}

func countOccurrences(sortedSlice []int64) occurrences {
	res := make(occurrences, len(sortedSlice))
	curr := 0
	res[curr][0] = sortedSlice[0]
	res[curr][1] = 1
	for _, next := range sortedSlice[1:] {
		if next != res[curr][0] {
			curr++
			res[curr][0] = next
		}
		res[curr][1] = res[curr][1] + 1
	}
	return res[:curr+1]
}
