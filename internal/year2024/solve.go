package year2024

import (
	"github.com/linusback/aoc/internal/year2024/day1"
	"github.com/linusback/aoc/internal/year2024/day2"
	"github.com/linusback/aoc/internal/year2024/day3"
	"github.com/linusback/aoc/internal/year2024/day4"
	"github.com/linusback/aoc/internal/year2024/day5"
	"github.com/linusback/aoc/internal/year2024/day6"
	"github.com/linusback/aoc/internal/year2024/day7"
	"github.com/linusback/aoc/internal/year2024/day8"
	"github.com/linusback/aoc/internal/year2024/day9"
	"github.com/linusback/aoc/internal/year2024/day10"
	"github.com/linusback/aoc/internal/year2024/day11"
	"github.com/linusback/aoc/internal/year2024/day12"
	"github.com/linusback/aoc/internal/year2024/day13"
	"github.com/linusback/aoc/internal/year2024/day14"
	"github.com/linusback/aoc/internal/year2024/day15"
	"github.com/linusback/aoc/internal/year2024/day16"
	"github.com/linusback/aoc/internal/year2024/day17"
	"github.com/linusback/aoc/internal/year2024/day18"
	"github.com/linusback/aoc/internal/year2024/day19"
	"github.com/linusback/aoc/internal/year2024/day20"
	"github.com/linusback/aoc/internal/year2024/day21"
	"github.com/linusback/aoc/internal/year2024/day22"
	"github.com/linusback/aoc/internal/year2024/day23"
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
	case "5":
		return day5.Solve()
	case "6":
		return day6.Solve()
	case "7":
		return day7.Solve()
	case "8":
		return day8.Solve()
	case "9":
		return day9.Solve()
	case "10":
		return day10.Solve()
	case "11":
		return day11.Solve()
	case "12":
		return day12.Solve()
	case "13":
		return day13.Solve()
	case "14":
		return day14.Solve()
	case "15":
		return day15.Solve()
	case "16":
		return day16.Solve()
	case "17":
		return day17.Solve()
	case "18":
		return day18.Solve()
	case "19":
		return day19.Solve()
	case "20":
		return day20.Solve()
	case "21":
		return day21.Solve()
	case "22":
		return day22.Solve()
	case "23":
		return day23.Solve()
	default:
		err = errorsx.NewSolverError(year, day, errorsx.ErrDayNotCreated)
		return
	}
}
