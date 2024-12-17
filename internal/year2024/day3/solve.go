package day3

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day3/example.txt"
	exampleFile2 = "./internal/year2024/day3/example2"
	inputFile    = "./internal/year2024/day3/input.txt"
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
	// wrong 29018484
	//reg2, err := regexp.Compile(`(?m)(don't\(\).*do\(\))|(don't\(\).*)`)
	//if err != nil {
	//	return
	//}
	//cleaned := reg2.ReplaceAllLiteral(b, nil)

	//fmt.Println("cleaned: ", string(cleaned))

	cleaned, err := cleanBytes(b)
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
	mul = reg.FindAllSubmatch(cleaned, -1)
	var acc2 int64
	for _, m := range mul {
		if len(m) < 3 {
			continue
		}
		acc2 += util.ParseInt64NoError(m[1]) * util.ParseInt64NoError(m[2])
	}

	solution1 = strconv.FormatInt(acc1, 10)
	//fmt.Println(acc2)
	solution2 = strconv.FormatInt(acc2, 10)

	return
}

func cleanBytes(src []byte) (res []byte, err error) {
	startRegex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		return
	}
	stopRegex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		return
	}
	startIdx := startRegex.FindAllIndex(src, -1)
	stopIdx := stopRegex.FindAllIndex(src, -1)

	startIgnore := toIntArray(startIdx, 0)
	stopIgnore := toIntArray(stopIdx, 1)

	// just in case
	slices.Sort(startIgnore)
	slices.Sort(stopIgnore)

	//fmt.Println(startIgnore)
	//fmt.Println(stopIgnore)
	res = make([]byte, 0, len(src))
	start := 0
	stop := startIgnore[0]
	res = append(res, src[start:stop]...)
	startIgnore = startIgnore[1:]

	for {
		start, stopIgnore = getBiggerThen(stopIgnore, stop)
		if start == -1 {
			break
		}
		stop, startIgnore = getBiggerThen(startIgnore, start)
		if stop == -1 {
			//fmt.Println(string(src[start:]))
			res = append(res, src[start:]...)
			break
		}
		//fmt.Println(string(src[start:stop]))
		res = append(res, src[start:stop]...)
	}
	// just in case sort

	return res, nil
}

func toIntArray(src [][]int, idx int) []int {
	res := make([]int, len(src))
	for i, pair := range src {
		res[i] = pair[idx]
	}
	return res
}

func getBiggerThen(res []int, n int) (int, []int) {
	for i, i2 := range res {
		if i2 > n {
			return i2, res[i+1:]
		}
	}
	return -1, nil
}
