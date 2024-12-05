package year2024

import (
	"github.com/linusback/aoc/internal/year2024/day1"
	"github.com/linusback/aoc/internal/year2024/day2"
	"github.com/linusback/aoc/internal/year2024/day3"
	"github.com/linusback/aoc/internal/year2024/day4"
	"github.com/linusback/aoc/pkg/errorsx"
)

const year = "2024"

func Solve(day string) (solution1, solution2 string, err error) {
	switch day {
	case "1":
		return day1.Solve()
	case "2":
		return day2.Solve()
	case "3":
		return day3.Solve()
	case "4":
		return day4.Solve()
	default:
		err = errorsx.NewSolverError(year, day, errorsx.ErrDayNotCreated)
		return
	}
}
