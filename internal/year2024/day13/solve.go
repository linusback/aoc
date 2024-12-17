package day13

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2024/day13/example"
	inputFile   = "./internal/year2024/day13/input"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
    log.Println("welcome to day 13 of advent of code")
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	log.Println(string(b))
	return
}
