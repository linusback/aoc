package internal

import (
	"fmt"
	"github.com/linusback/aoc2024/internal/day1"
)

func Solve(day int) (solution1, solution2 string, err error) {
	switch day {
	case 1:
		return day1.Solve()
	default:
		return "", "", fmt.Errorf("day %d not yet created", day)
	}
}
