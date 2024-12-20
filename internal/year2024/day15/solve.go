package day15

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"log"
	"slices"
	"strconv"
	"strings"
)

const (
	exampleFile  = "./internal/year2024/day15/example.txt"
	exampleFile2 = "./internal/year2024/day15/example2.txt"
	exampleFile3 = "./internal/year2024/day15/example3.txt"
	exampleFile4 = "./internal/year2024/day15/example4.txt"
	inputFile    = "./internal/year2024/day15/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type warehouseType struct {
	util.PositionMap[position.Pos8, position.Pos8, byte]
	map2         []byte
	robot        position.Pos8
	instructions []position.Pos8
}

type tile struct {
	pos position.Pos8
	b   byte
}

func (t tile) String() string {
	return fmt.Sprintf("%v: '%c'", t.pos, t.b)
}

type dirKey struct {
	pos, dir position.Pos8
}

const (
	up    = position.Pos8(uint8(255)) << 8
	right = position.Pos8(1)
	down  = position.Pos8(uint8(1)) << 8
	left  = position.Pos8(255)
)

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
		}
	}
	warehouse.map2 = make([]byte, len(warehouse.Map))
	copy(warehouse.map2, warehouse.Map)

	for _, ins := range warehouse.instructions {
		move(ins)
	}

	var (
		acc1, acc2 uint64
		pos        position.Pos8
	)
	for i, b := range warehouse.Map {
		if b != 'O' {
			continue
		}
		pos = position.Pos8(i)
		acc1 += 100 * uint64(pos.Y())
		acc1 += uint64(pos.X())
	}
	populateMap2()
	//log.Println(MapString())
	for _, ins := range warehouse.instructions {
		if !move2(ins) {
			//log.Println(moveString(ins), "impossible no change")
			continue
		}
		//log.Println(moveString(ins))
		//log.Println(MapString())
	}
	//log.Println(MapString())
	for i, b := range warehouse.Map {
		if b != '[' {
			continue
		}
		pos = position.Pos8(i)
		acc2 += 100 * uint64(pos.Y())
		acc2 += uint64(pos.X())
	}

	log.Println("target is:", 1492011)
	log.Println(acc2)
	//log.Println(warehouse.MaxPos)

	solution1 = strconv.FormatUint(acc1, 10)
	//solution2 = strconv.FormatUint(acc2, 10)
	return
}

func moveString(dir position.Pos8) string {
	switch dir {
	case up:
		return "Moving ^ (up)"
	case right:
		return "Moving > (right)"
	case down:
		return "Moving v (down)"
	case left:
		return "Moving < (left)"
	default:
		panic("unsupported direction")
	}
}

func populateMap2() {
	var lastIndex position.Pos8
	warehouse.Map = warehouse.map2
	warehouse.map2 = make([]byte, 0, len(warehouse.Map)*2)
	for _, b := range warehouse.Map {
		switch b {
		case 'O':
			warehouse.map2 = append(warehouse.map2, '[', ']')
		case '@':
			warehouse.map2 = append(warehouse.map2, '@', '.')
			warehouse.robot = position.Pos8(len(warehouse.map2) - 2)
		case 0:
			if lastIndex == 0 {
				lastIndex = position.Pos8(len(warehouse.map2) - 1)
				var sb strings.Builder
				for x := uint8(0); x <= lastIndex.X(); x++ {
					sb.WriteByte(warehouse.map2[position.New8(lastIndex.Y(), x)])
				}
				warehouse.map2 = util.AppendRepeat(warehouse.map2, 255-lastIndex.X(), 0)
				//log.Println("written:", sb.String())
				//log.Println("last index: ", lastIndex, warehouse.map2[lastIndex]+1)
			}
			continue
		default:
			warehouse.map2 = append(warehouse.map2, b, b)
		}
		lastIndex = 0
	}
	warehouse.MaxPos = position.Pos8(len(warehouse.map2) - 1)
	warehouse.Map = warehouse.map2
	clear(impossibleMoves)
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

}

func move2(ins position.Pos8) bool {
	if _, ok := impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}]; ok {
		return false
	}
	newPos := warehouse.robot.Add(ins)
	newTile := warehouse.Map[newPos]
	if newTile == '#' {
		impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}] = struct{}{}
		return false
	}
	if newTile == '.' {
		warehouse.Map[warehouse.robot] = '.'
		warehouse.robot = newPos
		warehouse.Map[newPos] = '@'
		return true
	}
	newPositions := make(map[position.Pos8]byte, 10)
	if ins == up || ins == down {
		//log.Println("first tile:", tile{pos: warehouse.robot, b: '@'})
		var (
			tiles []tile
			done  bool
		)
		tiles, newPositions, done = getBoxedTiles(ins, newPositions, tile{pos: warehouse.robot, b: '@'})
		//log.Println("initial tiles", tiles)
		for {
			if done {
				break
			}
			tiles, newPositions, done = getBoxedTiles(ins, newPositions, tiles...)
			//log.Println("tiles: ", tiles)
			if slices.ContainsFunc(tiles, func(t tile) bool {
				return t.b == '#'
			}) {
				return false
			}

		}
	} else {
		newPositions[warehouse.robot] = '.'
		newPositions[newPos] = '@'
		nextPos := newPos
		for newTile == '[' || newTile == ']' {
			nextPos = nextPos.Add(ins)
			newPositions[nextPos] = newTile
			newTile = warehouse.Map[nextPos]
			if newTile == '#' {
				//impossibleMoves[dirKey{pos: warehouse.robot, dir: ins}] = struct{}{}
				return false
			}
		}
	}
	//log.Println("change")
	warehouse.robot = newPos
	for pos, b := range newPositions {
		warehouse.Map[pos] = b
		//log.Println("set", pos, "to", string(b))
	}
	return true
}

func getBoxedTiles(dir position.Pos8, newPositions map[position.Pos8]byte, positions ...tile) (newTiles []tile, nPos map[position.Pos8]byte, done bool) {
	if len(positions) == 0 {
		return nil, newPositions, true
	}
	done = true
	newTiles = make([]tile, 0, len(positions)+2)
	for i, t := range positions {
		p := t.pos.Add(dir)
		val := warehouse.Map[p]
		if done && val != '.' {
			done = false
		}
		if val == '.' {
			continue
		}
		if i == 0 && val == ']' {
			newTiles = append(newTiles, tile{
				pos: p - 1,
				b:   warehouse.Map[p-1],
			})
		}
		newTiles = append(newTiles, tile{
			pos: p,
			b:   val,
		})
		if i == len(positions)-1 && val == '[' {
			newTiles = append(newTiles, tile{
				pos: p + 1,
				b:   warehouse.Map[p+1],
			})
		}
	}
	for _, t := range positions {
		if _, ok := newPositions[t.pos]; !ok {
			newPositions[t.pos] = '.'
		}
		newPositions[t.pos.Add(dir)] = t.b
	}
	return newTiles, newPositions, done
}

func MapString() string {
	var (
		sb       strings.Builder
		lastByte byte
	)
	sb.WriteString("\n\tMap: ")
	for _, b := range warehouse.Map {
		if b == 0 && lastByte != 0 {
			sb.WriteString("\n\t     ")
		} else {

		}
		sb.WriteByte(b)
		lastByte = b
	}
	return sb.String()
}
