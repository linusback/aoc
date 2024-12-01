package year2024

import (
	"fmt"
	"github.com/linusback/aoc2024/internal/year2024/day1"
)

func Solve(day string) (solution1, solution2 string, err error) {
	switch day {
	case "1":
		return day1.Solve()
	default:
		return "", "", fmt.Errorf("day %s not yet created", day)
	}
}
