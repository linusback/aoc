package day5

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"slices"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day5/example"
	inputFile   = "./internal/year2024/day5/input"
)

func Solve() (solution1, solution2 string, err error) {
	var (
		rules   [][2]int64
		updates [][]int64
	)

	err = util.DoEachRowFile(inputFile,
		func(row []byte, nr int) error {
			res := util.ParseInt64ArrNoError(row)
			if len(res) != 2 {
				return fmt.Errorf("expected len 2 got %d", len(res))
			}
			rules = append(rules, *(*[2]int64)(res))
			return nil
		}, func(row []byte, nr int) error {
			updates = append(updates, util.ParseInt64ArrNoError(row))
			return nil
		})
	if err != nil {
		return
	}

	//log.Println("rules: ", rules)
	//log.Println("pages: ", updates)

	ruleList := newRuleList(rules)
	//log.Println(ruleList)
	cmp := getComparer(rules)
	var acc1, acc2 int64
	for _, u := range updates {
		if isOrdered(u, ruleList) {
			acc1 += u[len(u)/2]
			//continue
		}
		if slices.IsSortedFunc(u, cmp) {
			acc2 += u[len(u)/2]
		}

		//log.Printf("%v is ordered: %v\n", u, isOrdered(u, ruleList))
	}
	log.Println(acc2)
	solution1 = strconv.FormatInt(acc1, 10)
	return
}

func isOrdered(update []int64, rl []rulelist) bool {
	var forbidden []int64
	for _, p := range update {
		if slices.Contains(forbidden, p) {
			return false
		}
		if idx := slices.IndexFunc(rl, containsY(p)); idx > -1 {
			forbidden = util.AppendUnique(forbidden, rl[idx].forbidden...)
		}
	}
	return true
}

func getComparer(rules [][2]int64) func(int64, int64) int {
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

//type comparer struct {
//	rl      []rulelist
//	applied []rulelist
//}
//
//func (c *comparer) reset() {
//	c.applied = c.applied[:0]
//}
//
//func (c *comparer) isSortedCmp() func(a, b int64) int {
//	var applied []rulelist
//	return func(a, b int64) int {
//		i
//		if idx := slices.IndexFunc(rulelist, containsY(p)); idx > -1 {
//			forbidden = util.AppendUnique(forbidden, rulelist[idx].forbidden...)
//		}
//	}
//
//}

type rulelist struct {
	y         int64
	forbidden []int64
	idx       int
}

//func (rl *rulelist) CloneWithIds(idx int) rulelist {
//	return rulelist{
//		y:         rl.y,
//		forbidden: rl.forbidden,
//		idx:       idx,
//	}
//}

func newRuleList(rules [][2]int64) (res []rulelist) {
	res = make([]rulelist, 0, len(rules))
	for _, r := range rules {
		if idx := slices.IndexFunc(res, containsY(r[1])); idx != -1 {
			res[idx].forbidden = append(res[idx].forbidden, r[0])
		} else {
			res = append(res, rulelist{
				y:         r[1],
				forbidden: []int64{r[0]},
			})
		}
	}
	return
}

func containsY(y int64) func(rulelist) bool {
	return func(r rulelist) bool {
		return y == r.y
	}
}
