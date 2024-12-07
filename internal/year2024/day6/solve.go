package day6

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day6/example"
	inputFile   = "./internal/year2024/day6/input"
)

func Solve() (solution1, solution2 string, err error) {
	log.Println("welcome to day 6 of advent of code")
	g := gridContainer{
		obstacles: make(map[pos]struct{}, 1000),
	}
	s := struct{}{}
	err = util.DoEachRowFile(inputFile, func(row []byte, nr int) error {
		if nr == 0 {
			g.xMax = len(row) - 1
		}
		g.yMax = nr
		y := nr
		for x, b := range row {
			switch b {
			case '#':
				g.obstacles[pos{y, x}] = s
			case '^':
				g.guard.y = y
				g.guard.x = x
			}
		}
		return nil
	})
	//log.Println("guard starting at", g.guard)
	//log.Println("obsticles", g.obstacles)
	visited := make(map[pos]struct{}, 10000)
	var (
		ok   bool
		next pos
	)
	for !g.guardOutside() {
		visited[g.guard] = s

		next.y = g.guard.y + directions[g.dir].y
		next.x = g.guard.x + directions[g.dir].x
		if _, ok = g.obstacles[next]; ok {
			g.rotate90()
		}
		g.moveGuard()
	}

	solution1 = strconv.Itoa(len(visited))
	return
}

type pos struct {
	y, x int
}

var directions = [...]pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (p pos) Equal(o pos) bool {
	return p.y == o.y && p.x == o.x
}

type gridContainer struct {
	guard      pos
	dir        int
	obstacles  map[pos]struct{}
	xMax, yMax int
}

func (g *gridContainer) checkNext(p pos) bool {
	return p.y == g.guard.y+directions[g.dir].y && p.x == g.guard.x+directions[g.dir].x
}

func (g *gridContainer) moveGuard() {
	g.guard.y += directions[g.dir].y
	g.guard.x += directions[g.dir].x
}

func (g *gridContainer) rotate90() {
	if g.dir == 3 {
		g.dir = 0
		return
	}
	g.dir++
}

func (g *gridContainer) guardOutside() bool {
	return 0 > g.guard.y || g.guard.y > g.yMax || 0 > g.guard.x || g.guard.x > g.xMax
}
