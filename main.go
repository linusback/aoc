package main

import (
	"github.com/linusback/aoc2024/internal/day1"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Need to specify day to solve")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("parse int ", err)
	}
	switch i {
	case 1:
		day1.Solve()
	default:
		log.Fatalf("day %d not ye created", i)
	}
}
