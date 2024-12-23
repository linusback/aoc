package year2023

import (
	"github.com/linusback/aoc/internal/year2023/day22"
	"github.com/linusback/aoc/pkg/errorsx"
)

const year = "2023"

func Solve(day string) (solution1, solution2 string, err error) {
	switch day {
	case "22":
		return day22.Solve()
	default:
		err = errorsx.NewSolverError(year, day, errorsx.ErrDayNotCreated)
		return
	}
}
