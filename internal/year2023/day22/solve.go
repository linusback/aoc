package day22

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2023/day22/example.txt"
	inputFile   = "./internal/year2023/day22/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(exampleFile)
}

func solve(filename string) (solution1, solution2 string, err error) {
	log.Println("welcome to day 22 of advent of code")
	b, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	log.Printf("\n%s\n", b)
	return
}
