package day14

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2024/day14/example.txt"
	inputFile   = "./internal/year2024/day14/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
    log.Println("welcome to day 14 of advent of code")
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	log.Println(string(b))
	return
}
