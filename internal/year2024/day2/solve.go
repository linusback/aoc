package day2

import (
	"github.com/linusback/aoc2024/pkg/util"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day2/example"
	inputFile   = "./internal/year2024/day2/input"
)

func Solve() (solution1, solution2 string, err error) {
	var (
		reports  [][]int64
		reports2 [][]int64
		report   []int64
	)
	err = util.DoEachRowFile(exampleFile, func(row []byte, nr int) error {
		report = util.ParseInt64ArrNoError(row)
		reports = appendIfSafe(reports, report, isSafe)
		reports2 = appendIfSafe(reports2, report, isSafe2)
		return nil
	})
	if err != nil {
		return
	}

	solution1 = strconv.Itoa(len(reports))
	return
}

func isSafe2(report []int64) bool {
	var isIncrease bool
	diff := report[1] - report[0]
	if diff == 0 {
		return false
	}
	isIncrease = diff > 0
	if !check(diff, isIncrease) {
		return false
	}
	for i, val := range report[2:] {
		diff = val - report[i+1]
		if !check(diff, isIncrease) {
			return false
		}
	}
	return true
}

func isSafe(report []int64) bool {
	var isIncrease bool
	diff := report[1] - report[0]
	if diff == 0 {
		return false
	}
	isIncrease = diff > 0
	if !check(diff, isIncrease) {
		return false
	}
	for i, val := range report[2:] {
		diff = val - report[i+1]
		if !check(diff, isIncrease) {
			return false
		}
	}
	return true
}

func check(diff int64, isIncrease bool) bool {
	if !isIncrease {
		diff *= -1
	}
	return checkIncrease(diff)
}

func checkIncrease(diff int64) bool {
	return 1 <= diff && diff <= 3
}

func appendIfSafe(reports [][]int64, report []int64, safeFunc func(report []int64) bool) [][]int64 {
	if len(report) > 0 && safeFunc(report) {
		return append(reports, report)
	}
	return reports
}
