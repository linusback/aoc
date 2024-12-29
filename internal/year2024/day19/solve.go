package day19

import (
	"cmp"
	"github.com/linusback/aoc/pkg/util"
	"iter"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	exampleFile = "./internal/year2024/day19/example.txt"
	inputFile   = "./internal/year2024/day19/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

var (
	patterns     []pattern
	towelMap     [826][]pattern
	oneStripeMap [26]uint64
	trie         = make([]uint16, 0, 5000)
	//knownPattern  = make(map[string]uint64, 19000)
)

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseTowels, parsePatterns)
	if err != nil {
		return
	}
	//log.Println("max len", maxlen)
	//log.Println("________________________")
	//for i, t := range towelMap {
	//	if len(t) == 0 {
	//		continue
	//	}
	//	mapPattern(keyToString(i), t)
	//	//log.Printf("%c: %s\n", i, ToString(t))
	//}
	//for i, u := range oneStripeMap {
	//	if u == 0 {
	//		continue
	//	}
	//	log.Printf("%c: 1", i+'a')
	//}

	//solution1, solution2 = solveTowels()
	solution1, solution2 = solveTowelsParallel()

	return
}

func solveTowels() (solution1, solution2 string) {
	var res1, res2 uint64
	//knownPattern := make(map[string]uint64, 18500)
	knownPattern := make([]int64, 61)
	for _, t := range patterns {
		knownPattern = knownPattern[:len(t)+1]
		for k := range knownPattern {
			knownPattern[k] = -1
		}
		if ways := canBeMade(t, 0, knownPattern); ways > 0 {
			res1++
			res2 += ways
		}
	}
	//log.Println("len:", len(knownPattern))
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func solveTowelsParallel() (solution1, solution2 string) {
	const parallel = 16
	var (
		res1, res2 uint64
		n          uint32
	)
	//ch := consume(parallel)
	wg := new(sync.WaitGroup)
	wg.Add(parallel)
	pLen := uint32(len(patterns))
	for i := 0; i < parallel; i++ {
		go func() {
			defer wg.Done()
			//knownPattern := swiss.NewMap[string, uint64](2500)
			knownPattern := make([]int64, 61)
			var t pattern
			for k := atomic.AddUint32(&n, 1) - 1; k < pLen; k = atomic.AddUint32(&n, 1) - 1 {
				t = patterns[k]
				knownPattern = knownPattern[:len(t)+1]
				for kk := range knownPattern {
					knownPattern[kk] = -1
				}
				if ways := canBeMade(t, 0, knownPattern); ways > 0 {
					atomic.AddUint64(&res1, 1)
					atomic.AddUint64(&res2, ways)
				}
			}
			//fmt.Println("len:", knownPattern.Count())
			//fmt.Println("len:", len(knownPattern))

		}()
	}
	wg.Wait()
	//for _, t := range patterns {
	//	if ways := canBeMade(t, 0); ways > 0 {
	//		res1++
	//		res2 += ways
	//	}
	//}
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func canBeMade(p pattern, res uint64, knownPattern []int64) uint64 {
	key := len(p)
	switch key {
	case 0:
		return res + 1
	case 1:
		return res + oneStripeMap[p[0]-'a']
	}
	var (
		newWays  uint64
		ways     uint64
		knownWay int64
	)
	knownWay = knownPattern[key]
	if knownWay > -1 {
		return res + uint64(knownWay)
	}

	if oneStripeMap[p[0]-'a'] == 1 {
		ways = canBeMade(p[1:], res, knownPattern)
		newWays += ways
		knownPattern[key-1] = int64(ways)
	}

	towelPatters := getTowelMap(p[0], p[1])
	if len(towelPatters) == 0 {
		knownPattern[key] = 0
		return res
	}

	for _, t := range towelPatters {
		if len(p) < len(t) {
			break
		}
		if notMatch(p, t) {
			continue
		}
		key = len(p) - len(t)
		knownWay = knownPattern[key]
		if knownWay > -1 {
			newWays += uint64(knownWay)
			continue
		}
		ways = canBeMade(p[len(t):], res, knownPattern)
		newWays += ways
		knownPattern[key] = int64(ways)
	}
	return res + newWays
}

func notMatch(pattern, t pattern) bool {
	for i := 2; i < len(t); i++ {
		if pattern[i] != t[i] {
			return true
		}
	}
	return false
}

func parsePatterns(row []byte, nr int) error {
	patterns = append(patterns, row)
	//if nr > 2 {
	//	return nil
	//}
	//log.Printf("pattern %s\n", row)
	//for i := 0; i < len(row); i++ {
	//	log.Printf("%s -> %+v (%v)\n", row[i:], slices.Collect(matches(row[i:])), ToString(slices.Collect(matchesGet(row[i:]))))
	//}
	return nil
}

func parseTowels(row []byte, _ int) error {
	var (
		start, i int
		b        byte
		towel    pattern
		towels   = make([]pattern, 0, 450)
	)
	for i, b = range row {
		if util.AsciiSpace[b] == 1 {
			start++
			continue
		}
		if b == ',' {
			towel = row[start:i]
			addTowelMaps(towel)
			towels = append(towels, towel)
			start = i + 1
		}
	}
	if start < i {
		towel = row[start:]
		towels = append(towels, towel)
		addTowelMaps(towel)
	}
	slices.SortFunc(towels, patternSortPerfectHash)
	for _, t := range towelMap {
		if len(t) == 0 {
			continue
		}
		slices.SortFunc(t, patternSort)
	}
	trie = append(trie, 0, 0, 0, 0, 0, 0)
	for _, t := range towels {
		setTrieTowel(t)
	}
	//log.Println(towels)
	//log.Println(trie)
	//log.Printf("trie len %d, cap %d\n", len(trie), cap(trie))
	//log.Println(math.MaxUint16)
	//log.Println("len", len(towels))
	return nil
}

func patternSort(a, b pattern) int {
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return cmp.Compare(string(a), string(b))
}

func patternSortPerfectHash(a, b pattern) int {
	var ha, hb uint16
	for i := 0; i < len(a) && i < len(b); i++ {
		ha, hb = perfectHash(a[i]), perfectHash(b[i])
		if ha < hb {
			return -1
		}
		if ha > hb {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func patternSortPerfectHashOld(a, b pattern) int {
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	var ha, hb uint16
	for i := 0; i < len(a); i++ {
		ha, hb = perfectHash(a[i]), perfectHash(b[i])
		if ha < hb {
			return -1
		}
		if ha > hb {
			return 1
		}
	}
	return 0
}

func getTowelMap(a, b byte) []pattern {
	return towelMap[uint16(a-'a')<<5|uint16(b-'a')]
}

func addTowelMaps(p pattern) {
	switch len(p) {
	case 0:
		return
	case 1:
		oneStripeMap[p[0]-'a'] = 1
		return
	}
	key := uint16(p[0]-'a')<<5 | uint16(p[1]-'a')
	towelMap[key] = append(towelMap[key], p)
}

// perfectHash is adapted from https://github.com/maneatingape/advent-of-code-rust/blob/main/src/year2024/day19.rs
// / Hashes the five possible color values white (w), blue (u), black (b), red (r), or green (g)
// / to 0, 2, 4, 5 and 1 respectively. This compresses the range to fit into an array of 6 elements.
func perfectHash(b byte) uint16 {
	return uint16((b ^ (b >> 4)) & 7)
}

func setTrieTowel(towel pattern) {
	var i, j, tLen uint16
	//log.Printf("setting towel %v\n", towel)
	for _, b := range towel {

		j = perfectHash(b)
		//log.Printf("hash: %c -> %d\n", b, j)
		//log.Printf("%d + %d = %d, %d", i, j, i+j, len(trie))
		if trie[i+j] == 0 {
			tLen = uint16(len(trie))
			trie[i+j] = tLen
			i = tLen
			trie = append(trie, 0, 0, 0, 0, 0, 0)
		} else {
			i = trie[i+j]
		}
	}
	trie[i+3] = 1
}

func matchesGet(p pattern) iter.Seq[pattern] {
	tLen := min(len(p), 8)
	towel := make(pattern, 0, 8)
	var trieI uint16
	//log.Println("initial:", trie[trieI:trieI+6])
	return func(yield func(pattern) bool) {
		for i := 0; i < tLen; i++ {
			towel = append(towel, p[i])
			//log.Printf("i: %d, hash: %d\n", trieI, perfectHash(p[i]))
			//log.Println("current:", trie[trieI:trieI+6])
			trieI = trie[trieI+perfectHash(p[i])]
			if trie[trieI+3] == 1 {
				if !yield(towel) {
					break
				}
			}
		}
	}
}

func matches(p pattern) iter.Seq[uint64] {
	tLen := min(len(p), 8)
	var (
		trieI uint16
		res   uint64
	)

	return func(yield func(uint64) bool) {
		for i := 0; i < tLen; i++ {
			res++
			trieI = trie[trieI+perfectHash(p[i])]
			if trie[trieI+3] == 1 {
				if !yield(res) {
					break
				}
			}
		}
	}
}

//func keyToString(i int) string {
//	k := uint16(i)
//	const mask = 1<<5 - 1
//	return fmt.Sprintf("%c%c", byte(k>>5)+'a', byte(k&mask)+'a')
//}

//func consume(parallel int) <-chan pattern {
//	ch := make(chan pattern, parallel)
//	go func() {
//		for _, p := range patterns {
//			ch <- p
//		}
//		close(ch)
//	}()
//	return ch
//}
