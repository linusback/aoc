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
	exampleFile = "./internal/year2024/day6/example.txt"
	inputFile   = "./internal/year2024/day6/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type pos uint16

func newPos(y, x uint8) pos {
	return pos(uint16(x) | uint16(y)<<8)
}
func (p pos) y() uint8 {
	return uint8(p >> 8)
}

func (p pos) x() uint8 {
	return uint8(p & math.MaxUint8)
}

type guard struct {
	pos
	dir uint8
}

//goland:noinspection GoMixedReceiverTypes
func (g *guard) rotate90() {
	g.dir = (g.dir + 1) % 4
}

//goland:noinspection GoMixedReceiverTypes
func (g *guard) move() {
	p := directions[g.dir]
	g.pos = newPos(p.y()+g.pos.y(), p.x()+g.pos.x())
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isNextObstacles() (ok bool) {
	if g.y() == 0 && g.dir == 0 || g.x() == 0 && g.dir == 3 {
		return false
	}
	p := directions[g.dir]
	p = newPos(p.y()+g.pos.y(), p.x()+g.pos.x())
	return obstacles[p] == 1
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isNextObstaclesWithNew(p pos) (ok bool) {
	if g.y() == 0 && g.dir == 0 || g.x() == 0 && g.dir == 3 {
		return false
	}
	np := directions[g.dir]
	np = newPos(np.y()+g.pos.y(), np.x()+g.pos.x())
	return np == p || obstacles[np] == 1
}

//goland:noinspection GoMixedReceiverTypes
func (g guard) isInside() bool {
	return g.pos <= maxPos && g.pos.x() <= maxPosX // -1 should wrap around to an even bigger number.
}

var (
	directions                = [...]pos{newPos(math.MaxUint8, 0), newPos(0, 1), newPos(1, 0), newPos(0, math.MaxUint8)}
	securityGuard, startGuard guard
	//obstacles                 = [33410]uint8{}
	obstacles    = [33410]uint8{}
	directedPath = make(map[guard]struct{}, 10000)
	maxPos       pos
	maxPosX      uint8
	//addedObstacles pos // for print debug
)

func solve(filename string) (solution1, solution2 string, err error) {
	startTime := time.Now()
	var maxY uint8
	err = util.DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			if len(row) > math.MaxUint8 {
				log.Fatal("len of row needs to be less then uint8 max")
			}
			maxPos |= pos(len(row) - 1)
		}
		if nr > math.MaxUint8 {
			log.Fatal("len of row needs to be less then uint8 max")
		}
		y := uint8(nr)
		maxY = y
		for xi, b := range row {
			x := uint8(xi)
			switch b {
			case '#':
				obstacles[uint16(x)|uint16(y)<<8] = 1
			case '^':
				securityGuard.pos = newPos(y, x)
				startGuard = securityGuard
			}
		}
		return nil
	})
	maxPos |= pos(maxY) << 8
	maxPosX = maxPos.x()
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
	path := [33410][4]uint8{}
	myStart := myGuard
	for p := range ch {
		path = [33410][4]uint8{}
		myGuard = myStart
		for {
			for myGuard.isNextObstaclesWithNew(p) {
				myGuard.rotate90()
			}
			myGuard.move()
			if !myGuard.isInside() {
				break
			}
			if path[myGuard.pos][myGuard.dir] == 1 {
				answer.Add(1)
				break
			}
			path[myGuard.pos][myGuard.dir] = 1
		}
	}
}

func solve2(visited map[pos]struct{}) int {
	loopCount := 0
	// ignore first position
	delete(visited, startGuard.pos)
	for p := range visited {
		//addedObstacles = p // for print debug
		obstacles[p] = 1
		if travelNewMap() {
			loopCount++
		}
		obstacles[p] = 0
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
	for y := uint8(0); y <= maxPos.y(); y++ {
		for x := uint8(0); x <= maxPos.x(); x++ {
			p = newPos(y, x)
			// add back if debug
			//if p.Equal(addedObstacles) {
			//	_ = sb.WriteByte('O')
			//	continue
			//}
			if p == startGuard.pos {
				_ = sb.WriteByte('^')
				continue
			}
			if obstacles[p] == 1 {
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
