package main

import (
	"fmt"
	"github.com/linusback/aoc2024/internal"
	"github.com/linusback/aoc2024/pkg/aoc"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Need to specify day to solve")
	}
	day := os.Args[1]
	i, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalln("parse int ", err)
	}

	var solution1, solution2 string
	solution1, solution2, err = internal.Solve(i)
	if err != nil {
		log.Fatalf("error from from solver %d: %v", i, err)
	}
	err = send(aoc.Part1, day, solution1)
	if err != nil {
		log.Println(err)
	}
	err = send(aoc.Part2, day, solution2)
	if err != nil {
		log.Println(err)
	}
}

func send(part aoc.Part, day, solution string) error {
	if len(solution) == 0 {
		return nil
	}
	err := aoc.Send(part, day, solution)
	if err != nil {
		return fmt.Errorf("error while sending solution \"%s\" to probblem %s of day %s: %v", solution, part, day, err)
	}
	return nil
}
