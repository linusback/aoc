package day4

import (
	"bytes"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"os"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day4/example.txt"
	inputFile   = "./internal/year2024/day4/input.txt"
	mas         = "MAS"
	amx         = "SAM"
)

func Solve() (solution1, solution2 string, err error) {
	log.Println("welcome to day 4 of advent of code")
	b, err := os.ReadFile(inputFile)
	//b, err := os.ReadFile(exampleFile)
	if err != nil {
		return
	}
	matrix := bytes.Split(b, []byte{'\n'})
	if len(matrix[len(matrix)-1]) == 0 {
		matrix = matrix[:len(matrix)-1]
	}
	acc1 := 0
	acc2 := 0
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1
	for y, row := range matrix {
		//log.Printf("checking row %s\n", row)
		for x, c := range row {
			acc1 += searchRowPart1(c, x, y, row, matrix)
			acc2 += searchRowPart2(c, x, y, maxX, maxY, matrix)
		}
	}
	//log.Println("acc2:", acc2)
	solution1 = strconv.Itoa(acc1)
	solution2 = strconv.Itoa(acc2)
	return
}

func searchRowPart2(c byte, x, y, xMax, yMax int, m [][]byte) (res int) {
	if c != 'A' || x == 0 || y == 0 || x == xMax || y == yMax {
		return 0
	}
	if ((m[y-1][x-1] == 'M' && m[y+1][x+1] == 'S') ||
		(m[y-1][x-1] == 'S' && m[y+1][x+1] == 'M')) && ((m[y-1][x+1] == 'M' && m[y+1][x-1] == 'S') ||
		(m[y-1][x+1] == 'S' && m[y+1][x-1] == 'M')) {
		return 1
	}
	return 0
}

func searchRowPart1(c byte, x, y int, row []byte, m [][]byte) (res int) {
	if c != 'X' {
		return 0
	}

	xMax := len(row)
	yMax := len(m)
	// horizontal
	checkBackward := x-3 >= 0
	checkForward := x+4 <= xMax
	checkUp := y-3 >= 0
	checkDown := y+4 <= yMax

	if checkBackward && util.ToUnsafeString(row[x-3:x]) == amx {
		//log.Printf("found backwards at x: %d, y: %d", x, y)
		res++
	}
	if checkForward && util.ToUnsafeString(row[x+1:x+4]) == mas {
		//log.Printf("found forwards at x: %d, y: %d", x, y)
		res++
	}

	//vertical
	if checkUp && search(x, y, 0, -1, mas, m) {
		//log.Printf("found up at x: %d, y: %d", x, y)
		res++
	}
	if checkDown && search(x, y, 0, 1, mas, m) {
		//log.Printf("found down at x: %d, y: %d", x, y)
		res++
	}

	// diagonally
	if checkBackward && checkUp && search(x, y, -1, -1, mas, m) {
		//log.Printf("found backward-up at x: %d, y: %d", x, y)
		res++
	}
	if checkBackward && checkDown && search(x, y, -1, 1, mas, m) {
		//log.Printf("found backward-down at x: %d, y: %d", x, y)
		res++
	}
	if checkForward && checkUp && search(x, y, 1, -1, mas, m) {
		//log.Printf("found forward-up at x: %d, y: %d", x, y)
		res++
	}
	if checkForward && checkDown && search(x, y, 1, 1, mas, m) {
		//log.Printf("found forward-down at x: %d, y: %d", x, y)
		res++
	}

	return res
}

func search(x, y, dX, dY int, target string, m [][]byte) bool {
	x += dX
	y += dY
	for i := 0; i < len(target); i++ {
		if m[y][x] != target[i] {
			return false
		}
		x += dX
		y += dY
	}
	return true
}
