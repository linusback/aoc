package day6

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2024/day6/example"
	inputFile   = "./internal/year2024/day6/input"
)

func Solve() (solution1, solution2 string, err error) {
	log.Println("welcome to day 6 of advent of code")
	b, err := os.ReadFile(exampleFile)
	if err != nil {
		return
	}
	log.Println(string(b))
	return
}
