package day2

import (
	"log"
	"os"
)

const (
	exampleFile = "./internal/year2024/day2/example"
	inputFile   = "./internal/year2024/day2/input"
)

func Solve() (solution1, solution2 string, err error) {
	log.Println("welcome to day 2 of advent of code")
	b, err := os.ReadFile(exampleFile)
	if err != nil {
		return
	}
	log.Println(string(b))
	return
}
