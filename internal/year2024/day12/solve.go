package day12

import (
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"log"
	"math"
)

const (
	exampleFile  = "./internal/year2024/day12/example.txt"
	exampleFile2 = "./internal/year2024/day12/example2.txt"
	exampleFile3 = "./internal/year2024/day12/example3.txt"
	inputFile    = "./internal/year2024/day12/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile3)
}

var directions = position.DirectionsPos8[:position.Dir_UpRight]

func solve(filename string) (solution1, solution2 string, err error) {
	var (
		gardenMap = make([]byte, math.MaxUint16)
		maxPos    position.Pos8
		y, x      uint8
	)
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			x = uint8(len(row) - 1)
		}
		y = uint8(nr)
		for i, b := range row {
			gardenMap[position.New8(y, uint8(i))] = b
		}
		return nil
	})
	if err != nil {
		return
	}

	maxPos = position.New8(y, x)
	//log.Println(gardenMap)
	log.Println(maxPos)
	log.Println(gardenMap[:maxPos+1])
	return
}
