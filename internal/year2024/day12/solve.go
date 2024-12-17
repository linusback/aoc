package day12

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2024/day12/example.txt"
	inputFile   = "./internal/year2024/day12/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
    log.Println("welcome to day 12 of advent of code")
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	log.Println(string(b))
	return
}
