package day6

import (
	"github.com/linusback/aoc/pkg/util"
	"log"
	"maps"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	exampleFile = "./internal/year2024/day6/example"
	inputFile   = "./internal/year2024/day6/input"
	selected    = inputFile
)

type pos struct {
	y, x int
}

func (p pos) Equal(o pos) bool {
	return p.y == o.y && p.x == o.x
}

type guard struct {
	pos
	dir int
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
func (g guard) isNextObstacles() bool {
	newPos := directions[g.dir]
	newPos.y += g.y
	newPos.x += g.x
	return obstacles[newPos]
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isNextObstaclesParallel(myObstacles map[pos]bool) bool {
	newPos := directions[g.dir]
	newPos.y += g.y
	newPos.x += g.x
	return myObstacles[newPos]
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isInside() bool {
	return 0 <= g.y && g.y <= yMax &&
		0 <= g.x && g.x <= xMax
}

var (
	directions                = [...]pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	s                         struct{} // empty no alloc special go val
	securityGuard, startGuard guard

	//addedObstacles pos // for print debug
	obstacles    = make(map[pos]bool, 1000)
	directedPath = make(map[guard]struct{}, 10000)
	yMax, xMax   int
)

func Solve() (solution1, solution2 string, err error) {
	return solve(selected)
}

func solve(filename string) (solution1, solution2 string, err error) {
	//startTime := time.Now()
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			xMax = len(row) - 1
		}
		yMax = nr
		y := nr
		for x, b := range row {
			switch b {
			case '#':
				obstacles[pos{y, x}] = true
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
		visited[securityGuard.pos] = s
		for securityGuard.isNextObstacles() {
			securityGuard.rotate90()
		}
		securityGuard.move()
	}
	solution1 = strconv.Itoa(len(visited))
	//log.Println("Time part 1: ", time.Since(startTime))
	//startTime = time.Now()
	//solution2 = strconv.Itoa(solve2(visited))
	solution2 = strconv.FormatUint(solve2Parallel(visited), 10)
	//log.Println("Time part 2: ", time.Since(startTime))

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
	path := make(map[guard]struct{}, 10000)
	myObstacles := maps.Clone(obstacles)
	myStart := myGuard
	for p := range ch {
		myObstacles[p] = true
		clear(path)
		myGuard = myStart
		for myGuard.isInside() {
			for myGuard.isNextObstaclesParallel(myObstacles) {
				myGuard.rotate90()
			}
			myGuard.move()
			if _, ok := path[myGuard]; ok {
				answer.Add(1)
				break
			}
			path[myGuard] = s
		}
		myObstacles[p] = false
	}
}

func solve2(visited map[pos]struct{}) int {
	loopCount := 0
	// ignore first position
	delete(visited, startGuard.pos)
	for p := range visited {
		//addedObstacles = p // for print debug
		obstacles[p] = true
		if travelNewMap() {
			loopCount++
		}
		obstacles[p] = false
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
		directedPath[securityGuard] = s
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
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
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
			if obstacles[p] {
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
