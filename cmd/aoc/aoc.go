package main

import (
	"fmt"
	"github.com/linusback/aoc2024/internal"
	"github.com/linusback/aoc2024/pkg/aoc"
	"github.com/linusback/aoc2024/pkg/util"
	"log"
	"os"
)

func main() {
	var (
		day  string
		err  error
		year string
	)

	year, day, err = util.GetYearDay(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	err = aoc.Download(year, day)
	if err != nil {
		log.Fatal(err)
	}

	var solution1, solution2 string
	solution1, solution2, err = internal.Solve(year, day)
	if err != nil {
		log.Fatalln(err)
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
