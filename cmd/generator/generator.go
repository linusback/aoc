package main

import (
	"github.com/linusback/aoc2024/pkg/aoc"
	"github.com/linusback/aoc2024/pkg/generate"
	"github.com/linusback/aoc2024/pkg/util"
	"log"
	"os"
	"slices"
)

func main() {
	year, days, err := util.GetYearDays(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	slices.Reverse(days)
	err = generate.Generate(year, days)
	if err != nil {
		log.Fatal(err)
	}

	err = aoc.Download(year, days)
	if err != nil {
		log.Fatal(err)
	}
}
