package internal

import (
	"fmt"
	"github.com/linusback/aoc2024/internal/year2024"
)

func Solve(year, day string) (solution1, solution2 string, err error) {
	switch year {
	case "2024":
		return year2024.Solve(day)
	default:
		return "", "", fmt.Errorf("year %s not yet created", year)
	}
}
