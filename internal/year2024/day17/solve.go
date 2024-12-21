package day17

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	exampleFile  = "./internal/year2024/day17/example.txt"
	exampleFile2 = "./internal/year2024/day17/example2.txt"
	inputFile    = "./internal/year2024/day17/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

var (
	registries         [3]int64
	originalRegistries [3]int64
	program            []int64
	output             []int64
	ptr                int64
)

func parseRegistries(row []byte, nr int) error {
	for i, b := range row {
		if b == ':' {
			registries[nr] = util.ParseInt64NoError(row[i+1:])
		}
	}
	return nil
}

func parseProgram(row []byte, _ int) error {
	program = util.ParseIntArr[int64](row)
	return nil
}

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseRegistries, parseProgram)
	if err != nil {
		return
	}

	copy(originalRegistries[:], registries[:])

	// part 1
	run()
	solution1 = OutputToString()

	reset()
	log.Println("looking for:", program)
	acc2, ok := runPart2(0, 0)
	log.Println("was ok:", ok, "value:", acc2)
	solution2 = strconv.FormatInt(acc2, 10)
	// 247 839 539 763 386
	return
}

func run() {
	for ptr < int64(len(program)) {
		ptr += instruction()
	}
}

func instruction() int64 {
	switch program[ptr] {
	case 0:
		registries[0] /= int64(math.Exp2(float64(operand())))
	case 1:
		registries[1] ^= program[ptr+1]
	case 2:
		registries[1] = operand() % 8
	case 3:
		if registries[0] != 0 {
			ptr = program[ptr+1]
			return 0
		}
	case 4:
		registries[1] ^= registries[2]
	case 5:
		output = append(output, operand()%8)
	case 6:
		registries[1] = registries[0] / int64(math.Exp2(float64(operand())))
	case 7:
		registries[2] = registries[0] / int64(math.Exp2(float64(operand())))
	}
	return 2
}

func operand() int64 {
	o := program[ptr+1]
	switch o {
	case 0, 1, 2, 3:
		return o
	case 4, 5, 6:
		return registries[o-4]
	default:
		panic(fmt.Sprintf("invalid combo operand %d", o))
	}
}

func runPart2(regA int64, idx int) (res int64, ok bool) {
	if idx > len(program)-1 {
		//log.Printf("reg A: %d, idx: %d -> %s\n", regA, idx, OutputToString())
		//log.Println("quit")
		return regA, true
	}
	regA = regA << 3
	var regANew int64
	for i := int64(0); i < 8; i++ {
		reset()
		regANew = regA | i
		registries[0] = regANew
		run()
		//log.Printf("reg A: %d, idx: %d -> %s\n", regANew, idx, OutputToString())
		if checkOutput(idx) {
			res, ok = runPart2(regANew, idx+1)
			if ok {
				return
			}
		}
	}
	return
}

func checkOutput(idx int) bool {
	return len(output) > idx && program[len(program)-1-idx] == output[len(output)-1-idx]
}

func reset() {
	copy(registries[1:], originalRegistries[1:])
	output = output[:0]
	ptr = 0
}

func OutputToString() string {
	return IntArrToString(output)
}

func IntArrToString(arr []int64) string {
	s := make([]string, len(arr))
	for i, i2 := range arr {
		s[i] = strconv.FormatInt(i2, 10)
	}
	return strings.Join(s, ",")
}
