package day15

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day15/example.txt"
	exampleFile2 = "./internal/year2024/day15/example2.txt"
	inputFile    = "./internal/year2024/day15/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type warehouseType struct {
	util.PositionMap[position.Pos8, position.Pos8, byte]
	map2         []byte
	robot        position.Pos8
	robotStart   position.Pos8
	instructions []position.Pos8
}

type dirKey struct {
	pos, dir position.Pos8
}

var (
	warehouse      warehouseType
	instructionSet = [255]position.Pos8{
		'^': position.New8Negative(-1, 0),
		'>': position.New8Negative(0, 1),
		'v': position.New8Negative(1, 0),
		'<': position.New8Negative(0, -1),
	}
	impossibleMoves = make(map[dirKey]struct{})
)

func solve(filename string) (solution1, solution2 string, err error) {
	warehouse.PositionMap, err = util.ToMapOfPositionsByte[position.Pos8](filename, func(row []byte, nr int) error {
		for _, b := range row {
			warehouse.instructions = append(warehouse.instructions, instructionSet[b])
		}
		return nil
	})
	if err != nil {
		return
	}
	for _, pos := range warehouse.Positions {
		if warehouse.Map[pos] == '@' {
			warehouse.robot = pos
			warehouse.robotStart = pos
		}
	}
	//log.Println(warehouse.MapString())
	for _, ins := range warehouse.instructions {
		move(ins)
		//log.Println(warehouse.MapString())
	}
	//log.Println(warehouse.MapString())
	var (
		acc1 uint64
		pos  position.Pos8
	)
	for i, b := range warehouse.Map {
		if b != 'O' {
			continue
		}
		pos = position.Pos8(i)
		acc1 += 100 * uint64(pos.Y())
		acc1 += uint64(pos.X())
	}
	fmt.Println("result is ", acc1)
	solution1 = strconv.FormatUint(acc1, 10)
	return
}

func move(ins position.Pos8) {
	if _, ok := impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}]; ok {
		return
	}
	newPos := warehouse.robot.Add(ins)
	newTile := warehouse.Map[newPos]
	if newTile == '#' {
		impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}] = struct{}{}
		return
	}
	if newTile == '.' {
		warehouse.Map[warehouse.robot] = '.'
		warehouse.robot = newPos
		warehouse.Map[newPos] = '@'
		return
	}
	nextPos := newPos
	for newTile == 'O' {
		nextPos = nextPos.Add(ins)
		newTile = warehouse.Map[nextPos]
		if newTile == '#' {
			//impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}] = struct{}{}
			return
		}
	}
	warehouse.Map[warehouse.robot] = '.'
	warehouse.Map[nextPos] = 'O'
	warehouse.robot = newPos
	warehouse.Map[newPos] = '@'
	return

	return
}
