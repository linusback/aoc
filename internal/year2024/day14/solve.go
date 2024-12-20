package day14

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"log"
	"strconv"
	"strings"
)

const (
	exampleFile = "./internal/year2024/day14/example.txt"
	inputFile   = "./internal/year2024/day14/input.txt"
	//YMax        = 7
	//XMax        = 11
	YMax = 103
	XMax = 101
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type robot struct {
	position.Pos
	velocity position.Pos
}

func (r *robot) String() string {
	return fmt.Sprintf("{p: %v, v: %v}", r.Pos, r.velocity)
}

func (r *robot) move() {
	r.AddSelf(r.velocity)
	if r.X < 0 {
		r.X += XMax
	}
	if r.X >= XMax {
		r.X -= XMax
	}

	if r.Y < 0 {
		r.Y += YMax
	}
	if r.Y >= YMax {
		r.Y -= YMax
	}
}

var robots []robot

func moveRobots() {
	for i := range robots {
		r := &robots[i]
		r.move()
		robots[i] = *r
	}
}

func solve(filename string) (solution1, solution2 string, err error) {
	var minVal, maxVal int64
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		//fmt.Println("row:", string(row))
		arr := util.ParseIntArr[int64](row)
		//fmt.Println("arr:", arr)
		//for _, i := range arr {
		//	if i < minVal {
		//		minVal = i
		//	}
		//	if i > maxVal {
		//		maxVal = i
		//	}
		//	//if i > 101 {
		//	//	fmt.Println(arr)
		//	//}
		//}
		//fmt.Println(string(row))
		r := robot{
			Pos:      position.NewNegative(arr[1], arr[0]),
			velocity: position.NewNegative(arr[3], arr[2]),
		}
		robots = append(robots, r)
		//fmt.Println("robot:", r)
		return nil
	})
	if err != nil {
		return
	}
	log.Println("min:", minVal, "max:", maxVal)
	//printBathroom()
	var res1, res2 uint64
	for {
		res2++
		moveRobots()
		if res2 == 100 {
			res1 = countQuadrants()
		}
		if foundUniquePos() {
			printBathroom()
			break
		}

	}

	//printBathroom()
	solution1 = strconv.FormatUint(res1, 10)
	solution2 = strconv.FormatUint(res2, 10)
	return
}

func countQuadrants() uint64 {
	var quadrants [4]uint64

	q := 0
	for y := int64(0); y < YMax; y++ {
		if y == YMax/2 {
			q += 2
			continue
		}
		for x := int64(0); x < XMax; x++ {
			if x == XMax/2 {
				q++
				continue
			}
			quadrants[q] += util.CountFunc(robots, robotInTile(y, x))
		}
		q--
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func foundUniquePos() bool {
	for y := int64(0); y < YMax; y++ {
		for x := int64(0); x < XMax; x++ {
			if util.CountFunc(robots, robotInTile(y, x)) > 1 {
				return false
			}
		}
	}
	return true
}

func robotInTile(y, x int64) func(robot) bool {
	return func(r robot) bool {
		return r.Y == y && r.X == x
	}
}

func printBathroom(bots ...robot) {
	var (
		u  uint64
		sb strings.Builder
	)
	if len(bots) == 0 {
		bots = robots
	}
	sb.WriteByte('\n')
	sb.WriteByte('\n')
	for y := int64(0); y < YMax; y++ {
		for x := int64(0); x < XMax; x++ {
			if u = util.CountFunc(bots, func(r robot) bool {
				return r.Y == y && r.X == x
			}); u > 0 {
				sb.Write(strconv.AppendUint(nil, u, 10))
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteByte('\n')
	log.Println(sb.String())
}
