package day8

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"maps"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day8/example"
	inputFile   = "./internal/year2024/day8/input"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
	var (
		antennaMap = make(map[byte][]position.Pos8)
		yMax, xMax uint8
		posMax     position.Pos8
	)
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		xMax = uint8(len(row) - 1)
		yMax = uint8(nr)
		for x, b := range row {
			if util.IsAlphaNumerical[b] == 0 {
				continue
			}
			antennaMap[b] = append(antennaMap[b], position.New8(yMax, uint8(x)))
		}
		return nil
	})
	if err != nil {
		return
	}
	posMax = position.New8(yMax, xMax)

	antiNodes1 := make(map[position.Pos8]struct{}, yMax*xMax)
	antiNodes2 := make(map[position.Pos8]struct{}, yMax*xMax)
	for k, v := range antennaMap {
		maps.Insert(antiNodes2, util.ToKeysSeq2(v, struct{}{}))
		calculateAntiNodes(k, v, posMax, antiNodes1, antiNodes2)
	}
	fmt.Println(len(antiNodes1))
	fmt.Println(len(antiNodes2))

	solution1 = strconv.Itoa(len(antiNodes1))
	solution2 = strconv.Itoa(len(antiNodes2))
	return
}

func calculateAntiNodes(antenna byte, v []position.Pos8, posMax position.Pos8, m, m2 map[position.Pos8]struct{}) {
	var antiNode, diff position.Pos8
	//fmt.Printf("checking %v\n", v)
	for i, p1 := range v[:len(v)-1] {
		for _, p2 := range v[i+1:] {
			diff = position.Sub8(p1, p2)

			antiNode = position.Add8(p1, diff)
			if antiNode.IsInside(posMax) {
				m[antiNode] = struct{}{}
			}
			for antiNode.IsInside(posMax) {
				m2[antiNode] = struct{}{}
				antiNode.Add(diff)
			}

			antiNode = position.Sub8(p2, diff)
			if antiNode.IsInside(posMax) {
				m[antiNode] = struct{}{}
			}

			for antiNode.IsInside(posMax) {
				m2[antiNode] = struct{}{}
				antiNode.Sub(diff)
			}
		}
	}
}
