package day2

import (
	"github.com/linusback/aoc/pkg/util"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day2/example"
	inputFile   = "./internal/year2024/day2/input"
)

func Solve() (solution1, solution2 string, err error) {
	var (
		//reports  [][]int64
		//reports2 [][]int64
		report       []int64
		safe1, safe2 int
	)

	err = util.DoEachRowFile(inputFile, func(row []byte, nr int) error {
		report = util.ParseInt64ArrNoErrorCache(row, report)
		if len(report) == 0 {
			return nil
		}
		if isSafe(report) {
			safe1++
			safe2++
			return nil
		}

		if isSafe2(report) {
			safe2++
		}
		//reports = appendIfSafe(reports, report, isSafe)
		//reports2 = appendIfSafe(reports2, report, isSafe2)

		return nil
	})
	if err != nil {
		return
	}

	solution1 = strconv.Itoa(safe1)
	//fmt.Println(reports2)
	solution2 = strconv.Itoa(safe2)
	return
}

func isSafe2(report []int64) bool {
	// remove first
	if isSafe(report[1:]) {
		//fmt.Println("  safe if remove first")
		return true
	}

	// try last
	if isSafe(report[:len(report)-1]) {
		//fmt.Println("  safe if remove last")
		return true
	}

	alt := make([]int64, len(report)-1)
	for remove := 1; remove < len(alt); remove++ {
		copyExcept(alt, report, remove)
		if isSafe(alt) {
			//fmt.Printf("  safe if remove %d: %v\n", report[remove], alt)
			return true
		}
	}

	return false
}

func copyExcept(dst, src []int64, except int) {
	n := copy(dst, src[:except])
	copy(dst[n:], src[except+1:])
}

func isSafe(report []int64) bool {
	diff := report[1] - report[0]
	if diff == 0 {
		return false
	}
	var (
		lower int64 = 1
		upper int64 = 3
	)
	if diff < 0 {
		lower = -3
		upper = -1
	}
	for i, val := range report[1:] {
		diff = val - report[i]
		if lower > diff || diff > upper {
			return false
		}
	}
	return true
}

func appendIfSafe(reports [][]int64, report []int64, safeFunc func(report []int64) bool) [][]int64 {
	if len(report) > 0 && safeFunc(report) {
		return append(reports, report)
	}
	return reports
}
