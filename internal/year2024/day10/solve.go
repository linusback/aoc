package day10

import (
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"math"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day10/example.txt"
	example2File = "./internal/year2024/day10/example2.txt"
	inputFile    = "./internal/year2024/day10/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type topographicMap struct {
	MaxPos position.Pos8
	Map    []int8
}

var tMap topographicMap

func solve(filename string) (solution1, solution2 string, err error) {
	const outside int8 = -1
	var (
		y, x   uint8
		starts []position.Pos8
	)

	tMap.Map = util.Repeat(math.MaxUint16, outside)
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			x = uint8(len(row) - 1)
		}
		y = uint8(nr)
		for i := uint8(0); i <= x; i++ {
			p := position.New8(y, i)
			tMap.Map[p] = int8(row[i] - '0')
			if row[i] == '0' {
				starts = append(starts, p)
			}
		}
		return nil
	})
	if err != nil {
		return
	}
	tMap.MaxPos = position.New8(y, x)
	//log.Println(tMap.Map)
	tMap.Map = tMap.Map[:tMap.MaxPos+1]
	//log.Println(tMap.Map)
	//log.Println(starts)

	part1, part2 := findPaths(starts)
	solution1 = strconv.FormatInt(part1, 10)
	solution2 = strconv.FormatInt(part2, 10)
	return
}

func findPaths(starts []position.Pos8) (part1, part2 int64) {
	var trailheads []position.Pos8
	for _, p := range starts {
		trailheads = findPath(p, 0, trailheads[:0])

		//log.Printf("pos: %v, has score %d: %v", p, len(trailheads), trailheads)
		part1 += int64(util.LenUnique(trailheads))
		part2 += int64(len(trailheads))
	}
	return part1, part2
}

func findPath(pos position.Pos8, val int8, ends []position.Pos8) []position.Pos8 {
	if val == 9 {
		return append(ends, pos)
	}
	for _, dir := range position.DirectionsPos8[:position.Dir_UpRight] {
		next := pos.Add(dir)
		if next.IsInside(tMap.MaxPos) && tMap.Map[next] == val+1 {
			ends = findPath(next, val+1, ends)
		}
	}
	return ends
}
