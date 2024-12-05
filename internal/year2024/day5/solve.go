package day5

import (
	"bytes"
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
)

const (
	exampleFile = "./internal/year2024/day5/example"
	inputFile   = "./internal/year2024/day5/input"
)

func Solve() (solution1, solution2 string, err error) {
	var (
		inUpdates bool
		rules     [][2]int64
	)

	err = util.DoEachRowFile(exampleFile, func(row []byte, nr int) error {
		row = bytes.TrimSpace(row)
		if len(row) == 0 {
			log.Println("len is 0")
			if inUpdates {
				log.Println("do nothing since in updates")
				return nil
			}
			log.Println("setting updates to true")
			inUpdates = true
			return nil
		}
		if inUpdates {
			log.Printf("u: %s\n", row)
			return nil
		}

		res := util.ParseInt64ArrNoError(row)
		if len(res) != 2 {
			return fmt.Errorf("expected len 2 got %d", len(res))
		}
		rules = append(rules, *(*[2]int64)(res))
		return nil
	})
	//b, err := os.ReadFile(exampleFile)
	if err != nil {
		return
	}
	log.Println("rules: ", rules)
	return
}

func generateCompare(rules [][2]int64) func(a, b []int64) int {
	return func(a, b []int64) int {
		return 0
	}
}
