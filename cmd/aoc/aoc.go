package main

import (
	"fmt"
	"github.com/linusback/aoc2024/internal"
	"github.com/linusback/aoc2024/pkg/aoc"
	"github.com/linusback/aoc2024/pkg/util"
	"log"
	"os"
	"time"
)

func main() {
	var (
		days []string
		err  error
		year string
	)

	year, days, err = util.GetYearDays(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if len(days) == 0 {
		log.Fatal("no days selected")
	}

	err = aoc.Download(year, days)
	if err != nil {
		log.Fatal(err)
	}

	var solution1, solution2 string
	start := time.Now()
	solution1, solution2, err = internal.Solve(year, days[0])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("result 1: ", solution1)
	log.Println("result 2: ", solution2)
	log.Println("Time elapsed:", time.Since(start))

	err = send(aoc.Part1, days[0], solution1)
	if err != nil {
		log.Println(err)
	}
	err = send(aoc.Part2, days[0], solution2)
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
