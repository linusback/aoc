package day12

import (
	"github.com/linusback/aoc/pkg/util"
	"github.com/linusback/aoc/pkg/util/position"
	"log"
	"slices"
	"strconv"
)

const (
	exampleFile  = "./internal/year2024/day12/example.txt"
	exampleFile2 = "./internal/year2024/day12/example2.txt"
	exampleFile3 = "./internal/year2024/day12/example3.txt"
	exampleFile4 = "./internal/year2024/day12/example4.txt"
	exampleFile5 = "./internal/year2024/day12/example5.txt"
	inputFile    = "./internal/year2024/day12/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

var (
	directions         = position.DirectionsPos8
	diagonalDirections = position.DirectionsDiagonalPos8
)

type gardenRegion struct {
	name      byte
	plots     []position.Pos8
	parameter int
	sides     int
}

type gardenMapType struct {
	util.PositionMap[position.Pos8, position.Pos8, byte]
	visited []position.Pos8
	regions []*gardenRegion
}

var gardenMap gardenMapType

func solve(filename string) (solution1, solution2 string, err error) {

	posMap, err := util.ToMapOfPositionsByte[position.Pos8](filename)
	if err != nil {
		return
	}
	gardenMap.PositionMap = posMap
	gardenMap.visited = make([]position.Pos8, 0, len(gardenMap.Positions))
	gardenMap.regions = make([]*gardenRegion, 0, len(gardenMap.Positions)/4)
	for _, pos := range gardenMap.Positions {
		if slices.Contains(gardenMap.visited, pos) {
			continue
		}
		region := new(gardenRegion)
		region.name = gardenMap.Map[pos]
		findRegion(pos, region)
		gardenMap.regions = append(gardenMap.regions, region)
	}
	var total1, total2 int
	for _, v := range gardenMap.regions {
		//log.Println(string(k))
		total1 += len(v.plots) * v.parameter
		total2 += len(v.plots) * v.sides

		//log.Printf("PART 1 %c with a price of: %d + %d = %d \n", v.name, len(v.plots), v.parameter, len(v.plots) * v.parameter)
		//log.Printf("PART 2 %c with a price of: %d + %d = %d \n", v.name, len(v.plots), v.sides, len(v.plots)*v.sides)
		//log.Println(v.plots)
		//log.Println("")
	}
	solution1 = strconv.Itoa(total1)
	solution2 = strconv.Itoa(total2)
	log.Println("total1 price: ", total1)
	log.Println("total2 price: ", total2)
	return
}

func findRegion(pos position.Pos8, region *gardenRegion) {
	gardenMap.visited = append(gardenMap.visited, pos)
	region.plots = append(region.plots, pos)
	plots, parameter, sides := getNearby(region.name, pos)
	region.parameter += parameter
	region.sides += sides
	for _, p := range plots {
		if slices.Contains(gardenMap.visited, p) {
			continue
		}
		findRegion(p, region)
	}
}

func getNearby(name byte, pos position.Pos8) (plots []position.Pos8, parameterToAdd, sides int) {
	// count corners instead of sides, should be the same and is easier to think about.
	var (
		candidate                  position.Pos8
		sameRegion, isInside       bool
		prevInside, prevSameRegion bool
	)
	plots = make([]position.Pos8, 0, 4)
	preCandidate := pos.Add(directions[3])

	prevInside = gardenMap.HasInside(preCandidate)
	if prevInside {
		prevSameRegion = name == gardenMap.Map[preCandidate]
	}
	for i, d := range directions {
		candidate = pos.Add(d)
		isInside = gardenMap.HasInside(candidate)
		if isInside {
			sameRegion = name == gardenMap.Map[candidate]
		} else {
			sameRegion = false
		}

		if isInside && sameRegion {
			plots = append(plots, candidate)
		} else {
			parameterToAdd++
		}

		// orthogonal normal corner
		if !sameRegion && !prevSameRegion {
			sides++
		}
		// diagonal inverted corner
		if sameRegion && prevSameRegion && name != gardenMap.Map[pos.Add(diagonalDirections[i])] {
			sides++
		}
		prevInside = isInside
		prevSameRegion = sameRegion
	}
	return plots, parameterToAdd, sides
}
