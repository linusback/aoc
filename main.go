package main

import (
	"errors"
	"fmt"
	"github.com/linusback/aoc2024/internal"
	"github.com/linusback/aoc2024/pkg/aoc"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		day  string
		err  error
		year string
	)
	if len(os.Args) > 2 {
		year = os.Args[1]
		day = os.Args[2]
		err = hasPassed(year, day)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		year, day, err = getYearDay()
		if err != nil {
			log.Fatal(err)
		}
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

func hasPassed(year, day string) (err error) {
	var i int
	_, err = strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("while parsing year string: %v", err)
	}
	i, err = strconv.Atoi(day)
	if err != nil {
		return fmt.Errorf("while parsing day string: %v", err)
	}
	if i < 1 || i > 24 {
		return errors.New("day need to have a value between 1 and 24 inclusive")
	}
	return nil
}

func getYearDay() (year, day string, err error) {
	var loc *time.Location

	loc, err = time.LoadLocation("EST")
	if err != nil {
		return year, day, fmt.Errorf("while parsing location: %v", err)

	}

	current := time.Now()

	start := time.Date(current.Year(), time.November, 30, 0, 0, 0, 0, loc)

	daysDiff := int64(current.Sub(start) / (24 * time.Hour))
	if daysDiff > 24 {
		daysDiff = 24
	}
	return strconv.Itoa(current.Year()), strconv.FormatInt(daysDiff, 10), nil
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
