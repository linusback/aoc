package day3

import (
	"github.com/linusback/aoc2024/pkg/util"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day3/example"
	inputFile   = "./internal/year2024/day3/input"
)

func Solve() (solution1, solution2 string, err error) {
	log.Println("welcome to day 3 of advent of code")
	b, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}
	reg, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return
	}
	mul := reg.FindAllSubmatch(b, -1)
	var acc1 int64
	for _, m := range mul {
		if len(m) < 3 {
			continue
		}
		acc1 += util.ParseInt64NoError(m[1]) * util.ParseInt64NoError(m[2])
	}
	solution1 = strconv.FormatInt(acc1, 10)

	return
}
