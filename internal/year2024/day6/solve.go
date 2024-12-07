package day6

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
	"maps"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	exampleFile = "./internal/year2024/day6/example"
	inputFile   = "./internal/year2024/day6/input"
	selected    = inputFile
)

type pos struct {
	y, x uint8
}

func (p pos) Equal(o pos) bool {
	return p.y == o.y && p.x == o.x
}

type guard struct {
	pos
	dir uint8
}

//goland:noinspection GoMixedReceiverTypes
func (g *guard) rotate90() {
	if g.dir == 3 {
		g.dir = 0
		return
	}
	g.dir++
}

//goland:noinspection GoMixedReceiverTypes
func (g *guard) move() {
	g.y += directions[g.dir].y
	g.x += directions[g.dir].x
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isNextObstacles() (ok bool) {
	newPos := directions[g.dir]
	newPos.y += g.y
	newPos.x += g.x
	_, ok = obstacles[newPos]
	return
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isNextObstaclesWithNew(p pos) (ok bool) {
	newPos := directions[g.dir]
	newPos.y += g.y
	newPos.x += g.x
	_, ok = obstacles[newPos]
	return ok || p.Equal(newPos)
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isInside() bool {
	return g.y <= yMax && g.x <= xMax // -1 should wrap around to an even bigger number.
}

var (
	directions                = [...]pos{{math.MaxUint8, 0}, {0, 1}, {1, 0}, {0, math.MaxUint8}}
	securityGuard, startGuard guard

	//addedObstacles pos // for print debug
	obstacles    = make(map[pos]struct{}, 1000)
	directedPath = make(map[guard]struct{}, 10000)
	yMax, xMax   uint8
)

func Solve() (solution1, solution2 string, err error) {
	return solve(selected)
}

func solve(filename string) (solution1, solution2 string, err error) {
	startTime := time.Now()
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			if len(row) > math.MaxUint8 {
				log.Fatal("len of row needs to be less then uint8 max")
			}
			xMax = uint8(len(row)) - 1
		}
		if nr > math.MaxUint8 {
			log.Fatal("len of row needs to be less then uint8 max")
		}
		y := uint8(nr)
		yMax = y
		for xi, b := range row {
			x := uint8(xi)
			switch b {
			case '#':
				obstacles[pos{y, x}] = struct{}{}
			case '^':
				securityGuard.y = y
				securityGuard.x = x
				startGuard = securityGuard
			}
		}
		return nil
	})

	visited := make(map[pos]struct{}, 10000)
	for securityGuard.isInside() {
		visited[securityGuard.pos] = struct{}{}
		for securityGuard.isNextObstacles() {
			securityGuard.rotate90()
		}
		securityGuard.move()
	}
	solution1 = strconv.Itoa(len(visited))
	log.Println("Time part 1: ", time.Since(startTime))
	startTime = time.Now()
	//solution2 = strconv.Itoa(solve2(visited))
	solution2 = strconv.FormatUint(solve2Parallel(visited), 10)
	log.Println("Time part 2: ", time.Since(startTime))

	return
}

func solve2Parallel(visited map[pos]struct{}) uint64 {
	parallel := 20
	delete(visited, startGuard.pos)
	wg, ch := util.SeqToChannel(maps.Keys(visited), parallel*2)
	wg.Add(parallel)
	answer := new(atomic.Uint64)
	for range parallel {
		go travelParallel(wg, ch, startGuard, answer)
	}
	wg.Wait()
	return answer.Load()
}

func travelParallel(wg *sync.WaitGroup, ch <-chan pos, myGuard guard, answer *atomic.Uint64) {
	defer wg.Done()
	path := make(map[guard]struct{}, 7000)
	myStart := myGuard
	for p := range ch {
		clear(path)
		myGuard = myStart
		for myGuard.isInside() {
			for myGuard.isNextObstaclesWithNew(p) {
				myGuard.rotate90()
			}
			myGuard.move()
			if _, ok := path[myGuard]; ok {
				answer.Add(1)
				break
			}
			path[myGuard] = struct{}{}
		}
	}
}

func solve2(visited map[pos]struct{}) int {
	loopCount := 0
	// ignore first position
	delete(visited, startGuard.pos)
	for p := range visited {
		//addedObstacles = p // for print debug
		obstacles[p] = struct{}{}
		if travelNewMap() {
			loopCount++
		}
		delete(obstacles, p)
	}
	return loopCount
}

func travelNewMap() (isLoop bool) {
	//directedPrintPath := make(map[pos]int, 10000)
	clear(directedPath)
	securityGuard = startGuard
	for securityGuard.isInside() {
		//if dir, ok := directedPrintPath[securityGuard.pos]; ok {
		//	// cross
		//	if dir%2 != securityGuard.dir%2 {
		//		directedPrintPath[securityGuard.pos] = 4
		//	}
		//} else {
		//	directedPrintPath[securityGuard.pos] = securityGuard.dir
		//}
		for securityGuard.isNextObstacles() {
			//directedPrintPath[securityGuard.pos] = 4
			securityGuard.rotate90()
		}
		securityGuard.move()
		if _, ok := directedPath[securityGuard]; ok {
			//printObstaclesMap(added, directedPrintPath)
			return true
		}
		directedPath[securityGuard] = struct{}{}
	}
	return false
}

// printObstaclesMap only used for debugging highly unoptimized
func printObstaclesMap(directedPath map[pos]int) {
	sb := strings.Builder{}
	var (
		p   pos
		ok  bool
		dir int
	)
	for y := uint8(0); y <= yMax; y++ {
		for x := uint8(0); x <= xMax; x++ {
			p.y = y
			p.x = x
			// add back if debug
			//if p.Equal(addedObstacles) {
			//	_ = sb.WriteByte('O')
			//	continue
			//}
			if p.Equal(startGuard.pos) {
				_ = sb.WriteByte('^')
				continue
			}
			if _, ok = obstacles[p]; ok {
				_ = sb.WriteByte('#')
				continue
			}

			dir, ok = directedPath[p]
			if !ok {
				_ = sb.WriteByte('.')
				continue
			}
			switch dir {
			case 0, 2:
				_ = sb.WriteByte('|')
			case 1, 3:
				_ = sb.WriteByte('-')
			case 4:
				_ = sb.WriteByte('+')
			default:
				log.Fatalf("unknown dir %d while printing", dir)
			}
		}
		_ = sb.WriteByte('\n')
	}
	log.Printf("\n%s", sb.String())
}
