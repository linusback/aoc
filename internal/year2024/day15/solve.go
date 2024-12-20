package day15

import (
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"log"
)

const (
	exampleFile = "./internal/year2024/day15/example.txt"
	inputFile   = "./internal/year2024/day15/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

var wMap util.PositionMap[position.Pos8, position.Pos8, byte]

func solve(filename string) (solution1, solution2 string, err error) {
	wMap, err = util.ToMapOfPositionsByte[position.Pos8](filename, func(row []byte, nr int) error {
		log.Println("handle instructions")
		log.Println(string(row))
		return nil
	})
	if err != nil {
		return
	}
	log.Println(wMap)
	return
}
